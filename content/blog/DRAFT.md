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

* `cicd` replatformed from Concourse CI to... JENKINS! Believe it or not!

* Vault is up! Still need to get a) secrets into it, b) HA, and c) a way to get
  the unseal keys off the machine post-build, but that will come in time.

* For now, cut `rhad` linting functionality down to just what GitHub's
  Super-Linter doesn't handle well

* Blog has RSS feed link (but the feed itself had existed already)
