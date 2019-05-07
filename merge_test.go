package mergi_test

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/impexp"
	"image"
	"strings"
	"testing"
)

func TestMergeTemplateWithNil(t *testing.T) {
	template := "TB"
	res, err := mergi.Merge(template, []image.Image{nil, nil})
	if err == nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	if err != nil {
		t.Log(err)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplateMismatch(t *testing.T) {
	template := "TB"
	res, err := mergi.Merge(template, []image.Image{nil})
	if err == nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	if err != nil {
		t.Log(err)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate1(t *testing.T) {
	template := "TBB"
	imgs := getImages(template)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	b := res.Bounds().Max
	w := imgs[0].Bounds().Max.X
	h := imgs[0].Bounds().Max.Y + imgs[1].Bounds().Max.Y + imgs[2].Bounds().Max.Y
	if b.X != w {
		t.Fatalf("Expect width [%d] got [%d]", w, b.X)
	}
	if b.Y != h {
		t.Fatalf("Expect height [%d] got [%d]", h, b.Y)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate2(t *testing.T) {
	template := "TBTB"
	imgs := getImages(template)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	b := res.Bounds().Max
	w := imgs[0].Bounds().Max.X + imgs[1].Bounds().Max.X
	h := imgs[0].Bounds().Max.Y + imgs[1].Bounds().Max.Y
	if b.X != w {
		t.Fatalf("Expect width [%d] got [%d]", w, b.X)
	}
	if b.Y != h {
		t.Fatalf("Expect height [%d] got [%d]", h, b.Y)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate3(t *testing.T) {
	template := "TTBBB"
	large, _ := mergi.Import(impexp.NewFileImporter("./testdata/nature-3042751_960_720.jpg"))
	small, _ := mergi.Import(impexp.NewFileImporter("./testdata/evraiki-2514543_240_180.jpg"))
	imgs := make([]image.Image, 0)
	imgs = append(imgs, large)
	imgs = append(imgs, small)
	imgs = append(imgs, small)
	imgs = append(imgs, small)
	imgs = append(imgs, small)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	b := res.Bounds().Max
	w := imgs[0].Bounds().Max.X + imgs[1].Bounds().Max.X
	h := imgs[0].Bounds().Max.Y
	if b.X != w {
		t.Fatalf("Expect width [%d] got [%d]", w, b.X)
	}
	if b.Y != h {
		t.Fatalf("Expect height [%d] got [%d]", h, b.Y)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate4(t *testing.T) {
	template := "TTBBBT"
	large, _ := mergi.Import(impexp.NewFileImporter("./testdata/nature-3042751_960_720.jpg"))
	small, _ := mergi.Import(impexp.NewFileImporter("./testdata/evraiki-2514543_240_180.jpg"))
	imgs := make([]image.Image, 0)
	imgs = append(imgs, large)
	imgs = append(imgs, small)
	imgs = append(imgs, small)
	imgs = append(imgs, small)
	imgs = append(imgs, small)
	imgs = append(imgs, large)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	b := res.Bounds().Max
	w := imgs[0].Bounds().Max.X + imgs[1].Bounds().Max.X + imgs[5].Bounds().Max.X
	h := imgs[0].Bounds().Max.Y
	if b.X != w {
		t.Fatalf("Expect width [%d] got [%d]", w, b.X)
	}
	if b.Y != h {
		t.Fatalf("Expect height [%d] got [%d]", h, b.Y)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate5(t *testing.T) {
	template := "TBBTBB"
	imgs := getImages(template)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	b := res.Bounds().Max
	w := imgs[0].Bounds().Max.X + imgs[3].Bounds().Max.X
	h := imgs[0].Bounds().Max.Y + imgs[1].Bounds().Max.Y + imgs[2].Bounds().Max.Y
	if b.X != w {
		t.Fatalf("Expect width [%d] got [%d]", w, b.X)
	}
	if b.Y != h {
		t.Fatalf("Expect height [%d] got [%d]", h, b.Y)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate6(t *testing.T) {
	template := "TBBTBBTBB"
	imgs := getImages(template)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	b := res.Bounds().Max
	w := imgs[0].Bounds().Max.X + imgs[3].Bounds().Max.X + imgs[6].Bounds().Max.X
	h := imgs[0].Bounds().Max.Y + imgs[1].Bounds().Max.Y + imgs[2].Bounds().Max.Y
	if b.X != w {
		t.Fatalf("Expect width [%d] got [%d]", w, b.X)
	}
	if b.Y != h {
		t.Fatalf("Expect height [%d] got [%d]", h, b.Y)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate7(t *testing.T) {
	template := "TBBBTBBBTBBBTBBB"
	imgs := getImages(template)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	b := res.Bounds().Max
	w := imgs[0].Bounds().Max.X + imgs[4].Bounds().Max.X + imgs[8].Bounds().Max.X + imgs[12].Bounds().Max.X
	h := imgs[0].Bounds().Max.Y + imgs[1].Bounds().Max.Y + imgs[2].Bounds().Max.Y + imgs[3].Bounds().Max.Y
	if b.X != w {
		t.Fatalf("Expect width [%d] got [%d]", w, b.X)
	}
	if b.Y != h {
		t.Fatalf("Expect height [%d] got [%d]", h, b.Y)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate8(t *testing.T) {
	template := "TBBBBBBTBBBBBBTBBBBBBTBBBBBB"
	imgs := getImages(template)
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate9(t *testing.T) {
	img, _ := mergi.Import(impexp.NewFileImporter("./testdata/evraiki-2514543_240_180.jpg"))
	imgs := []image.Image{img, img, img, img, img, img, img, img, img}
	res, err := mergi.Merge("TBBTBBTBB", imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

func TestMergeTemplate10(t *testing.T) {
	template := "TTTTT"
	imgs := make([]image.Image, 0)
	img, _ := mergi.Import(impexp.NewFileImporter("./testdata/evraiki-2514543_240_180.jpg"))
	for i := 0; i < len(strings.Split(template, "")); i++ {
		imgs = append(imgs, img)
	}
	res, err := mergi.Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	mergi.Export(impexp.NewFileExporter(res, "out.png"))
}

// helper functions
func getImages(template string) []image.Image {
	imgs := make([]image.Image, 0)
	img, _ := mergi.Import(impexp.NewFileImporter("./testdata/evraiki-2514543_240_180.jpg"))
	for i := 0; i < len(strings.Split(template, "")); i++ {
		imgs = append(imgs, img)
	}
	return imgs
}
