package main

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"image"
	"log"
)

// Mergi supports url to image, Lets try to merge this placeholders
func main() {
	// Opps Sorry ! let's ignore error for now
	img1, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/500x500"))

	img2, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/250x250"))
	img3, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/250x250"))

	img4, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/100x100"))
	img5, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/100x100"))
	img6, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/100x100"))
	img7, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/100x100"))
	img8, _ := mergi.Import(loader.NewURLImporter("https://via.placeholder.com/100x100"))

	imgs := []image.Image{img1, img2, img3, img4, img5, img6, img7, img8}

	res, err := mergi.Merge("TTBTBBBB", imgs)
	if err != nil {
		log.Fatal(err)
	}
	anim, _ := mergi.Animate([]image.Image{img1, img2}, 20)

	// export image file png/jpg
	mergi.Export(loader.NewFileExporter(res, "examples/import_export/res/out.png"))

	// Export animation gif
	mergi.Export(loader.NewAnimationExporter(anim, "examples/import_export/res/out.gif"))
}
