---
layout: post
title:  "Progress Report 2022-01-02"
date:   2022-01-02 19:30:00 -0500
# categories: jekyll update
author: Ryan J. Price
---

Happy New Year!

Over the holiday season, I've had about a week of dedicated time to do a lot of
net-new work for OSC, as well as to go back and do some much-needed cleanup
& consolidation across several of the core codebases.

First & foremost, we've now got a working OCI image registry!
[`gnar`](https://github.com/opensourcecorp/gnar) needs one for custom images, so
it turned into a priority pretty quick. The registry is called
[`photobook`](https://github.com/opensourcecorp/photobook), which is our
implementation of the [Harbor registry](https://goharbor.io) software. Getting
Harbor up & running was more of a challenge than I expected -- it's
documentation isn't as robust as I'd have liked to see, particularly its HTTP
API docs (you can only browse said docs on a running instance of Harbor). I
really needed the API docs because there are several things `photobook` needs to
do at build time: set up a Docker Hub mirror, set up replication jobs for common
base images from said mirror, and surely more administrative tasks in the future
(like building some OSC images so they're available at boot time). I managed to
cobble together what I could from the docs, and my own trial & error, but it's
become a bit of an ugly set of repeated `curl` calls in the
[`aether`](https://github.com/opensourcecorp/aether) code for `photobook`.
Something for us to work on in the future might be making a Harbor client
library -- but, we'll see how much more complex the build-time config grows to
be first, I think.

However, Harbor was chosen *because* of its broad set of features as touched on
above, so it was developer time I was willing to spend, and maintenance I'm
willing to keep up with.

The second major update (and much more invisibly impactful, IMHO) is that all
OSC images will now pull & store common scripts from the
[`ymir`](https://github.com/opensourcecorp/ymir) codebase at build time. This
means that *any* repos using OSC's image-build toolchain now need to bring much
less of their own config for builds & boots -- much has now been abstracted to
`ymir`'s `scripts/{build,run}/` directories. `ymir` ended up being a great
location to house common build & runtime logic (outside of `aether`) because
image builds happen within the `ymir` codebase anyway, so anything stored there
is already available for use.

Additionally, those common config scripts now ensure that `aether` minion IDs
are automatically suffixed with a random hex string. This allows for scaled
services without conflicting IDs or private keys, which has been a major issue
until now unless every single build-time job provided their own randomized
suffix. The last reamining duplicate-ID pain point is that `gnar` workers
provide their node IDs to the web server via their hostnames -- which means that
`n` number of worker nodes will all show up as distinct workers in the cluster,
but as a *single hostname* in the `fly` CLI. This isn't a difficult fix, but
it's not one we've broached just yet, since local testing just uses a single
worker node.

And the final update: we've managed to secure a domain!

[opensourcecorp.org](https://opensourcecorp.org) is where we'll eventually be
building out site content for the project at large. As of today, it just
redirects the DNS query to the project Github page, but be on the lookout for
when we start adding records to point to subdomains. ***However***, paths *do*
get forwarded on that redirect, so e.g. `opensourcecorp.org/ymir` sends you
straight to `github.com/opensourcecorp/ymir`! That's pretty neat for the time
being!

All in all, there's been a lot of great progress this past week or two at OSC,
and I'm SUPER excited about what 2022 will bring. I'm looking forward to getting
the community more involved in the work OSC wants to do, to explore other
use-cases that we can come up with, and to work to make even more cool stuff.
See you in the new year!
