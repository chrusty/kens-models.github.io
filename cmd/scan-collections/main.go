package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	flickrclient "github.com/chrusty/kens-models.github.io/internal/flickr-client"
)

const (
	jekyllCollectionPathPrefix = "./collections/_"
	modelTemplateFile          = "model.tmpl"
)

var (
	flickrAPIKey = os.Getenv("FLICKR_APIKEY")
	flickrSecret = os.Getenv("FLICKR_SECRET")
	flickrUserID = os.Getenv("FLICKR_USERID")
)

type model struct {
	FlickrSetID  string
	Description  string
	Title        string
	ThumbnailURL string
}

func main() {

	// Make a new Flickr client:
	client := flickrclient.New(flickrAPIKey).WithSecret(flickrSecret).WithUserId(flickrUserID)

	// Retrieve collections:
	collectionsResponse, err := client.CollectionsGetTree()
	if err != nil {
		log.Fatalf("Error retrieving collections: %s", err.Error())
	}
	log.Printf("Found %d collections:", len(collectionsResponse.Collections.Collection))

	// Retrieve photosets:
	photosetsResponse, err := client.PhotosetsGetList()
	if err != nil {
		log.Fatalf("Error retrieving photosets: %s", err.Error())
	}
	log.Printf("Found %d photosets:", len(photosetsResponse.Photosets.Photoset))

	// Go through the collections:
	for _, collection := range collectionsResponse.Collections.Collection {
		jekyllCollectionPath := fmt.Sprintf("%s%s", jekyllCollectionPathPrefix, strings.ReplaceAll(strings.ToLower(collection.Title), " ", ""))
		log.Printf(" * %s => %s", collection.Title, jekyllCollectionPath)

		// Go through each set (album) we find:
		for _, set := range collection.Set {
			modelFileName := strings.ReplaceAll(strings.ToLower(set.Title), " ", "-")
			modelFilePath := fmt.Sprintf("%s/%s.md", jekyllCollectionPath, modelFileName)

			// Put together a model (which we can run the template on):
			model := model{
				Description: set.Description,
				FlickrSetID: set.ID,
				Title:       set.Title,
			}

			// See if we have any photo URLs in the photosets we found:
			for _, foundPhotoset := range photosetsResponse.Photosets.Photoset {
				if foundPhotoset.ID == set.ID {
					model.ThumbnailURL = foundPhotoset.PrimaryPhotoExtras.UrlSmall
				}
			}

			log.Printf("   * Model (\"%s\") [ID=%s, Thumbnail=%v] => %s", set.Title, set.ID, (model.ThumbnailURL != ""), modelFilePath)

			// Render the model template:
			tmpl, err := template.New(modelTemplateFile).ParseFiles(modelTemplateFile)
			if err != nil {
				log.Fatalf("Error parsing model template: %s", err.Error())
			}

			// Make a file for the album:
			modelFile, err := os.Create(modelFilePath)
			if err != nil {
				log.Fatalf("Error opening modle file: %s", err.Error())
			}

			// Execute the template:
			if err := tmpl.Execute(modelFile, model); err != nil {
				log.Fatalf("Error executing model template: %s", err.Error())
			}

			// Close the file:
			modelFile.Close()
		}
	}
}
