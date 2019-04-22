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

func TestCover(t *testing.T) {
	img1, img2 := getImages()
	res := Cover(img1, img2, 10, 0.5)

	gif, _ := mergi.Animate(res, 10)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}

func TestSplit(t *testing.T) {
	img1, img2 := getImages()
	res := Split(img1, img2, 1, 0.03)

	gif, _ := mergi.Animate(res, 10)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}

func TestAll(t *testing.T) {
	img1, img2 := getImages()
	res1 := FadeIn(img2, img1, 1, 0.05)
	res2 := Split(img1, img2, 1, 0.1)
	res3 := Split(img2, img1, 1, 0.1)
	res4 := FadeIn(img1, img2, 1, 0.05)

	//res2 := Cover(img1, img2, 1, 0.05)
	//res3 := FadeIn(img1, img2, 1, 0.05)
	//
	res1 = append(res1, res2...)
	res1 = append(res1, res3...)
	res1 = append(res1, res4...)


	gif, _ := mergi.Animate(res1, 10)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}