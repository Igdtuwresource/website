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
and publishing. After each commit made to this repo, you can run `make publish`
to try and automatically publish website content.
