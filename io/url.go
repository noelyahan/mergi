package io

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"net/url"
	"strings"
)

type httpURL struct {
	url string
}

// NewURLImporter uses to import jpg, png images from given url
func NewURLImporter(url string) Importer {
	return httpURL{url}
}

func (o httpURL) Import() (image.Image, error) {
	errMsg := fmt.Sprintf("Mergi can't find the file or file is invalid %s", o.url)
	if !isValidURL(o.url) {
		return nil, fmt.Errorf("Mergi found a invalid URL: %s", o.url)
	}
	resp, err := http.Get(o.url)
	if err != nil {
		return nil, errors.New(errMsg)
	}
	reader := resp.Body
	cType := resp.Header.Get("Content-Type")
	cType = strings.Replace(cType, "/", ".", -1)
	ext := getExt(cType)

	var img image.Image
	if ext == "jpg" || ext == "jpeg" {
		img, err = jpeg.Decode(reader)
	} else if ext == "png" {
		img, err = png.Decode(reader)
	} else {
		return nil, errors.New(errMsg)
	}

	if err != nil {
		return nil, errors.New(errMsg)
	}
	return img, nil
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true
}
