package transition

import (
	"testing"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
	"github.com/noelyahan/mergitrans"
)

func getImages() (img1 image.Image, img2 image.Image) {
	scale := 4
	img1, _ = mergi.Import(io.NewFileImporter("../testdata/nature-3042751_960_720.jpg"))
	img1, _ = mergi.Resize(img1, uint(img1.Bounds().Max.X/scale), uint(img1.Bounds().Max.Y/scale))

	img2, _ = mergi.Import(io.NewFileImporter("../testdata/soldier-708711_960_720.jpg"))
	img2, _ = mergi.Resize(img2, uint(img2.Bounds().Max.X/scale), uint(img2.Bounds().Max.Y/scale))

	img1, _ = mergi.Crop(img1, image.ZP, image.Pt(img2.Bounds().Max.X, img2.Bounds().Max.Y))
	return
}

func TestImage(t *testing.T) {
	img1, img2 := getImages()

	frames := Image(img1, img2, mergitrans.Ink2, mergi.MaskBlack, 0, float64(len(mergitrans.Ink2)-1) , 1)

	gif, _ := mergi.Animate(frames, 1)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}


func TestImages(t *testing.T) {
	frames := Images(mergitrans.Videos.VideoPoppyField, mergitrans.Videos.VideoClouds, mergitrans.Ink1, mergi.MaskBlack, 0, float64(len(mergitrans.Ink1)-1), 1)

	gif, _ := mergi.Animate(frames, 1)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}
