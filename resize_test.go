package mergi_test

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"image"
	"log"
	"testing"
)

func TestResizeWithNil(t *testing.T) {
	mergi.Resize(nil, 100, 100)
}

func TestResizeScaleDown(t *testing.T) {
	img, _ := mergi.Import(loader.NewFileImporter("./testdata/hedgehog-child-3636026_960_720.jpg"))
	w := uint(img.Bounds().Max.X / 2)
	h := uint(img.Bounds().Max.Y / 2)
	res, err := mergi.Resize(img, w, h)
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}
	b := res.Bounds().Max
	if b.X != int(w) || b.Y != int(h) {
		t.Errorf("Expected [width %d, height %d] got [width %d, height %d]",
			w, h, b.X, b.Y)
	}
	mergi.Export(loader.NewFileExporter(res, "out.png"))
}

func TestResizeScaleUp(t *testing.T) {
	img, _ := mergi.Import(loader.NewFileImporter("./testdata/hedgehog-child-3636026_960_720.jpg"))
	w := uint(img.Bounds().Max.X * 2)
	h := uint(img.Bounds().Max.Y * 2)
	res, err := mergi.Resize(img, w, h)
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}
	b := res.Bounds().Max
	if b.X != int(w) || b.Y != int(h) {
		t.Errorf("Expected [width %d, height %d] got [width %d, height %d]",
			w, h, b.X, b.Y)
	}
	mergi.Export(loader.NewFileExporter(res, "out.png"))
}

func TestMergeScaleDown(t *testing.T) {
	template := "TT"
	img1, _ := mergi.Import(loader.NewFileImporter("./testdata/cherry-3074284_960_720.jpg"))
	img2, _ := mergi.Import(loader.NewFileImporter("./testdata/avocado-3210885_960_720.jpg"))
	imgs := []image.Image{img1, img2}

	img, err := mergi.Merge(template, imgs)
	if err != nil {
		log.Fatalf("failed to merge: %s", err)
	}
	if img == nil {
		t.Errorf("Expected a image, got [%v]", img)
		return
	}
	w := uint(img.Bounds().Max.X / 2)
	h := uint(img.Bounds().Max.Y / 2)
	res, err := mergi.Resize(img, w, h)
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}
	b := res.Bounds().Max
	if b.X != int(w) || b.Y != int(h) {
		t.Errorf("Expected [width %d, height %d] got [width %d, height %d]",
			w, h, b.X, b.Y)
	}
	mergi.Export(loader.NewFileExporter(res, "out.png"))
}

func TestMergeScaleUp(t *testing.T) {
	template := "TB"
	img1, _ := mergi.Import(loader.NewFileImporter("./testdata/cherry-3074284_960_720.jpg"))
	img2, _ := mergi.Import(loader.NewFileImporter("./testdata/avocado-3210885_960_720.jpg"))
	imgs := []image.Image{img1, img2}

	img, err := mergi.Merge(template, imgs)
	if err != nil {
		log.Fatalf("failed to merge: %s", err)
	}
	if img == nil {
		t.Errorf("Expected a image, got [%v]", img)
		return
	}
	w := uint(img.Bounds().Max.X * 2)
	h := uint(img.Bounds().Max.Y * 2)
	res, err := mergi.Resize(img, w, h)
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}
	b := res.Bounds().Max
	if b.X != int(w) || b.Y != int(h) {
		t.Errorf("Expected [width %d, height %d] got [width %d, height %d]",
			w, h, b.X, b.Y)
	}
	mergi.Export(loader.NewFileExporter(res, "out.png"))
}
