---
layout: post
title:  "Announcing OpenSourceCorp!"
date:   2021-07-05 23:50:00 -0500
# categories: jekyll update
author: Ryan J. Price
---

Today, I'm very excited to finally announce the public launch of
[OpenSourceCorp](https://github.com/opensourcecorp) -- the free and open-source
enterprise. This is something I've been working towards for a few months now,
and while it's been public the entire time, it's finally in a spot worth sharing
with the rest of the world.

My contributions to the open-source community all started with [an ndJSON
logging framework for R](https://github.com/ryapric/loggit) that I wrote at
work, to solve a very specific set of challenges I was having. I figured that
even if the tool itself wasn't useful to anyone but me, then at least the
codebase could be helpful to look through as an example of how a simple R
package should be designed, since I was reading out of [the book on the
subject](https://r-pkgs.org/).

Since then, I've been involved in [various open-source
activities](https://github.com/ryapric) over the past several years, from your
typical [mountain of mishmash
repos](https://github.com/ryapric?tab=repositories), to minor contributions to
larger open-source projects (wow this summary is hard to find a GH link for), to
[less-personalized side projects](https://github.com/ops-utils) intended for
larger community engagement, and [talks that I've
given](https://github.com/ryapric/talks) to the various communities in which I
work. Those communities now span a pretty wide range of programming languages,
software & data platforms, automation solutions, and more. I'm incredibly
grateful for all of the learning opportunities I have had to embrace such a
range of experiences, and have grown to appreciate really quality learning
content for skills like that.

Most of my public work has been in pursuit of sharing knowledge, as it had been
shared with me as I was learning these various tools, technologies, coding
styles, etc. But a career in consulting for different, diverse companies has led
me to a realization: you can read all the docs, and run through all the
tutorials & workshops in the world, and you *still* might not be able to
actually ***do the f@ckin' thing*** in "real life". Maybe you'll be an "expert"
on all the ins, outs, and edge behaviors of some complex tool or programming
language, but when it comes time to put your skills into a production
environment, you might end up... stranded.

>Wtf is this 'corporate reverse proxy' I have to register with? And what's this
'Vault' thing where I'm supposed to store passwords in? And what do you MEAN I
can't pull this public image and `kubectl apply` my sweet Kubernetes manifest to
the prod cluster?!"

(You can read more (and growing) motivation for the "why's" in [OpenSourceCorp's
main repo](https://github.com/opensourcecorp/opensourcecorp))

So much of this "how do I actually do this" knowledge is gated behind employers'
walls. Which is to be expected; companies don't want to share what they've got
going on internally, whether that's to protect data, trade secrets, operational
advantages, etc. But this means that if someone wants to learn a certain (or
alternative) way to architect or design something at work, they have fewer and
less-available resources with which to engage.

To be clear, this is not a plea for software maintainers to broaden their
documentation's scope so much that these kinds of things are covered -- you'd
never catch all the use cases! But ultimately, tool & language documentation
alone inevitably will leave hope for enterprise integration wanting.

***This*** is the ultimate motivation for OpenSourceCorp: to be a place where
developers, engineers, sysadmins, designers, and all the other walks of
work-life can work together to deliver real enterprise solutions. Knowledge is
power, and our mission is to bring more of this power to our industry -- a
mission we pursue with one simple, unwavering promise: to do ***every bit of our
work 100% in the open***. Development, discussion, hosting, and anything else
(that respects the privacy or personal security of our people) will be fully
visible to anyone in the world.

So... wait, what's actually going on, here, then?
-------------------------------------------------

What you can expect us to be working on here at OpenSourceCorp is... well, lots
of stuff! The groundwork will be platform solutions like an [image builder
framework](https://github.com/opensourcecorp/ymir), [deployment tooling for our
CI/CD](https://github.com/opensourcecorp/gnar) and
[SCM](https://github.com/opensourcecorp/sauce) systems, and others. Once we have
some platforms in place, we'll start to deploy *actual* software onto those
platforms, integrating them into the broader OSC ecosystem as it grows and
matures. The dream is that if someone ever wants to see how to hook up my
company's source code management to a CI/CD pipeline, have it fetch its secrets,
deploy the application to some virtualization platform, and monitor its
lifecycle, then they can find out how all that fits together across the relevant
codebases across OpenSourceCorp. We'll be leaning into comprehensive
documentation pretty heavily, to make this traversal easier. And if you can't
find specifically what you're looking for, ask us! Like we said above,
everything we do is 100% public.

Someday, we even hope to have *other* open-source projects do work on our
platforms, run their tests on our CI/CD deployment, and deploy their public
content to our infrastructure solutions. We hope to have others work with us
across oceans of target audiences to help share in a common vision to make the
hardware/software industry a better place for everyone. But, one step at a time
:)

One final thing...
------------------

We need help from YOU! We're "hiring", so to speak. OpenSourceCorp cannot
realize its vision in a vacuum/echo chamber, because then it's just one person's
opinions and vision shaping the content. Our strength will come from the
collaboration of our peers, and the more of those peers we can bring on board to
help realize our shared vision, the better. So if you're interested in working
towards this shared vision of a more open tech world, have something to
contribute to the project, or just want to learn more, we want to hear from you!

So keep an eye out for future developments from us, and we hope to see you all
around the office! ;)
