package mergi_test

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/impexp"
	"image"
	"testing"
)

func TestAnimationWithNil(t *testing.T) {
	res, err := mergi.Animate([]image.Image{nil, nil}, 10)
	if len(res.Delay) != 0 {
		t.Error("Expects res to be empty gif", res)
	}
	if err == nil {
		t.Error("Expects error cannot be nil", err)
	}
}

func TestAnimationSlide(t *testing.T) {
	img, _ := mergi.Import(impexp.NewURLImporter("https://cdn.pixabay.com/photo/2014/06/11/17/00/cook-366875__340.jpg"))
	cropSize := image.Pt(110, 130)
	slide1, _ := mergi.Crop(img, image.Pt(450, 200), cropSize)
	slide2, _ := mergi.Crop(img, image.Pt(340, 200), cropSize)
	slide3, _ := mergi.Crop(img, image.Pt(230, 200), cropSize)
	slide4, _ := mergi.Crop(img, image.Pt(120, 200), cropSize)
	slide5, _ := mergi.Crop(img, image.Pt(10, 200), cropSize)

	gif, err := mergi.Animate([]image.Image{slide1, slide2, slide3, slide4, slide5}, 50)
	if err != nil {
		t.Error("Expets no error!", err)
	}

	mergi.Export(impexp.NewAnimationExporter(gif, "out.gif"))

}
