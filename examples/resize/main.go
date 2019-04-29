package main

import (
	"fmt"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/eximp"
	"log"
)

// This example will guide how to resize a image using mergi.Resize API
func main() {
	img, err := mergi.Import(eximp.NewFileImporter("testdata/coffee-171653_960_720.jpg"))
	if err != nil {
		fmt.Println(err)
		return
	}
	scale := 2

	// lets scale up
	newWidth := uint(img.Bounds().Max.X * scale)
	newHeight := uint(img.Bounds().Max.Y * scale)
	res, err := mergi.Resize(img, newWidth, newHeight)
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(eximp.NewFileExporter(res, getPath("scale_up.png")))

	// lets scale down
	newWidth = uint(img.Bounds().Max.X / scale)
	newHeight = uint(img.Bounds().Max.Y / scale)
	res, err = mergi.Resize(img, newWidth, newHeight)
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(eximp.NewFileExporter(res, getPath("scale_down.png")))

}

func getPath(name string) string {
	return fmt.Sprintf("examples/resize/res/%s", name)
}
