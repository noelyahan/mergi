package mergi

import (
	"testing"
	"github.com/noelyahan/eximp"
	"image"
	"github.com/noelyahan/mergitrans"
)

func _getImages() (img1 image.Image, img2 image.Image) {
	scale := 4
	img1, _ = Import(eximp.NewFileImporter("./testdata/nature-3042751_960_720.jpg"))
	img1, _ = Resize(img1, uint(img1.Bounds().Max.X/scale), uint(img1.Bounds().Max.Y/scale))

	img2, _ = Import(eximp.NewFileImporter("./testdata/soldier-708711_960_720.jpg"))
	img2, _ = Resize(img2, uint(img2.Bounds().Max.X/scale), uint(img2.Bounds().Max.Y/scale))

	img1, _ = Crop(img1, image.ZP, image.Pt(img2.Bounds().Max.X, img2.Bounds().Max.Y))
	return
}

func TestImage(t *testing.T) {
	img1, img2 := _getImages()

	frames := Transit([]image.Image{img1}, []image.Image{img2}, mergitrans.Ink1, MaskBlack, 0, float64(len(mergitrans.Ink1)-1) , 1)

	gif, _ := Animate(frames, 1)
	Export(eximp.NewAnimationExporter(gif, "out.gif"))
}


func TestImages(t *testing.T) {
	trans := mergitrans.Ink2
	frames := Transit(mergitrans.Videos.PoppyField, mergitrans.Videos.Clouds, trans, MaskBlack, 0, float64(len(trans)-1), 1)

	gif, _ := Animate(frames, 1)
	Export(eximp.NewAnimationExporter(gif, "out.gif"))
}