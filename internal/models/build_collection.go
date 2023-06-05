package models

import (
	"fmt"
	"strings"
)

type BuildCollection struct {
	Description string
	IconSmall   string
	ID          string
	Title       string
}

func (m *BuildCollection) FileName(extension string) string {
	return fmt.Sprintf("%s.%s", m.Label(), extension)
}

func (m *BuildCollection) Label() string {
	return strings.ReplaceAll(strings.ToLower(m.Title), " ", "-")
}
