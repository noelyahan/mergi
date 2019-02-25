package main

import (
	"fmt"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
	"log"
)

// This example will guide how to crop a image using mergi.Crop API
func main() {
	// Let's open up friends_group image
	tigerLarge, err := mergi.Import(io.NewFileImporter("testdata/tiger-2320819_960_720.jpg"))
	if err != nil {
		log.Fatal(err)
	}
	rabbitLarge, err := mergi.Import(io.NewFileImporter("testdata/rabbit-1882699_960_720.jpg"))
	if err != nil {
		log.Fatal(err)
	}

	tigerSize := image.Pt(495, 600)

	// lets crop tiger image :)
	tiger, err := mergi.Crop(tigerLarge, image.Pt(0, 0), tigerSize)
	if err != nil {
		fmt.Printf("Cannot continue the crop [%v]", err)
	}

	// lets save tiger
	mergi.Export(io.NewFileExporter(tiger, "examples/crop/res/tiger.png"))

	rabbitSize := image.Pt(350, 600)

	// lets crop rabbit too
	rabbit, err := mergi.Crop(rabbitLarge, image.Pt(350, 0), rabbitSize)
	if err != nil {
		fmt.Printf("Cannot continue the crop [%v]", err)
	}

	// lets save rabbit too
	mergi.Export(io.NewFileExporter(rabbit, "examples/crop/res/rabbit.png"))

	// Let's see how we can merge tiger and rabbit together :) in merge example
}
