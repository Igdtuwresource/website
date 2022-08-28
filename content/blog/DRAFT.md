---
title: "DRAFT"
publishdate: 2021-01-01T00:00:00-05:00
draft: true
---

This document is a living draft of changes across OSC to put into blog posts. It
should remain in `draft` mode in the front matter, and as such can be committed
to source control without appearing on the site.

Updates
-------

* Work has been done to the `osc-infra` bootstrapper to support AWS deployments!
  We're using TF provisioners like sinners. It's neat.

* `osc-infra` now v0.2.0

* `local-vm` bootstrapper refactored to not build separate images, and just
  provision at runtime. Way faster to bootstrap, no persistent storage of other
  images, etc. This came along with a number of other cleanup changes while I
  was crawling around in the codebase, e.g. the `imgbuilder` subsystem is now
  named `baseimg`.

* `redis` is now installed via APT repo, vs. source build. Somehow the docs page
  I read from initially now shows how to get via `apt` instead :shrug:

* The Salt block adding a swapfile was commented out, but now it's not, and is
  working as expected

* `rhad` is now at v0.3.0. Most of the updates were a lot of internal
  refactoring that aren't going to be very visible to users, but the
  minor-version bump is because `rhad` now runs `go vet` as part of its Go
  linter runs. New feature, new minor version. Internal changes include a
  filesystem-based state tracking mechanism (previously `rhad` used global vars
  to track run state), and decoupling of the logging facilities to [a separate
  repo for OSC's common Go
  tooling](https://github.com/opensourcecorp/go-common).
