package models

import (
	"fmt"
	"strings"
	"time"
)

const blogTimeStampFormat = "Monday 02 of January 2006"

type BlogPost struct {
	Author  string
	Content string
	Icon    string
	Time    time.Time
	Title   string
}

func (b *BlogPost) FileName(prefix, extension string) string {
	timeStamp := b.Time.Format(time.DateOnly)
	kebabName := strings.ReplaceAll(strings.ToLower(b.Title), " ", "-")
	return fmt.Sprintf("%s/%s-%s.%s", prefix, timeStamp, kebabName, extension)
}

func (b *BlogPost) TimeStamp() string {
	return b.Time.Format(blogTimeStampFormat)
}
