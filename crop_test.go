package mergi

import (
	"github.com/noelyahan/eximp"
	"image"
	"strings"
	"testing"
)

func TestCropWithNil(t *testing.T) {
	Crop(nil, image.Pt(0, 0), image.Pt(0, 0))
}

func TestCropWithNegativeBounds(t *testing.T) {
	imgs := getImages("T")
	_, err := Crop(imgs[0], image.Pt(-10, -10), image.Pt(-1, -1))
	if err == nil {
		t.Errorf("Expect error got [%v]", err)
		return
	}
}

func TestCrop(t *testing.T) {
	imgs := getImages("T")
	p1 := image.Pt(40, 0)
	p2 := image.Pt(120, 160)
	img, err := Crop(imgs[0], p1, p2)
	if err != nil {
		t.Errorf("Crop test fails [%v]", err)
		return
	}
	Export(eximp.NewFileExporter(img, "out.png"))
}

func TestCropWithHeigthWidth(t *testing.T) {
	imgs := getImages("T")
	p1 := image.Pt(40, 10)
	h := 120
	w := 160
	p2 := image.Pt(p1.X+h, p1.Y+w)
	img, err := Crop(imgs[0], p1, p2)
	if err != nil {
		t.Errorf("Crop test fails [%v]", err)
		return
	}
	Export(eximp.NewFileExporter(img, "out.png"))
}

func TestGetYAxisCount(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"TBB", 2},
		{"TBTB", 1},
		{"TTTB", 0},
		{"TTBB", 0},
		{"TTBBT", 0},
		{"TBBBBT", 4},
		{"TBBTBBTBB", 2},
	}

	for _, test := range tests {
		c := getYAxisCount(test.in)
		if test.out != c {
			t.Errorf("Want %d got %d", test.out, c)
		}
	}

}

func TestCropAndMerge(t *testing.T) {
	img, _ := Import(eximp.NewFileImporter("./testdata/tree-146874_960_720.png"))
	imgs := make([]image.Image, 5)
	template := "TTTTT"
	imgs[0], _ = Crop(img, image.Pt(0, 0), image.Pt(100, 140))
	imgs[1], _ = Crop(img, image.Pt(0, 145), image.Pt(100, 140))
	imgs[2], _ = Crop(img, image.Pt(0, 290), image.Pt(100, 140))
	imgs[3], _ = Crop(img, image.Pt(0, 435), image.Pt(100, 140))
	imgs[4], _ = Crop(img, image.Pt(0, 580), image.Pt(100, 140))
	res, err := Merge(template, imgs)
	if err != nil {
		t.Fatalf("Expect error got [%v]", err)
	}
	Export(eximp.NewFileExporter(res, "out.png"))
}

// helper functions
func getImages(template string) []image.Image {
	imgs := make([]image.Image, 0)
	img, _ := Import(eximp.NewFileImporter("./testdata/evraiki-2514543_240_180.jpg"))
	for i := 0; i < len(strings.Split(template, "")); i++ {
		imgs = append(imgs, img)
	}
	return imgs
}