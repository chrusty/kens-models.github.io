package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	docsClient "github.com/chrusty/kens-models.github.io/internal/google-clients/docs"
	driveClient "github.com/chrusty/kens-models.github.io/internal/google-clients/drive"
	"github.com/chrusty/kens-models.github.io/internal/models"
	"github.com/erikh/gdocs-export/pkg/converters"
	"github.com/erikh/gdocs-export/pkg/util"
)

const (
	defaultBlogAuthor = "Ken"
	defaultBlogIcon   = "fa-clock"
	jekyllPostsPath   = "./collections/_posts"
	templateFileName  = "blog.tmpl"
	templatePath      = "./internal/templates/"
)

var (
	googleAuthKey  = os.Getenv("GOOGLE_AUTH_KEY")
	googleFolderId = os.Getenv("GOOGLE_FOLDER_ID")
)

func main() {

	// Make a new Google drive client:
	driveClient, err := driveClient.New().WithAPIKey(googleAuthKey)
	if err != nil {
		log.Fatalf("Unable to prepare a Google Drive client: %s", err.Error())
	}

	// Make a new Google docs client:
	docsClient, err := docsClient.New().WithAPIKey(googleAuthKey)
	if err != nil {
		log.Fatalf("Unable to prepare a Google Docs client: %s", err.Error())
	}

	// Retrieve a list of docs in the folder:
	driveResponse, err := driveClient.Service().Files.List().Q("mimeType='application/vnd.google-apps.document'").MaxResults(1000).Do()
	if err != nil {
		log.Fatalf("Unable to list docs: %s", err.Error())
	}

	// Process each doc:
	for _, doc := range driveResponse.Items {

		// Get the document:
		docsResponse, err := docsClient.Service().Documents.Get(doc.Id).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve doc: %s", err.Error())
		}

		// Download the assets:
		manifest, err := util.DownloadAssets(docsClient.HTTP(), docsResponse, "img/blog", true)
		if err != nil {
			log.Fatalf("Unable to download assets")
		}

		// Convert to markdown:
		convertedDoc, err := converters.Convert("md", docsResponse, manifest)
		if err != nil {
			log.Fatalf("Unable to convert document: %s", err.Error())
		}

		// Parse the creation date of the Google Drive file:
		createdTimeStamp, err := time.Parse(time.RFC3339, doc.CreatedDate)
		if err != nil {
			log.Fatalf("Unable to parse CreatedDate from drive file: %s", err.Error())
		}

		// Turn it into a blog post:
		blogPost := &models.BlogPost{
			Author:  defaultBlogAuthor,
			Content: convertedDoc,
			Icon:    defaultBlogIcon,
			Time:    createdTimeStamp,
			Title:   docsResponse.Title,
		}

		// Parse the doc description to get the publishing metadata:
		if err := blogPost.ParseDescription(doc.Description); err != nil {
			log.Fatalf("Unable to parse description: %s", err.Error())
		}

		// Render the blog template:
		tmpl, err := template.New(templateFileName).ParseFiles(fmt.Sprintf("%s%s", templatePath, templateFileName))
		if err != nil {
			log.Fatalf("Error parsing model template: %s", err.Error())
		}

		// Make a file for the blog post:
		blogPostFile, err := os.Create(blogPost.FileName(jekyllPostsPath, "md"))
		if err != nil {
			log.Fatalf("Error opening modle file: %s", err.Error())
		}

		// Execute the template:
		if err := tmpl.Execute(blogPostFile, blogPost); err != nil {
			log.Fatalf("Error executing blog template: %s", err.Error())
		}

		// Close the file:
		blogPostFile.Close()

		log.Printf(" * Doc (\"%s\") [ID=%s] => %s", docsResponse.Title, doc.Id, blogPost.FileName(jekyllPostsPath, "md"))
	}
}
