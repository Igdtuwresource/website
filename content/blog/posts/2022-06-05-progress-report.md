---
title: "Progress Report 2022-06-05"
publishdate: 2022-06-05T21:20:16-05:00
author: "Ryan J. Price"
---

Not too long between updates this time, but I've got a few neat pieces of news
to share.

* First and foremost, you may recall that I said we'd gotten a domain name
  (`opensourcecorp.org`) in a previous post. For some time now, it's really just
  had some redirect rules to point folks to the GitHub Org, and to the GitHub
  Pages site for blog posts.

  But now, it's a ***[real-ass website!](https://opensourcecorp.org)*** You
  might even be reading this post on it right now!
  
  It's still being served through GH Pages, but in a sort of nonstandard way
  (i.e. no longer using Jekyll natively), with the OSC domain instead of the GHP
  domain format. I'm also using [Hugo](https://gohugo.io) as the static-site
  generator instead -- but not using much of its fancy features; it's pretty
  plain HTML & CSS at the end of the day. Which, is how I'd like it to be.

  The website generator repo is
  [here](https://github.com/opensourcecorp/website), and is being served from
  [here](https://github.com/opensourcecorp/opensourcecorp.github.io) (just like
  it was before).

  I also set us up with [a sick sk8rboi
  logo](https://github.com/opensourcecorp/website/blob/main/static/images/osc-logo.png)!
  It's not the prettiest thing in the world (I'm not an artist by any stretch of
  the imagination), but I thought the "Super-S" style was a meta-nostalgic
  callback to how OSC is intended to be built -- with accessible,
  eventually-familiar tooling.

  The logo is visible on the site's navbar, as well as being its favicon on your
  browser tab. It was designed using a custom Shape on
  [diagrams.net](https://app.diagrams.net), and you can find the XML-y code used
  to create it
  [here](https://github.com/opensourcecorp/website/blob/main/osc-logo.drawio.xml)
  (with embedded comments on how to recreate it).

* Next, you may recall seeing or hearing about the linter aggregator I had
  started as the very first OSC project:
  [`rhadamanthus`](https://github.com/opensourcecorp/rhad). Its name has now
  been shortened to just `rhad`, and has been entirely rewritten from a
  collection of shell scripts (which worked fine, and that I'm still super proud
  of) to a Go CLI utility (which also works fine but will offer more flexibility
  in the future).
  
  I have also moved towards having `rhad` serve as much more than just a linter
  aggregator, but to be a holistic CI/CD solution akin to how a Jenkins shared
  library might work -- the CI/CD subsytem will pick up your repo, and `rhad`
  will process it for all the relevant steps it needs to go from raw idea to
  production deployment. I'm really looking forward to have `rhad` continue to
  expand in functionality as use more cases appear across OSC.

* Finally, I received some helpful feedback about some of the naming conventions
  used within OSC projects. Before all the infra tooling was consolidated into
  [the monorepo](https://github.com/opensourcecorp/osc-infra), each infra
  subsystem had a "cool" name -- `aether`, `faro`, `chonk`, etc. I still think
  they're cool (lol), but they're not doing anything helpful for folks looking
  to get involved with the project.
  
  These subsystems have now been renamed throughout the codebase into more
  meaningful names according to their actual function -- e.g. `configmgmt`,
  `netsvc`, and `datastore`, respectively in the above example. In the longer
  term, this should make the infra codebase more accessible to others.

  You may still find some lingering naming conventions in other repos across the
  GH Org, but those will get cleaned up eventually. Let me know if you find any,
  or have any feedback.

The update that I think is the coolest is obviously the website, so please check
it out and let me know what you think! But before you come after me too hard:
yes, I know it's "ugly" -- it's mostly intentional. But if there's something
you'd like to change, feel free to open a PR to the [website
repo](https://github.com/opensourcecorp/website) with any layout or CSS changes
you have in mind! I especially could use some help finding a (non-JS) way to
make the site more mobile-friendly.
