---
layout: post
title:  "OSC updates: hosting decision & live services!"
date:   2021-10-04 21:15:00 -0500
# categories: jekyll update
author: Ryan J. Price
---

[OpenSourceCorp](https://github.com/opensourcecorp/opensourcecorp) (OSC)
development has slowed down in recent months as
[Ryan](https://github.com/ryapric) (founder & maintainer) has gotten caught up
in *work*-work. However, with renewed motivation (looming conference talk at
[SaltConf21](https://saltconf.com/), *presenting on OSC!*), we're back at it
again!

One of the chief goals of OSC has been to have its services live on an
as-unmanaged-as-possible platform, in the service of helping to teach the
community all about the core functionality of computing, storage, networking,
and security. We have deliberately steered *away* from public cloud hosting such
as Amazon Web Services or DigitalOcean in pursuit of this goal. However, thanks
in part to the current market prices for computer/server cluster components,
setting up a meaningfully-powerful & redundant home system (on
[Proxmox](https://www.proxmox.com/en/), for example) is cost-prohibitive, and
until OSC builds more momentum this cost is likely not very warranted.

So, for the time being, OSC will indeed be hosted on AWS. We will continue to
work towards our goal of self-hosting, but the flexibility & cost (especially
using Spot Instances) will allow us to iterate in the open much more quickly.

The other good news here is that public clouds like AWS tend to be more
approachable out of the gate than our desired solution, so this should still be
beneficial to users of the platform.

All that being said: ***we now have live services!***

You can't see them yet, because they're all internal
([ymir](https://github.com/opensourcecorp/ymir),
[aether](https://github.com/opensourcecorp/aether), &
[chonk](https://github.com/opensourcecorp/chonk), specifically, with
[gnar](https://github.com/opensourcecorp/gnar) coming next); but seeing all the
work & development tooling come together and so well-replicate the local dev
environments has been *very* rewarding.

The promise & reality of OSC is starting to all come together, and we're looking
forward to you being a part of it!
