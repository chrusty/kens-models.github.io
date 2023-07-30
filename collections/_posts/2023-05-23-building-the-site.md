---
layout: post
title: Building the site
author: Chris
time: Tuesday 23 of May 2023
icon: fa-hammer
---


We&#39;re currently working to bring the site online. Please bear with us!



\&lt;!--more--&gt;

## The tech stack

For those that are interested:
* The site is powered by Jekyll
* Running on GitHub pages
* How the site works:
  * Jekyll &#34;collections&#34; for build categories (ships / aircraft / rail / buildings / cars-and-trucks / military vehicles)
  * Blog section powered by Jekyll &#34;posts&#34; for Ken to post project updates / tools-and-techniques articles
  * Ken maintains the Flickr content, managing collections and albums on that side
  * A modelsDB is managed by a Google spreadsheet, which contains metadata, content, and links to Flickr
  * A custom script reads this spreadsheet and creates an item for each model which Jekyll reads upon deployment

### Why do it this way?

Yes this could all be wordpress or something, but the benefits of doing it this way are:
* We don&#39;t have to pay for Wordpress
* We don&#39;t have to deal with all of the cruft of wordpress
* We don&#39;t need a DB to serve this content as we get Jekyll to turn it into a static site whenever we need it to
* Ken loves spreadsheets!

