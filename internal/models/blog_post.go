package models

import (
	"fmt"
	"html"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const blogTimeStampFormat = "Monday 02 of January 2006"

type BlogPost struct {
	Author  string `yaml:"author"`
	Content string
	Icon    string    `yaml:"icon"`
	Time    time.Time `yaml:"time"`
	Title   string
}

func (b *BlogPost) FileName(prefix, extension string) string {
	timeStamp := b.Time.Format(time.DateOnly)
	kebabName := strings.ReplaceAll(strings.ToLower(b.Title), " ", "-")
	return fmt.Sprintf("%s/%s-%s.%s", prefix, timeStamp, kebabName, extension)
}

func (b *BlogPost) TimeStamp() string {
	return b.Time.UTC().Format(blogTimeStampFormat)
}

func (b *BlogPost) ParseDescription(description string) error {
	return yaml.Unmarshal([]byte(description), b)
}

func (b *BlogPost) RenderContent() string {

	// Remove leading "\" characters (added by the markdown converter for some reason):
	stripped := strings.ReplaceAll(b.Content, "\n\\", "\n")

	// Add a slash to image paths:
	stripped = strings.ReplaceAll(stripped, "img/blog", "/img/blog")

	// Unescape the stripped content:
	return html.UnescapeString(stripped)
}
