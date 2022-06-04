SHELL = /usr/bin/env bash

PKGNAME ?= website

DOCKER ?= docker
OCI_REGISTRY ?= ociregistry.opensourcecorp.org
OCI_REGISTRY_OWNER ?= library

.PHONY: %

all: render

render: clean
	@hugo

render-dev: clean
	@hugo -D

serve: clean
	@hugo server --disableFastRender

serve-dev: clean
	@hugo server -D --disableFastRender

test: run
	@bash ./scripts/test.sh

clean:
	@rm -rf \
		public/

image-build: clean
	@$(DOCKER) build -f Containerfile -t $(OCI_REGISTRY)/$(OCI_REGISTRY_OWNER)/$(PKGNAME):latest .

run-container: image-build stop-container
	@$(DOCKER) run \
		--rm -dit \
		-p 2015:2015 \
		--name $(PKGNAME) \
		$(OCI_REGISTRY)/$(OCI_REGISTRY_OWNER)/$(PKGNAME):latest
	@sleep 2
	@$(DOCKER) logs $(PKGNAME)
	@printf '\n^^^ Website should be running; review logs above to confirm ^^^\n'

stop-container:
	@$(DOCKER) stop $(PKGNAME) > /dev/null 2>&1 || true

# This is here until we're no longer serving via GitHub Pages -- but it's free, so.
github-pages: clean render
	@cp -r public/* ../opensourcecorp.github.io/

publish: github-pages
	@git -C ../opensourcecorp.github.io commit -am "Updates from website repo"
	@git -C ../opensourcecorp.github.io push
