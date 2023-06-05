---
layout: post
title:  "Building the site"
author: Chris
time: "Wednesday 24th of May 2023"
icon: fa-gears
---

We're currently working to bring the site online. Please bear with us!

<!--more-->

The tech stack
--------------

For those that are interested:

* Site is powered by Jekyll
* Running on GitHub pages
* How the site works:
    - Jekyll "collections" for build categories (ships / aircraft / rail / buildings / cars-and-trucks / military vehicles)
    - Blog section powered by Jekyll "posts" for Ken to post project updates / tools-and-techniques articles    
    - Ken maintains the Flickr content, managing collections and albums on that side
    - A DB models is managed by a Google spreadsheet, which contains metadata, content, and links to Flickr
    - A custom script reads this spreadsheet and creates an item for each model which Jekyll reads upon deployment


### Why do it this way?

Yes this could all be wordpress or something, but the benefits of doing it this way are:

* We don't have to pay for Wordpress
* We don't have to deal with all of the cruft of wordpress
* We don't need a DB to serve this content as we get Jekyll to turn it into a static site whenever we need it to
* Ken loves spreadsheets!
