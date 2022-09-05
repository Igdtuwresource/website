---
title: "Progress Report 2022-09-05"
publishdate: 2022-09-05T01:45:00-05:00
author: "Ryan J. Price"
---

Wooooo boy howdy, what a STRETCH it's been. Lots of things in life & work have
come up and kept me away from working on OSC, but I wanted to throw out
something quick to let folks know that neither I nor the project are dead.

* [`rhad`](https://github.com/opensourcecorp/rhad) has been under a lot of
  recent work. Some months ago, the concept of a `Rhadfile` was introduced to
  the codebase, which is a config-file format for `rhad` executions. The first
  pass at this was targeting an INI format, for its simplicity & readability. As
  it turns out, parsing an INI file is... interesting, to say the least. So,
  `Rhadfiles` are now [TOML](https://toml.io)-formatted. TOML is a nice
  middle-ground format of being readable, writable, and parseable. Right now,
  `Rhadfiles` only support providing a version number for your code tree
  modules, but more options will spring up as they become apparent.

  There's also [a helper
  script](https://github.com/opensourcecorp/rhad/pkgs/container/rhad/tree/main/scripts/run-rhad-lint.sh)
  that runs the GitHub Super-Linter, and `rhad` following, as well as a `make
  add-local-symlinks` target to symlink the script onto (what I hope is) your
  `$PATH`. Try it out and let me know if you run into any notable bugs.

  `rhad` is now version `v0.3.0`, and the [container image on
  GHCR](https://github.com/opensourcecorp/rhad/pkgs/container/rhad) reflects
  this in its `:latest` tag (I'll get around to getting version-tagged builds
  done in GHA at some point).

* While working on `rhad`, I started to think that abstracting out some of the
  reusable Go functionality might be beneficial in the long run -- so now we
  have [`a repo to house that`](https://github.com/opensourcecorp/go-common)
  (currently it's just logging functionality). This kind of felt like premature
  optimization, since nothing else except `ghostwriter` is using this shareable
  functionality, but it ended up teaching me an awful lot about how to work with
  dependency packages in Go that *you* control -- like being able to override
  exported fields in child packages to disable log output for tests, set custom
  log prefixes without boilerplate, etc. Seeing it all come together was a cool
  feeling.

And now for the parts that I haven't touched in long enough that I forget most
of the finer details:

* The [`osc-infra`
  bootstrapper](https://github.com/opensourcecorp/osc-infra/tree/feature/bootstrapper-add-aws)
  is almost ready to support AWS! we're using Terraform Provisioners like
  sinners. It's neat. Also, as mentioned in the last post `osc-infra` is now
  under tagged releases, and the most recent version is now v0.2.0.

* The `local-vm` bootstrapper has been refactored to not build separate images
  for ever ysingle subsystem, and just run provision calls at runtime. This
  makes the OSC cluster ***way faster*** to bootstrap, saves on disk space since
  there's no persistent storage of other images, etc. This also came along with
  a number of other cleanup changes while I was crawling around in the codebase,
  e.g. the `imgbuilder` subsystem is now named `baseimg`.

* Previously, when setting up a `redis` service on `datastore`, I must have
  stumbled onto a docs page that told readers that the ideal way to install
  `redis` was to compile it from source, and ***not*** use something like an APT
  repo. I'm not sure what I found, or if something changed, but `redis` is now
  installed via an APT repo, and it works fine `:shrug:`

* `configmgmt` had a Salt block that added a swapfile to cluster nodes, and that
  block was commented out. I don't recall ***why*** it was commented out, but
  now it's not, and is working as expected. All nodes now have a swapfile
  (insert Oprah "you get a car" meme here).

Like I said above, this isn't a massive update, nor is it super clean, but
posting all this to say that I've not disappeared, nor stopped working on OSC.
Life, uh, finds a way (to get in the way).
