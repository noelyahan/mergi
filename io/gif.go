package io

import (
	"errors"
	"fmt"
	"image/gif"
	"os"
)

type animation struct {
	path string
	anim gif.GIF
}

// NewAnimationExporter uses to export gif animation to given path
func NewAnimationExporter(anim gif.GIF, path string) Exporter {
	return animation{path, anim}
}

func (o animation) Export() error {
	anim := o.anim

	ext := getExt(o.path)
	f, err := os.Create(o.path)
	if err != nil {
		msg := fmt.Sprintf("Sorry Mergi failed to create: %s", o.path)
		return errors.New(msg)
	}
	if ext == "gif" {
		err = gif.EncodeAll(f, &anim)
	}
	if err != nil {
		msg := fmt.Sprintf("Sorry Mergi cannot encode the image: %s", o.path)
		return errors.New(msg)
	}
	defer f.Close()
	return nil
}
