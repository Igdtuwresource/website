---
layout: post
title:  "OSC local bootstrapper available"
date:   2021-12-11 16:00:00 -0500
# categories: jekyll update
author: Ryan J. Price
---

[OpenSourceCorp](https://github.com/opensourcecorp/opensourcecorp) (OSC)
development continues to chug along. Below are some of the updates we've been
hard at work on, when time presents itself.

One of the net-new things we've worked on recently is the introduction of a
(working) service-discovery implementation via [HashiCorp
Consul](https://www.consul.io) (named
"[faro](https://github.com/opensourcecorp/faro)", which means "lighthouse" in a
few Latin-root languages). Updates to
[aether](https://github.com/opensourcecorp/aether) allow for each platform/app
to easily register themselves with the Consul cluster, and be discoverable via
clean domain names and not IP addresses anymore, with two foundational
exceptions that still need known IPs at their first launch (`faro` itself, and
`aether`). We took out a lot of manual IP pointing with this change, and it
feels *awesome!*

The other major news is that we've been developing a local infrastructure
bootstrapper on & off for some time now, and it's finally in a state that should
work for others! Using all of the regular OSC local-workstation dev tools, the
utility builds the VM images, starts them up in order, runs some sanity checks
on the cluster, and then the user can poke around to learn more about how the
OSC infrastructure works together, is extended upon, etc.

***We would really appreciate the community taking the time to test it out***.
Ryan Price (founder/maintainer of OSC) works off of a native Linux development
environment, and so there's been no testing of the bootstrapper done against
macOS or Windows (the latter is sure to be hacky). The goal of OSC is for all
dev work to be done on a Linux-based OS in the first place for consistency, but
we recognize that not everyone has a high-powered Linux dev machine to use to
reap the fruits of all this labor `:)`

[The bootstrapper is available
here](https://github.com/opensourcecorp/osc-infra-bootstrap), with instructions
for use in the `README`. For any questions, suggestions, or comments, please
reach out, open a Github PR/Issue, etc!
