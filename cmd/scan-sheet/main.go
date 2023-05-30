package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	flickrclient "github.com/chrusty/kens-models.github.io/internal/flickr-client"
	"github.com/chrusty/kens-models.github.io/internal/models"
	sheetsdata "github.com/chrusty/kens-models.github.io/internal/sheets-data"
)

const (
	headerBuildHours           = "Build Hours"
	headerBuildPeriod          = "Build Period"
	headerCategory             = "Category"
	headerComments             = "Comments"
	headerCompletionDate       = "Compl. Date"
	headerModelID              = "Model ID"
	headerName                 = "Name"
	headerFlickrSetID          = "Flickr Photoset ID"
	headerScale                = "Scale"
	headerSummary              = "Subject Full Title"
	jekyllCollectionPathPrefix = "./collections/_"
	modelTemplateFile          = "model.tmpl"
)

var (
	flickrAPIKey        = os.Getenv("FLICKR_APIKEY")
	flickrSecret        = os.Getenv("FLICKR_SECRET")
	flickrUserID        = os.Getenv("FLICKR_USERID")
	googleAuthKey       = os.Getenv("GOOGLE_AUTH_KEY")
	googleSpreadsheetId = os.Getenv("GOOGLE_SHEET_ID")
)

func main() {

	// Make a new Google sheets client:
	googleSheetsClient, err := sheetsdata.New(googleSpreadsheetId).WithAPIKey(googleAuthKey)
	if err != nil {
		log.Fatalf("Unable to prepare a Google sheets client: %s", err.Error())
	}

	// Retrieve values from the Google sheet:
	values, err := googleSheetsClient.ValuesGet("1:1000", true)
	if err != nil {
		log.Fatalf("Unable to retrieve Google sheets values: %s", err.Error())
	}

	// Build a map of the values so we can search them (by ModelID):
	if err := values.CreateMap(headerModelID); err != nil {
		log.Fatalf("Unable to create map from sheets data: %s", err.Error())
	}

	// Make a new Flickr client:
	flickrClient := flickrclient.New(flickrAPIKey).WithSecret(flickrSecret).WithUserId(flickrUserID)

	// Retrieve photosets:
	photosetsResponse, err := flickrClient.PhotosetsGetList()
	if err != nil {
		log.Fatalf("Error retrieving photosets: %s", err.Error())
	}
	log.Printf("Found %d photosets:", len(photosetsResponse.Photosets.Photoset))

	// Go through each row entry we found:
	for _, entry := range values.Entries {

		// Turn it into an actual model:
		model := models.Model{
			BuildHours:     entry[headerBuildHours],
			BuildPeriod:    entry[headerBuildPeriod],
			Category:       entry[headerCategory],
			Comments:       entry[headerComments],
			CompletionDate: entry[headerCompletionDate],
			FlickrSetID:    entry[headerFlickrSetID],
			ID:             entry[headerModelID],
			Name:           entry[headerName],
			Scale:          entry[headerScale],
			Summary:        entry[headerSummary],
		}

		// See if we have any photo URLs in the photosets we found:
		for _, foundPhotoset := range photosetsResponse.Photosets.Photoset {
			if foundPhotoset.ID == model.FlickrSetID {
				model.ThumbnailURL = foundPhotoset.PrimaryPhotoExtras.UrlSmall
			}
		}

		// Figure out the paths:
		flatCategory := strings.ReplaceAll(strings.ToLower(model.Category), " ", "")
		jekyllCollectionPath := fmt.Sprintf("%s%s", jekyllCollectionPathPrefix, flatCategory)

		log.Printf("   * Model (\"%s\") [ID=%s, Publish=%v] => %s", model.Name, model.ID, model.Publish(), model.FileName(jekyllCollectionPath, "md"))

		if !model.Publish() {
			continue
		}

		// Render the model template:
		tmpl, err := template.New(modelTemplateFile).ParseFiles(modelTemplateFile)
		if err != nil {
			log.Fatalf("Error parsing model template: %s", err.Error())
		}

		// Make a file for the album:
		modelFile, err := os.Create(model.FileName(jekyllCollectionPath, "md"))
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
