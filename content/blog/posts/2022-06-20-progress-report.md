---
title: "Progress Report 2022-06-20"
publishdate: 2022-06-20T15:00:00-05:00
author: "Ryan J. Price"
---

Some heavy updates in this one, so buckle up for what's changed!

### Secrets management

Up until now, OSC has been strictly leveraging an implementation of Salt's
Pillar functionality for secrets management -- and that implementation is the
default, which is "store all secrets on the `configmgmt` disk". While this has
worked fine (and can be a valid strategy in some enterprise environments), it's
not what you would call "best-practice". If the `configmgmt` subsystem itself is
ever compromised, *every single piece of sensitive data* is accessible to the
attacker at once -- default Pillar data is just stored in plaintext on disk.

I've had a dedicated secrets-management solution on the roadmap since the
beginning at OSC, but have spent all time since getting the other subsystems up
& running first. And now, that solution is successfully set up as another OSC
subsystem. Named simply `secretsmgmt`, OSC's implementation of [HashiCorp
Vault](https://vaultproject.io) is now successfully building as part of the
bootstrapper!

There's still a lot of work to do here though:

* There's no actual secrets in Vault right now. The deployment & unsealing of
  the subsystem is live, but it has no secrets, nor are any other subsystems
  currently using it. The good news here though is that Salt Pillar supports
  adding Vault as an "external Pillar", which means the other subsystems may not
  even need to know anything about Vault and just continue using Pillar data as
  usual. I'll explore more what those security implications are though.

* In order to get a Vault cluster actually ready for use, it must first be
  "unsealed". Until a cluster is unsealed, Vault cannot encrypt/decrypt any
  secrets, and indeed is unusable as a consequence. While the deployment does
  currently handle the unseal operation, the unseal key shards used are
  discarded post-setup. This is OK for now, but if the cluster ever goes down or
  needs to be recreated, the original unseal key shards are the only way to
  reactivate the cluster to a working state. This is flagged as a `TODO:`, so
  we'll revisit in time.

* The cluster is not currently highly-available (HA) -- it's just a single node.
  Vault has native facilities for multi-node use, I just haven't set them up
  yet.

Overall, this is still a great start, and having a dedicated secrets manager is
a good thing to have, and something you will see in mature enterprise
environments.

### CI/CD replatform

Being the very first subsystem worked on after the `imgbuilder` & `configmgmt`
bases, the `cicd` subsystem is the oldest in OSC's infrastructure. Originally, I
wanted to find a tool that:

(a) was not "cloud-native" (specifically: not requiring a Kubernetes cluster)

(b) did not have any kind of commercial tier available to it (so, no
"open-core")

(c) has centralizable pipeline definitions so every repo need not bring its own
def

(d) has pipeline jobs defined in something readable/maintainable and not
pseudo-script YAML.

There was also an implicit (e) requirement of not requiring the JVM, because I
have a visceral reaction to needing to extend the subsystem using Java (I don't
like it lol). So that removed from consideration a number of tools, such as
[Spinnaker](https://spinnaker.io), [Jenkins](https://jenkins.io), and others.

So, we went with [Concourse CI](https://concourse-ci.org/). It's a really slick
CI/CD system that hits every single one of the above requirements. As time went
on however -- whether due to warts of my own doing in defining the config or
Concourse's own quirks -- I kept running into more & more issues with actually
running the services. Sometimes it was shoddy container DNS that wasn't agreeing
with `netsvc` (or with anything), sometimes the service would say it was running
but it really wasn't, sometimes the worker nodes would just... hang for no
reason.

Overall, Concourse is super cool and I have nothing bad to say about it; I just
ran into too many issues with the rest of our stack, and in a fit of frustration
I decided to completely rip it out and choose a new platform. And that platform
is...

***Jenkins!*** It's not a change I ever thought I would make, but hear me out.
Jenkins is a *very* mature piece of software, and its failure modes are very
well-understood in general. In addition, Jenkins is probably the most common
CI/CD system you will run into in enterprise environments (based on my own
experiences in consulting). Those two points alone serve the broader mission of
OSC as a teaching instrument.

But, Jenkins also has some very nice things about it -- it has a very
straightforward (albeit "ugly", which I like) web interface, it has a *wealth*
of plugins available to extend its functionality to varying degrees, jobs can be
run in any kind of environment you configure (Concourse requires all jobs to run
in containers), and more that I don't have time to list. Additionally, Jenkins
has a "configuration as code" plugin that can not only itself be installed
automatically, but can then be used to define an entire Jenkins server exactly
the way you want it to be -- as, well, code. It's something I've known about for
many years but never actually had the chance to try out myself, and I'm
pleasantly surprised that it works as well as it does.

The only real thing keeping me away from Jenkins initially was that it's written
in Java. But I've used Jenkins ***a lot*** in my career, and despite it's not
just Stockholm syndrome -- it's a tool that does its job well, and has been
doing that job for a very long time.

So, that changeover is now complete, and Jenkins is the new underpinning of the
`cicd` subsystem. I'm unironically excited for the change, because it's already
working even better than I'd hoped. Check it out!

### `rhad` changes

As mentioned before, `rhad` is our CI/CD lifecycle management tool. It started
out as a linter aggregator, and then we decided to roll in full lifecycle
functionality once we revisited it (and rewrote it in Go, from Bash).

However, despite `rhad` having a generally straightforward way to add new
linters to the codebase, that list would inevitably grow to be ***absolutely
massive*** to support every possible linting case. We only had 6 or so linters,
and already the copy-paste was showing. Plus, we'd need to understand every
single linter in use, how to manage around them, how to provide them config
overrides, etc. This was already tiresome once I introduced `staticcheck` for Go
linting, and when I tried to add `tflint` for Terraform (you wanna talk about a
frustrating tool to configure, WOOO BOY).

So, I took a step back and revisited [GitHub's
Super-Linter](https://github.com/github/super-linter) for inspiration -- and
walked away deciding to just use Super-Linter for most of our linting purposes
at OSC. The developers of that tool have already done all the hard work solving
for cases mentioned above & more, plus it has ***so many included linters*** out
of the box.

However, Super-Linter currently has some outstanding bugs that prevent a
successful linting experience -- specifically, a bug where [`golangci-lint`
isn't configured correctly to scan packages vs. individual
files](https://github.com/github/super-linter/issues/1599), and so fails
incorrectly when linting a Go repo.

So, what I've decided to do is first run Super-Linter with buggy languages
disabled, and then run `rhad`'s linters for those cases that we can better
control. At this moment, this feels like a good compromise, doesn't require me
to trash all the linter code I've already written, and still allows for `rhad`
to be used as a lifecycle manager post-lint.

### Website updates

It turns out that Hugo (the tool used to build the site content you're reading)
automatically generates an RSS feed for any of its "list" page types. The
landing page for blog posts here is a "list" page. Hence, we have an RSS feed!
It's the link under the orange RSS buttons at the top of the blog landing page,
and at the bottom of *this* page.

If you're not into feeds, that's ok! Stay tuned on the usual social media feeds
for OSC updates. If you *are* into feeds, and are about to come after me about
putting out an RSS feed and not an Atom feed, bear with me! Generating an Atom
feed is more work than RSS using Hugo, because RSS comes for free out of the
box. But it's on the roadmap!

### Miscellaneous

* For some time, at least one of the subsystem `Vagrantfiles` has specified a
  larger-than-default disk size at runtime. This works as expected (in that a
  disk that was 10GB would then be 20GB), but I was reminded the hard way that
  ***simply increasing disk size does not increase available space***. You need
  to also expand a partition to fill the new unused space, and then resize the
  filesystem to take advantage of it. I noticed this when trying to pull very
  large container images for `cicd` jobs, and the disk filled up when it
  "shouldn't have". So, in `imgbuilder/scripts/run/main.sh`, there is now a
  block that handles the two-step process of expanding the root partition and
  resizing the filesystem. It's working as expected, which is a pleasant
  surprise.

* `osc-infra` will start having tagged releases from now on. This should help
  others (and myself, frankly) keep cleaner track of when each commit was marked
  as a) buildable and b) of some notable milestone.

* I've started to notice some really frustrating issues in the local VM
  bootstrapper where builds fail for spraodic (but usually consistent) reasons.
  Most commonly, it's when trying to do and `apt-get update` anywhere at all, at
  any stage of the lifecycle install. I first noticed it when trying to install
  Docker on a subsystem node (via
  `osc-infra/configmgmt/salt/salt/_common/docker.sls`) and APT errors out when
  it finds a hash sum mismatch for the Debian package repo. It's only sometimes,
  but it's been so far very difficult to isolate (let alone reproduce) because
  the acual repo that throws the mismatch error changes -- sometimes it's the
  `containerd.io` repo, sometimes it's the `docker-ce-rootless-extras` one, etc.
  This has been making progress & iteration on `ociregistry` and `cicd` very
  slow, since both subsystems use Docker. But additionally, this has started to
  happen during the base-image builds too (without even running Salt calls). I'm
  beginning to suspect it's either a Debian repo, OR VirtualBox bug. Both of
  which are not easy fixes on my end.

  * UPDATE: restarting my host machine seems to have gotten things to run
    smoothly again at least once -- so I'm now wondering if it is indeed a
    VirtualBox bug.

  Another strange error is a sporadic failure during the internal TLS cert
  generation (`osc-infra/configmgmt/salt/salt/_common/internal_tls_certs.sls`).
  Sometimes (with no discernable pattern), the whole script fails at the very
  first command (`openssl genrsa ...`). I can't yet tell if it's a race
  condition thing, or something more opaque, but that's another thing to explore
  more.
