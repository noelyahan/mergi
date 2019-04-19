package mergi_test

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
	"log"
	"testing"
	"strings"
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