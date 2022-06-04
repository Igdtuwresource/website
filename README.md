OpenSourceCorp website content
==============================

This repo houses all of the content on [OpenSourceCorp's public
website](https://opensourcecorp.org).

Currently, the site is served via [GitHub
Pages](https://docs.github.com/en/pages), and you can therefore see the public
content published at [the canonical GH Pages
repo](https://github.com/opensourcecorp/opensourcecorp.github.io).

Development (in progress)
-------------------------

The OSC website is a static site rendered via [Hugo](https://gohugo.io). As
such, adding new content can be added according to Hugo's documentation. New
pages for existing sections will always go in the relevant folder under
`content/`. If you want to add a new section entirely, review Hugo's docs for
what that would entail (spoiler: slightly more files).

The `Makefile` in this repo helps to orchestrate website generation, preview,
and publishing. Some notable `make` targets are listed in the table below.

| Make target            | Description
| :----------            | :----------
| `render<-dev>`         | Generate static content from templates, with or without drafts
| `serve<-dev>`          | Run Hugo server locally (`:1313`) to serve content, with or without drafts
| `image-build`          | Build OCI image to serve content via a Caddy server instead of Hugo's server
| `<run/stop>-container` | Starts or stops the running Caddy container (`:2015`)
| `github-pages`         | Copy site content to adjacent GitHub Pages repo, if you need to review before publishing
| `publish`              | Copy site content to adjacent GitHub Pages repo, commit, and push to publish changes
