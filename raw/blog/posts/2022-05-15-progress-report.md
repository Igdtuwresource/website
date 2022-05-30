---
layout: post
title:  "Progress Report 2022-05-15"
date:   2022-05-15 12:00:00 -0600
# categories: jekyll update
author: Ryan J. Price
---

(*Checks date of last post, sweats profusely*)

Not a ton of changes in this update, I've had a lot of things come up these past
few months that have prioritized my attention (a few conference talks to prepare
for, dog & personal health issues, house repairs, work-work, etc). But some
impactful updates nonetheless!

* All OSC infrastructure repos have now been consolidated into a single
  monorepo, [`osc-infra`](https://github.com/opensourcecorp/osc-infra). This
  implies that you should expect old links to the individual infra repos to be
  broken.

  OSC applications themselves (like `ghostwriter`) will still live in top-level
  repos for now, until I decide otherwise -- I'd *like* to have an `osc-apps`
  monorepo as well, but still need to tackle that with a more academic approach
  with respect to versioning, release, etc. This will probably be a custom-built
  toolset and *not* something like [Bazel](https://bazel.build/); though we can
  probably learn a lot from some parts of that tool.

* Each infra component has historically been referred to as "apps" throughout
  the source code. To better represent what these components actually do, they
  will now be referred to as "subsystems". `osc-infra` docs and most of its code
  should currently reflect that, but I will likely need to update more code as
  well. So, if you see any straggling references to an infra component referred
  to as "app"/"app_name", know that this means "subsystem" in the new
  vocabulary.

* The `osc-infra` bootstrapper (now a subdir of that repo) now supports another
  env var, `$OSC_SUBSYSTEMS`, that allows safer control over locally-deployed
  VMs with respect to their aggregate memory usage. Deploying the entire infra
  stack locally at this time should consume around 12GB of free memory, and this
  is unrealistic for most host systems. So, the bootstrapper currently defaults
  to *building* all subsystems, but only *running* the core three (`aether`,
  `faro`, and `chonk`). To run others, export the `OSC_SUBSYSTEMS` var as a
  comma-separated string of subsystem names, e.g.:

      $ export OSC_SUBSYSTEMS=photobook,gnar
      $ ./bootstrap.sh local-vm up

* I had previously gotten a working deployment of [Harbor](https://goharbor.io/)
  to serve as the OCI image registry for OSC infra. Harbor is now gone, and
  replaced with the [Docker Registry server](https://docs.docker.com/registry/)
  instead. Harbor's HTTP API was poorly-documented, and poorly-implemented, in
  my not-so-humble opinion -- plus the setup using its bundled installer is too
  opaque for my liking, and goes against the principles of OSC as a learning
  utility. I find the Docker Registry server by comparison to be way simpler to
  deploy & to understand, for me & others alike.

* To serve many different, relevant use cases, I have waffled back & forth on
  whether to use [GitLab](https://gitlab.com) or not. GitLab provides many
  consolidated features that are of excellent quality & diversity -- source code
  management, a CI/CD system, an OCI image registry, an artifact store, and
  more. It's a great tool.

  However, I have ultimately decided against using GitLab for OSC. Perhaps the
  primary reason being that with so many features tied to a single platform, I
  am worried that the modularity is far too limited to allow for
  individualization of OSC subsystems.

  So, OSC infra will stick with the original plan of having fully-modular
  subsystems instead of the consolidated features that GitLab provides:
  * [Gitea](https://gitea.io) as the SCM, as originally planned
  * Docker Registry server as OCI image store, and [an unofficial registry
    UI](https://github.com/Joxit/docker-registry-ui) as a companion service (the
    latter is WIP)
  * [Concourse](concourse-ci.org/) will remain as the CI/CD subsystem
  * A generic artifact store is still up for decision -- other options I've come
    across aren't... well, aren't that great, or aren't free (or both)

* Unrelated directly to OSC, but my host PC now runs Debian instead of Manjaro
  Linux. In the past, some discrepancies between host tooling, versions, and
  package names has made it difficult to jump between development machines'
  operating systems, etc. This host change will hopefully leading to a more
  consistent developer experience between hosts, and between hosts & deployments
  (since all infra subsystems run on Debian Stable).

There are surely some more minor notes to add somewhere, but the `osc-infra`
consolidation took away the VCS history of the individual infra repos -- and
naturally I didn't comb through their logs before deleting them. But overall,
the monorepo consolidation has already led to a lot of codebase & development
benefits that I'd hoped to see, and I look forward to seeing this tooling
continue to mature!
