OpenSourceCorp website content
==============================

This repo houses all of the content on OpenSourceCorp's public website.

Development (in progress)
-------------------------

To add a new page, just add a new Go HTML template. If your page needs render
data, add it to (or create) a `_data.json` file in the same directory as the
page.

To add a new *section*, add a file of identical formatting as the `_data.json`
files, but name them `_index.json` instead. This tells the renderer to start with
a site-consistent index file at that section's root.

The renderer processes `_index.json` files, and then adds additional data via
discovered `_data.json` files. There is currently no name key conflict checking,
so render data in `_data.json` effectively overrides what's found in
`_index.json`.
