package ease

import (
	"testing"
	"image"
	"github.com/noelyahan/mergi/io"
	"github.com/noelyahan/mergi"
)

func _getImages() (img1 image.Image, img2 image.Image) {
	scale := 4
	img1, _ = mergi.Import(io.NewFileImporter("../testdata/nature-3042751_960_720.jpg"))
	img1, _ = mergi.Resize(img1, uint(img1.Bounds().Max.X/scale), uint(img1.Bounds().Max.Y/scale))

	img2, _ = mergi.Import(io.NewFileImporter("../testdata/soldier-708711_960_720.jpg"))
	img2, _ = mergi.Resize(img2, uint(img2.Bounds().Max.X/scale), uint(img2.Bounds().Max.Y/scale))

	img1, _ = mergi.Crop(img1, image.ZP, image.Pt(img2.Bounds().Max.X, img2.Bounds().Max.Y))
	return
}

func TestEase(t *testing.T) {
	square, _ := mergi.Import(io.NewFileImporter("../testdata/square.jpg"))
	bg, _ := mergi.Import(io.NewFileImporter("../testdata/white_bg.jpg"))

	frames := make([]image.Image, 0)

	to := bg.Bounds().Max.X - square.Bounds().Max.X
	posY := bg.Bounds().Max.Y/2 - square.Bounds().Max.Y/2
	speed := 4

	for i := 0; i < to; i += speed {
		v := Ease(float64(i), float64(to), 0, InOutSine)
		img, _ := mergi.Watermark(square, bg, image.Pt(int(v), posY))
		frames = append(frames, img)
	}

	gif, _ := mergi.Animate(frames, 1)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}

func TestEaseFont(t *testing.T) {
	img1, _ := _getImages()
	frames := make([]image.Image, 0)
	text, _ := mergi.Import(io.NewFileImporter("../testdata/mergi_text.jpg"))
	maskImg, _ := mergi.Import(io.NewFileImporter("../testdata/white.jpg"))
	//maskImg, _ = mergi.Opacity(maskImg, 0)
	text, _ = mergi.Resize(text, uint(img1.Bounds().Max.X)/2, uint(img1.Bounds().Max.Y)/2)
	maskImg, _ = mergi.Resize(text, uint(text.Bounds().Max.X), uint(text.Bounds().Max.Y))

	text, _ = mergi.Mask(text, maskImg, mergi.MaskBlack)
	//text, _ = mergi.MaskColor(text, mergi.MaskBlack, mergi.MaskWhite)
	to := img1.Bounds().Max.X / 3
	for i := 0; i < to; i += 5 {
		v := Ease(float64(i), float64(to), 0, InElastic)
		img, _ := mergi.Watermark(text, img1, image.Pt(int(v), img1.Bounds().Max.Y/2))
		frames = append(frames, img)
	}

	gif, _ := mergi.Animate(frames, 1)
	mergi.Export(io.NewAnimationExporter(gif, "out.gif"))
}
