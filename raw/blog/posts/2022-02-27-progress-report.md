---
layout: post
title:  "Progress Report 2022-02-27"
date:   2022-02-27 12:00:00 -0600
# categories: jekyll update
author: Ryan J. Price
---

More radio silence as ever more work has been done across OSC. As you'll see
from the list below though, this is a really big set of changes/new stuff!

* The conflation of infrastructure repos with application repos has grown to be
  a potential source of confusion for readers. As such, OSC should shortly be
  moving from a mono-org source, to orgs separated by identifier
  (`opensourcecorp-infra`, `opensourcecorp-apps`, etc). This should make it more
  clear what type of repo you're looking at when you're browsing a certain org.

* All infra services across OSC now use internal HTTPS / TLS -- OSC itself
  behaves as a root CA, its root CA pubkey is in the system cert store of each
  service, and each service signs its own intermediate TLS certificates using
  that root CA key. Public-facing services will eventually use LetsEncrypt in
  addition, but we're not there yet since we have no public-facing services.

* [`chonk`](https://github.com/opensourcecorp/chonk) Postgres is now configured
  for TLS connections, and requires all clients to specify some manner of TLS in
  their connection config. Verification method is up to the client, since some
  are much more difficult to `verify-full` (namely,
  [`photobook`](https://github.com/opensourcecorp/photobook) because of the
  finnicky nature of its cert mounts, but there will undoubtedly be others).

* `chonk` Postgres now has synchronous replication configured successfully.

* `chonk`'s Vagrantfile now optionally mounts a persistent data dir that runs a
  backup & restore during shutdown & startup, respectively. It's not working
  exactly as intended just yet, but the code is there to tweak in the future.

* `chonk` now includes a [`redis`](https://redis.io) server, secured with TLS
  and (currently blanket) password auth. Some infra services require a `redis`
  server, and most even bring their own embedded -- but that's not something we
  want to do, and so rolling our own will allow for consistency in expectation
  of "where does this data live/go". Replication for the `redis` server will be
  coming soon as well.

* [`rhad`](https://github.com/opensourcecorp/rhadamanthus) is now on the GitHub
  Container Registry, and apps use `rhad` in their own GHA pipelines, just like
  they will within [`gnar`](https://github.com/opensourcecorp/gnar). However,
  the GHA container build stage recently started to fail seemingly without any
  code changes (some kind of auth issue), so the image is outdated until that
  gets fixed.

* `rhad` now runs the [Black](https://github.com/psf/black) formatter check for
  Python code.

* Fixed a bug in `rhad` that made it exit on the first error in some linter
  runs. It will now run in its entirety before exiting.

* [`sauce`](https://github.com/opensourcecorp/sauce) (source code management)
  is... not MVP yet. I wanted to go with [Gitea](https://gitea.io), which is a
  great option, and even had a working deployment for it. But as other
  components across OSC were being explored (artifact repository, container
  image registry, etc.), [GitLab](https://gitlab.com) started to make more &
  more sense. I'll be getting an instance of that up & working relatively soon.
  `sauce` will take over `photobook`'s functionality (I really don't like
  Harbor), as well as the other features that GitLab provides.

* [`scrawl`](https://github.com/opensourcecorp/scrawl) is a new app -- it's a
  simple note-taker that supports tag search. Lots of stuff planned for it,
  while keeping it super simple & lightweight. Stay tuned.

The above is mostly a collective brain dump of what's been being worked on
across OSC, with more to come all the time. So as always, stay tuned for
continued updates as we build out the rest of the platform!
