package main

import (
	"fmt"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
	"log"
)

func main() {
	// Let's import some images from local file system
	img, err := mergi.Import(io.NewFileImporter("testdata/tiger-2320819_960_720.jpg"))
	if err != nil {
		log.Fatal(err)
	}
	watermarkImage, err := mergi.Import(io.NewFileImporter("testdata/mergi_logo_watermark.png"))
	if err != nil {
		log.Fatal(err)
	}

	watermarkTopLeft(watermarkImage, img)

	watermarkResizeTopRightWith(watermarkImage, img)

	watermarkEverywhere(watermarkImage, img)

	// Just for fun ;)
	// Lets create a identity card
	identityWatermark()
}

func watermarkTopLeft(watermarkImage, img image.Image) {
	res, err := mergi.Watermark(watermarkImage, img, image.Pt(0, 0))
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(io.NewFileExporter(res, getPath("watermark_1.png")))

}

func watermarkResizeTopRightWith(watermarkImage, img image.Image) {
	b := watermarkImage.Bounds()
	w := b.Max.X / 4
	h := b.Max.Y / 4
	newWatermarkImage, _ := mergi.Resize(watermarkImage, uint(w), uint(h))
	res, err := mergi.Watermark(newWatermarkImage, img, image.Pt(img.Bounds().Max.X-w, 0))
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(io.NewFileExporter(res, getPath("watermark_2.png")))
}

func watermarkEverywhere(watermarkImage, img image.Image) {
	b := watermarkImage.Bounds()
	w := b.Max.X / 4
	h := b.Max.Y / 4
	newWatermarkImage, _ := mergi.Resize(watermarkImage, uint(w), uint(h))

	res := img
	var err error
	//	Lets add watermarks all over the image
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			res, err = mergi.Watermark(newWatermarkImage, res, image.Pt(img.Bounds().Max.X-(w*x), h*y))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	mergi.Export(io.NewFileExporter(res, getPath("watermark_3.png")))

}

func identityWatermark() {
	idCard, _ := mergi.Import(io.NewURLImporter("https://cdn.pixabay.com/photo/2013/07/12/19/03/id-154285_960_720.png"))
	profileImage, _ := mergi.Import(io.NewURLImporter("https://cdn.pixabay.com/photo/2017/08/30/17/27/business-woman-2697954_960_720.jpg"))

	cropedProfile, _ := mergi.Crop(profileImage, image.Pt(300, 50), image.Pt(217, 254))
	finalImage, _ := mergi.Watermark(cropedProfile, idCard, image.Pt(20, 63))
	mergi.Export(io.NewFileExporter(finalImage, getPath("watermark_4.png")))
}

func getPath(name string) string {
	return fmt.Sprintf("examples/watermark/res/%s", name)
}
