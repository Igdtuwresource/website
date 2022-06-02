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

serve:
	@hugo server -D --disableFastRender

test:
	@bash ./scripts/test.sh

clean:
	@rm -rf \
		public/

image-build: clean
	@$(DOCKER) build -f Containerfile -t $(OCI_REGISTRY)/$(OCI_REGISTRY_OWNER)/$(PKGNAME):latest .

image-run:
	@$(DOCKER) run --rm -it -p 2015:2015 --name $(PKGNAME) $(OCI_REGISTRY)/$(OCI_REGISTRY_OWNER)/$(PKGNAME):latest
