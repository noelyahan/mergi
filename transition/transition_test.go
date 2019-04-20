package transition

import (
	"testing"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
)

func getImages() (img1 image.Image, img2 image.Image){
	img1, _ = mergi.Import(io.NewFileImporter("../testdata/nature-3042751_960_720.jpg"))
	img1, _ = mergi.Resize(img1, uint(img1.Bounds().Max.X / 4), uint(img1.Bounds().Max.Y / 4))

	img2, _ = mergi.Import(io.NewFileImporter("../testdata/soldier-708711_960_720.jpg"))
	img2, _ = mergi.Resize(img2, uint(img2.Bounds().Max.X / 4), uint(img2.Bounds().Max.Y / 4))

	img1, _ = mergi.Crop(img1, image.ZP, image.Pt(img2.Bounds().Max.X, img2.Bounds().Max.Y))
	return
}

func TestFadeIn(t *testing.T) {
	img1, img2 := getImages()
	res := FadeIn(img1, img2, 1, 0.05)

	gif, _ := mergi.Animate(res, 10)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}

func TestFadeOut(t *testing.T) {
	img1, img2 := getImages()
	res := FadeOut(img1, img2, 1, 0.05)

	gif, _ := mergi.Animate(res, 10)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}