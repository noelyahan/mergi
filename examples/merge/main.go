package main

import (
	"fmt"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"image"
	"log"
)

// This example will guide how to merge multiple images using mergi.Merge API

/*
	Specialty is mergi use merge mechanism called template to line up images
	before merge
*/

func main() {
	img, err := mergi.Import(loader.NewFileImporter("testdata/evraiki-2514543_240_180.jpg"))
	if err != nil {
		log.Fatal(err)
	}

	tigerLarge, err := mergi.Import(loader.NewFileImporter("testdata/tiger-2320819_960_720.jpg"))
	if err != nil {
		log.Fatal(err)
	}
	rabbitLarge, err := mergi.Import(loader.NewFileImporter("testdata/rabbit-1882699_960_720.jpg"))
	if err != nil {
		log.Fatal(err)
	}

	//	lets try to merge 2 images in horizontal manner
	res := mergeHorizontal2([]image.Image{img, img})
	mergi.Export(loader.NewFileExporter(res, getPath("horizontal_2.png")))

	// now lets try vertical with 2 images
	res = mergeVertical2([]image.Image{img, img})
	mergi.Export(loader.NewFileExporter(res, getPath("vertical_2.png")))

	// now lets increase the image count to 4 and do the same
	res = mergeHorizontal4([]image.Image{img, img, img, img})
	mergi.Export(loader.NewFileExporter(res, getPath("horizontal_4.png")))

	res = mergeVertical4([]image.Image{img, img, img, img})
	mergi.Export(loader.NewFileExporter(res, getPath("vertical_4.png")))

	// now lets explore how 2x2 align looks like
	res = merge2x2([]image.Image{img, img, img, img})
	mergi.Export(loader.NewFileExporter(res, getPath("2x2.png")))

	// lets crop all friends characters
	tiger, _ := mergi.Crop(tigerLarge, image.Pt(0, 0), image.Pt(495, 600))
	rabbit, _ := mergi.Crop(rabbitLarge, image.Pt(350, 0), image.Pt(350, 600))

	// as we discuss in crop section time to bring chandler and monica together :)
	res, err = mergi.Merge("TT", []image.Image{tiger, rabbit})
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(loader.NewFileExporter(res, getPath("tiger_rabbit.png")))

	res, err = mergi.Merge("TT", []image.Image{rabbit, tiger})
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(loader.NewFileExporter(res, getPath("rabbit_tiger.png")))

	// lets build some next level merge template
	// for this I'm gonna get some help from placeholder site
	img500x500, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/500x500"))
	img250x250, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/250x250"))
	img200x200, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/200x200"))
	img200x100, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/200x100"))

	// let's define a template
	template := "TTBTBBTBBBBT"

	// now let's layout images
	imgs := []image.Image{
		img500x500, // TOP
		img250x250, // TOP
		img250x250, // BOTTOM
		img200x200, // TOP
		img200x100, // BOTTOM
		img200x200, // BOTTOM
		img200x100, // TOP
		img200x100, // BOTTOM
		img200x100, // BOTTOM
		img200x100, // BOTTOM
		img200x100, // BOTTOM
		img500x500, // TOP
	}

	res, err = mergi.Merge(template, imgs)
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(loader.NewFileExporter(res, getPath("next_level.png")))
}

func mergeHorizontal2(imgs []image.Image) image.Image {
	// TT -> represent top, top means 2 images will be align both top
	res, err := mergi.Merge("TT", imgs)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func mergeVertical2(imgs []image.Image) image.Image {
	// TB -> represent top, bottom means 1 image will be top and 2nd will be bottom
	res, err := mergi.Merge("TB", imgs)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func mergeHorizontal4(imgs []image.Image) image.Image {
	// TTTT -> top, top, top, top align
	res, err := mergi.Merge("TTTT", imgs)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func mergeVertical4(imgs []image.Image) image.Image {
	// TBTB -> top, bottom, bottom, bottom align, 1 image will be top other 3 will be below to top one
	res, err := mergi.Merge("TBBB", imgs)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func merge2x2(imgs []image.Image) image.Image {
	// TBTB -> top, bottom, top, bottom align, 2 image will be top other 2 will be below to top ones
	res, err := mergi.Merge("TBTB", imgs)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func getPath(name string) string {
	return fmt.Sprintf("examples/merge/res/%s", name)
}
