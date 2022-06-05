---
title: 'OpenSourceCorp - The Free and Open-Source Enterprise'
---

---

Here at OpenSourceCorp, we strive to provide an opinionated perspective on how
software systems work "in real life". We don't have some grand monetization
vision for ourselves, nor do we propose anything particularly novel in *what* we
do. What we do is collaborate to write application code that runs on
infrastructure deployed via CI/CD pipelines, secured via applicable industry
best-practices, and supported by comprehensive documentation.

The novel component that we provide is that we do all of this work ***100% in
the open***.

We do this because we have developed a general frustration with ***how*** modern
teams are expected to learn different things -- especially how so many tools,
technologies, and processes are used in concert with each other in a broader
enterprise ecosystem. There are plenty of tutorials and "getting started" guides
for any number of tools or topics, but while several of these have supplements
for their use in a larger ecosystem, they've never feel like quite enough to
document the holistic lifecycle of a use case. The most valuable learning
material for *real* use cases happens within company walls with propietary
software -- where if you're not an employee, you don't get to learn.

Take any given, relatively-complex software, like [Apache
Kafka](https://kafka.apache.org/) or [HashiCorp
Vault](https://www.vaultproject.io/docs). Their respective documentation portals
are very comprehensive, and provide a wealth of use cases, simple & complex
configurations, and more.

But what if you're self-hosting a production-quality cluster at your company?
How do you *actually* deploy that cluster, and not rely on manually launching it
via the (usually very good) installation instructions? Do you bootstrap some EC2
machines in a CloudFormation template for an AWS deployment? Do you do that
bootstrapping with raw shell code, [Ansible](https://docs.ansible.com/), or
something else?

Or maybe you stuff containers into [Kubernetes](https://kubernetes.io/docs/)?
The docs for Vault, for example, provide a [Helm
chart](https://www.vaultproject.io/docs/platform/k8s/helm), sure -- but what if
your team/company doesn't use Helm?

Or, isn't ***allowed*** to use Helm? We haven't even touched on the security
considerations that your company might take against any tools, and those
considerations vary wildly for the *same* tool across orgs. A version bump for a
Java library might be a few hours of testing & validation in some orgs, but
(literally) *years* of security review in others, *before* you're even allowed
to start doing those same upgrade validations.

All of this begs the same question again -- ***how do you actually work with
this stuff in real life?***

These, and ever more questions, are what OpenSourceCorp aims to provide answers
& perspective for. Real code, real solutions, for real enterprise challenges.

If you'd like to get involved, reach out to us!

---

### Site Roadmap

* Wiki section for the site
* "About" page (which might partially be *this* page's contents moved)
* Core values
* `osc-infra` arch diagram
* Contribution/onboarding model
* Target platform commentary -- "why not (strictly) Kubernetes/AWS/etc.", etc.
  Maybe that's part of core values though.
* Partnership model (i.e. how do we get other OSS projects to USE our platforms,
  to help enhance the value of doing everything in the open)
