package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/opensourcecorp/website/logging"
	"github.com/yuin/goldmark"
)

// const doNotEditHTMLComment = "<!-- This document was rendered from a template, DO NOT EDIT DIRECTLY -->"

type tplInfo struct {
	inPath   string
	tpl      string
	rendered string
}

type renderData map[string]any

func getAllFiles(root string) []string {
	var files []string

	// This is such an ugly way to walk directories and get the files
	// (filepath.Glob doesn't recurse deep enough), but... Go things
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		// Skip .git, and already-rendered directories
		if filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}

		fileInfo, err := os.Stat(path)
		if err != nil {
			logging.Error("Could not process filepath %s for some reason; error specifics below\n", path)
			logging.Error(err.Error())
		}

		// Only return files, not directories; and also do a weak skip of the
		// config file
		if !fileInfo.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		logging.Error(err.Error())
	}

	return files
}

func getNearbyRenderData(tplFilePath string) renderData {
	var err error
	var data renderData

	tplDir := filepath.Dir(tplFilePath)
	dataPath := filepath.Join(tplDir, "_data.json")
	
	_, err = os.Lstat(dataPath)
	if errors.Is(err, fs.ErrNotExist) {
		// No need to have actual data if there's none provided
		return renderData{}
	} else if err != nil {
		logging.Error("Error trying to determine stat on path '%s'", dataPath)
		logging.Error(err.Error())
		os.Exit(1)
	}

	dataBytes, err := os.ReadFile(dataPath)
	if err != nil {
		logging.Error("Error reading in file at '%s'", tplFilePath)
		logging.Error(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		logging.Error("Error while processing _data.json file at '%s'", tplFilePath)
		logging.Error(err.Error())
		os.Exit(1)
	}

	return data
}

func render(tplFilePath string) tplInfo {
	var tplBytes []byte
	var rendered bytes.Buffer
	var data renderData
	var err error

	tplBytes, err = os.ReadFile(tplFilePath)
	if err != nil {
		logging.Error("Error reading in file at '%s'", tplFilePath)
		logging.Error(err.Error())
		os.Exit(1)
	}

	indexJSONRegex := regexp.MustCompile(`.*_index\.json$`)
	markdownRegex := regexp.MustCompile(`.*\.(md|markdown)$`)

	if indexJSONRegex.MatchString(tplFilePath) {
		// Populate subdir index.html files as rendered versions of the site
		// root index.html
		tplBytes, err = os.ReadFile("./raw/index.html.tpl")
		if err != nil {
			logging.Error("Error reading in file at '%s'", tplFilePath)
			logging.Error(err.Error())
			os.Exit(1)
		}
		dataBytes, err := os.ReadFile(tplFilePath)
		if err != nil {
			logging.Error("Error reading in file at '%s'", tplFilePath)
			logging.Error(err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(dataBytes, &data)
		if err != nil {
			logging.Error("Error while processing _index.json file at '%s'", tplFilePath)
			logging.Error(err.Error())
			os.Exit(1)
		}

		tpl, err := template.New("x").Parse(string(tplBytes))
		if err != nil {
			logging.Error("Couldn't parse file '%s' for some reason; bad template formatting?")
			logging.Error(tplFilePath)
			logging.Error(err.Error())
			os.Exit(1)
		}

		// Add additional render data, if any
		addlData := getNearbyRenderData(tplFilePath)
		for k := range addlData {
			data[k] = addlData[k]
		}

		err = tpl.Execute(&rendered, data)
		if err != nil {
			logging.Error("Could not succesfully render the text from the template from '%s'", tplFilePath)
			logging.Error(err.Error())
			os.Exit(1)
		}

		tplFilePath = strings.ReplaceAll(tplFilePath, ".json", ".html")
		tplFilePath = strings.ReplaceAll(tplFilePath, "_index.html", "index.html")

	} else if markdownRegex.MatchString(tplFilePath) {
		goldmark.Convert(tplBytes, &rendered)
		tplFilePath = strings.ReplaceAll(tplFilePath, ".md", ".html")

	} else { // assume everything else is OK to parse as an HTML or text template
		tpl, err := template.New("x").Parse(string(tplBytes))
		if err != nil {
			logging.Error("Couldn't parse file '%s' for some reason; bad template formatting?")
			logging.Error(tplFilePath)
			logging.Error(err.Error())
			os.Exit(1)
		}

		data = make(renderData)

		// Add additional render data, if any
		addlData := getNearbyRenderData(tplFilePath)
		for k := range addlData {
			data[k] = addlData[k]
		}

		err = tpl.Execute(&rendered, data)
		if err != nil {
			logging.Error("Could not succesfully render the text from the template from '%s'", tplFilePath)
			logging.Error(err.Error())
			os.Exit(1)
		}
	}

	return tplInfo{
		tplFilePath,
		string(tplBytes),
		rendered.String(),
	}
}

func writeRendered(tpl tplInfo, outDir string) {
	var err error

	// outPath is built by removing the root 'raw/' directory from the start of
	// the inPath -- otherwise 'raw/' is the top of the rendered tree
	outPath := filepath.Join(outDir, strings.Replace(tpl.inPath, "raw/", "", 1))
	if filepath.Ext(outPath) == ".tpl" {
		outPath = strings.ReplaceAll(outPath, ".tpl", "")
	}

	_, err = os.Lstat(filepath.Dir(outPath))
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(filepath.Dir(outPath), 0755)
		if err != nil {
			logging.Error("Could not succesfully create some director(y|ies) needed for rendering output")
			logging.Error(err.Error())
			os.Exit(1)
		}
	} else if err != nil {
		logging.Error("Error trying to determine stat on path '%s'", outDir)
		logging.Error(err.Error())
		os.Exit(1)
	}

	err = os.WriteFile(
		outPath,
		[]byte(tpl.rendered),
		0644,
	)
	if err != nil {
		logging.Error("Could not successfully write to output path '%s'", outPath)
		logging.Error(err.Error())
	}
}

func main() {
	files := getAllFiles("./raw/")

	for _, file := range files {
		writeRendered(render(file), "site")
	}
}
