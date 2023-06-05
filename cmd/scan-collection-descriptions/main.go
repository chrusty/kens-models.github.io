package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	flickrclient "github.com/chrusty/kens-models.github.io/internal/flickr-client"
	"github.com/chrusty/kens-models.github.io/internal/models"
)

const (
	templateFileName = "build-collection.tmpl"
	templatePath     = "./internal/templates/"
)

var (
	flickrAPIKey = os.Getenv("FLICKR_APIKEY")
	flickrSecret = os.Getenv("FLICKR_SECRET")
	flickrUserID = os.Getenv("FLICKR_USERID")
)

func main() {

	// Make a new Flickr client:
	flickrClient := flickrclient.New(flickrAPIKey).WithSecret(flickrSecret).WithUserId(flickrUserID)

	// Retrieve collections:
	collectionsResponse, err := flickrClient.CollectionsGetTree()
	if err != nil {
		log.Fatalf("Error retrieving collections: %s", err.Error())
	}
	log.Printf("Found %d collections:", len(collectionsResponse.Collections.Collection))

	// Go through the collections:
	for _, collection := range collectionsResponse.Collections.Collection {

		// Put together a model (which we can run the template on):
		buildCollection := &models.BuildCollection{
			Description: collection.Description,
			IconSmall:   collection.IconSmall,
			ID:          collection.ID,
			Title:       collection.Title,
		}

		// Render the model template:
		tmpl, err := template.New(templateFileName).ParseFiles(fmt.Sprintf("%s%s", templatePath, templateFileName))
		if err != nil {
			log.Fatalf("Error parsing model template: %s", err.Error())
		}

		// Make a file for the collection:
		collectionFile, err := os.Create(buildCollection.FileName("html"))
		if err != nil {
			log.Fatalf("Error opening build-collection file: %s", err.Error())
		}

		// Execute the template:
		if err := tmpl.Execute(collectionFile, buildCollection); err != nil {
			log.Fatalf("Error executing build-collection template: %s", err.Error())
		}

		// Close the file:
		collectionFile.Close()

		log.Printf(" * %s => %s", collection.Title, buildCollection.FileName("html"))
	}
}
