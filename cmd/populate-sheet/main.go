package main

import (
	"context"
	"log"
	"os"
	"strings"

	flickrclient "github.com/chrusty/kens-models.github.io/internal/flickr-client"
	sheetsdata "github.com/chrusty/kens-models.github.io/internal/sheets-data"
	"google.golang.org/api/sheets/v4"
)

const (
	sheetName     = "Models"
	spreadsheetId = "1FKv2l9z69RdR3ZIJ5dRmYO1ntTR6Dwnxao0wfSr5DEU"
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

	// Make a new Flickr client:
	flickrClient := flickrclient.New(flickrAPIKey).WithSecret(flickrSecret).WithUserId(flickrUserID)

	// Retrieve collections:
	collectionsResponse, err := flickrClient.CollectionsGetTree()
	if err != nil {
		log.Fatalf("Error retrieving collections: %s", err.Error())
	}
	log.Printf("Found %d collections:", len(collectionsResponse.Collections.Collection))

	// We'll build up new rows here:
	newRows := &sheets.ValueRange{
		Values: [][]interface{}{},
	}

	// Go through the collections:
	for _, collection := range collectionsResponse.Collections.Collection {

		// Go through each set (album) we find:
		for _, set := range collection.Set {

			// Prepare an empty row:
			modelRow := make([]interface{}, 8)

			// Fill in what we can:
			modelRow[0] = modelNumber(set.Title)      // Model number
			modelRow[1] = collection.Title            // Category
			modelRow[2] = modelTitle(set.Title)       // Model title
			modelRow[3] = set.Description             // Description
			modelRow[6] = modelScale(set.Description) // Scale
			modelRow[7] = set.ID                      // Flickr photoset ID

			// Add our new row for this model:
			newRows.Values = append(newRows.Values, modelRow)
		}
	}

	// Add the new rows to our sheet:
	response, err := googleSheetsClient.Service().Spreadsheets.Values.Append(spreadsheetId, "Models", newRows).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(context.TODO()).Do()
	if err != nil || response.HTTPStatusCode != 200 {
		log.Fatalf("Unable to insert sheet data: %s", err.Error())
	}

	log.Printf("Added %d rows to the spreadsheet", len(newRows.Values))
}

func modelNumber(source string) string {

	// Split the source string up into parts:
	parts := strings.Split(source, " ")

	// Split the "M" (if we found it):
	if len(parts) > 1 {
		if withoutPrefix, found := strings.CutPrefix(parts[0], "M"); found {
			return withoutPrefix
		}
	}

	return "n/a"
}

func modelTitle(source string) string {

	// Split the source string up into parts:
	parts := strings.Split(source, " ")

	// Split the "M" (if we found it):
	if len(parts) > 1 {
		if strings.HasPrefix(parts[0], "M") {
			return strings.Join(parts[1:], " ")
		}
	}

	return source
}

func modelScale(source string) string {

	// Look for parts containing a "/" character:
	for _, part := range strings.Split(source, " ") {
		if strings.Contains(part, "/") {
			return part
		}
	}

	return "n/a"
}
