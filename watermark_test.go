package mergi_test

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
	"log"
	"testing"
	"strings"
	"image/color"
)

func TestWatermarkLogo(t *testing.T) {
	watermark, _ := mergi.Import(io.NewFileImporter("./testdata/mergi_logo_watermark.png"))
	max := watermark.Bounds().Max
	watermark, err := mergi.Resize(watermark, uint(max.X/2), uint(max.Y/2))
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}
	img, _ := mergi.Import(io.NewFileImporter("./testdata/coffee-206142_960_720.jpg"))
	max = img.Bounds().Max

	res, err := mergi.Watermark(watermark, img, image.Pt(max.X-watermark.Bounds().Max.X, max.Y-watermark.Bounds().Max.Y))
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}

	res, err = mergi.Resize(res, uint(res.Bounds().Max.X*2), uint(res.Bounds().Max.Y*2))
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}

	mergi.Export(io.NewFileExporter(res, "out.png"))
}

func TestWatermarkWithNil(t *testing.T) {
	img, _ := mergi.Import(io.NewFileImporter("./testdata/coffee-1291656_960_720.jpg"))
	res, err := mergi.Watermark(nil, img, image.Pt(0, 0))
	if err == nil {
		t.Errorf("Expect error got [%v]", err)
	}
	mergi.Export(io.NewFileExporter(res, "out.png"))
}


// https://stackoverflow.com/questions/12484403/setting-opacity-of-image-in-golang
func TestOpacity(t *testing.T) {
	bg, _ := mergi.Import(io.NewFileImporter("./testdata/lion-3576045_960_720.jpg"))
	wm, _ := mergi.Import(io.NewFileImporter("./testdata/mergi_logo_watermark.png"))
	max := wm.Bounds().Max
	wm, err := mergi.Resize(wm, uint(max.X/2), uint(max.Y/2))
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}

	wmAlpha, _ := mergi.Opacity(wm, 0.3)

	wmarks := make([]image.Image, 0)
	tmplate := "TBBBBBBBTBBBBBBBTBBBBBBBTBBBBBBBTBBBBBBBTBBBBBBB"
	for _ = range strings.Split(tmplate, "") {
		wmarks = append(wmarks, wmAlpha)
	}
	wmAlpha, _ = mergi.Merge(tmplate, wmarks)
	res, err := mergi.Watermark(wmAlpha, bg, image.ZP)

	mergi.Export(io.NewFileExporter(res, "out.png"))

}

func get2Images() (img1 image.Image, img2 image.Image) {
	scale := 4
	img1, _ = mergi.Import(io.NewFileImporter("./testdata/nature-3042751_960_720.jpg"))
	img1, _ = mergi.Resize(img1, uint(img1.Bounds().Max.X/scale), uint(img1.Bounds().Max.Y/scale))

	img2, _ = mergi.Import(io.NewFileImporter("./testdata/soldier-708711_960_720.jpg"))
	img2, _ = mergi.Resize(img2, uint(img2.Bounds().Max.X/scale), uint(img2.Bounds().Max.Y/scale))

	img1, _ = mergi.Crop(img1, image.ZP, image.Pt(img2.Bounds().Max.X, img2.Bounds().Max.Y))
	return
}

func TestMaskAdvanced(t *testing.T) {
	img, _ := mergi.Import(io.NewFileImporter("./testdata/black_circle.jpg"))
	img, _ = mergi.Resize(img, uint(img.Bounds().Max.X/3), uint(img.Bounds().Max.Y/3))
	img, _ = mergi.Merge("TBBTBBTBBTBBTBB", []image.Image{
		img, img, img,
		img, img, img,
		img, img, img,
		img, img, img,
		img, img, img,
		})
	img1, img2 := get2Images()
	msk, _ := mergi.Mask(img, img1, color.RGBA{0, 0, 0, 0})
	res, _ := mergi.Watermark(msk, img2, image.ZP)
	mergi.Export(io.NewFileExporter(res, "out.png"))
}