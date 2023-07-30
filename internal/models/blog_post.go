package models

import (
	"fmt"
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
