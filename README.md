Ken's Models site
=================

A Jekyll-powered site to showcase Ken's Models.

- Links to categorised collections of models
- Thumbnails for models come from Flickr
- Each model links to its own Flickr gallery
- Also a blog


How it works (v1)
-----------------

- Custom scraper written which lists all of Ken's Models Flickr collections (includinga list of photosets in each). Also lists all photosets in full (including thumbnail URLs).
- Scraper matches these lists together and generates markdown files for each model under Jekyll's category folders.
- Jekyll build then picks these up and generates content.
- GitHub automatically deploys to pages for us.


Plans for the future
--------------------

Ken loves spreadsheets, so we could control the metadata using Google sheets and its public API. The scraper could use this spreadsheet as the primary source of information, allowing us to convey as much metadata as we need for each model. The sheet would need the ID for each Flickr photoset (so we can discover the thumbnail URLs dynamically).


Local development
-----------------

For local dev, run:

``` docker-compose up ```

Then open http://localhost:4000 in your browser.



Made using the [Creative Theme](http://startbootstrap.com/template-overviews/creative/) for Jekyll.
