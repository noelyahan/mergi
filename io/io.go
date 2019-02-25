package io

import (
	"image"
	"strings"
)

// Importer uses to plug any kind of image.Image importer
type Importer interface {
	Import() (image.Image, error)
}

// Exporter uses to plug any kind of image.Image exporter
type Exporter interface {
	Export() error
}

func getExt(p string) string {
	s := strings.Split(p, ".")
	ext := s[len(s)-1]
	if ext == "jpeg" || ext == "jpg" || ext == "png" || ext == "gif" {
		return ext
	}
	return ""
}
