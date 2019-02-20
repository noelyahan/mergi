package main

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"image"
	"github.com/noelyahan/mergi/ease"
)

func mapVlaues(value, start1, stop1, start2, stop2 float64) float64 {
	return start2 + (stop2-start2)*((value-start1)/(stop1-start1))
}


func main() {
	cherry, _ := mergi.Import(loader.NewFileImporter("testdata/cherry-3074284_960_720.jpg"))
	// scale down
	cherry, _ = mergi.Resize(cherry, uint(cherry.Bounds().Max.X/4), uint(cherry.Bounds().Max.Y/4))

	images := make([]image.Image, 0)

	for x := float64(cherry.Bounds().Max.X); x > 0.0; x -= 2.5 {
		y := ease.AnimateXY(ease.InBounce, x, 0, float64(cherry.Bounds().Max.X), 0, float64(cherry.Bounds().Max.Y))
		img, _ := mergi.Crop(cherry, image.Pt(int(x), int(y)), image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y))
		images = append(images, img)
	}

	gif, _ := mergi.Animate(images, 2)
	mergi.Export(loader.NewAnimationExporter(gif, "examples/easing/res/ease.gif"))
}