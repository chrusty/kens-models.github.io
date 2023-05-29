package models

import (
	"fmt"
	"strings"
)

type Model struct {
	BuildHours     string
	BuildPeriod    string
	Category       string
	Comments       string
	CompletionDate string
	FlickrSetID    string
	ID             string
	Name           string
	Scale          string
	Summary        string
	ThumbnailURL   string
}

func (m *Model) FileName(prefix, extension string) string {
	kebabName := strings.ReplaceAll(strings.ToLower(m.ID), " ", "-")
	return fmt.Sprintf("%s/%s.%s", prefix, kebabName, extension)
}
