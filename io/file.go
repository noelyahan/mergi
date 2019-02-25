package io

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type file struct {
	path string
	img  image.Image
}

// NewFileImporter uses to import jpg, png images from given path
func NewFileImporter(path string) Importer {
	return file{path, nil}
}

// NewFileExporter uses to export jpg, png images to given path
func NewFileExporter(img image.Image, path string) Exporter {
	return file{path, img}
}

func (o file) Import() (image.Image, error) {
	errMsg := "Mergi cannot read or decode !"
	ext := getExt(o.path)
	f, err := os.Open(o.path)
	if err != nil {
		return nil, errors.New(errMsg)
	}

	var img image.Image
	if ext == "jpg" || ext == "jpeg" {
		img, err = jpeg.Decode(f)
	} else if ext == "png" {
		img, err = png.Decode(f)
	} else {
		return nil, errors.New(errMsg)
	}

	if err != nil {
		return nil, errors.New(errMsg)
	}
	return img, nil
}

func (o file) Export() error {
	img := o.img
	if img == nil {
		return errors.New("Mergi found a invalid file ")
	}

	ext := getExt(o.path)
	f, err := os.Create(o.path)
	if err != nil {
		msg := fmt.Sprintf("Sorry Mergi failed to create: %s", o.path)
		log.Printf(msg)
		return errors.New(msg)
	}
	if ext == "jpg" || ext == "jpeg" {
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	} else if ext == "png" {
		err = png.Encode(f, img)
	}
	if err != nil {
		msg := fmt.Sprintf("Sorry Mergi cannot encode the image: %s", o.path)
		return errors.New(msg)
	}
	defer f.Close()
	return nil
}
