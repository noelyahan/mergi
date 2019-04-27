package mergi

import (
	"testing"
	"github.com/noelyahan/mergi/io"
	"image"
)

func TestMaskAdvanced(t *testing.T) {
	img, _ := Import(io.NewFileImporter("./testdata/black_circle.jpg"))
	img, _ = Resize(img, uint(img.Bounds().Max.X/3), uint(img.Bounds().Max.Y/3))
	img, _ = Merge("TBBTBBTBBTBBTBB", []image.Image{
		img, img, img,
		img, img, img,
		img, img, img,
		img, img, img,
		img, img, img,
	})
	img1, img2 := _getImages()
	msk, _ := Mask(img, img1, MaskBlack)
	res, _ := Watermark(msk, img2, image.ZP)
	Export(io.NewFileExporter(res, "out.png"))
}
