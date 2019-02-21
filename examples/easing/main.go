package main

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"image"
	"github.com/noelyahan/mergi/ease"
)

func main() {
	cherry, _ := mergi.Import(loader.NewFileImporter("testdata/cherry-3074284_960_720.jpg"))
	coffee, _ := mergi.Import(loader.NewFileImporter("testdata/coffee-171653_960_720.jpg"))

	// scale down
	cherrySmall, _ := mergi.Resize(cherry, uint(cherry.Bounds().Max.X/4), uint(cherry.Bounds().Max.Y/4))
	coffeeSmall, _ := mergi.Resize(coffee, uint(coffee.Bounds().Max.X/4), uint(coffee.Bounds().Max.Y/4))

	cropAnimation(cherrySmall)
	watermarkAnimation(coffee, cherrySmall)
	cropMergeAnimation(coffeeSmall, cherrySmall)
}

func cropAnimation(cherry image.Image) {
	images := make([]image.Image, 0)

	arr := ease.AnimatePoints(ease.OutBounce, image.Pt(0, 0), image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y), 3.5)
	// animate with crop operation
	for _, v := range arr {
		img, _ := mergi.Crop(cherry, v, image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y))
		images = append(images, img)
	}

	gif, _ := mergi.Animate(images, 2)
	mergi.Export(loader.NewAnimationExporter(gif, "examples/easing/res/crop_ease.gif"))
}

func watermarkAnimation(coffee, cherry image.Image) {
	images := make([]image.Image, 0)
	arr := ease.AnimatePoints(ease.OutBounce, image.Pt(0, 0), image.Pt(coffee.Bounds().Max.X-cherry.Bounds().Max.X, coffee.Bounds().Max.Y-cherry.Bounds().Max.Y), 3.5)
	// animate with watermark operation
	for _, v := range arr {
		img, _ := mergi.Watermark(cherry, coffee, v)
		images = append(images, img)
	}

	gif, _ := mergi.Animate(images, 2)
	mergi.Export(loader.NewAnimationExporter(gif, "examples/easing/res/watermark_ease.gif"))
}

func cropMergeAnimation(coffee, cherry image.Image) {
	images := make([]image.Image, 0)

	// animate with resize, crop and merge operation
	cherryAnimArr := ease.AnimatePoints(ease.InBounce, image.Pt(0, 0), image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y), 3.5)
	coffeeAnimArr := ease.AnimatePoints(ease.InElastic, image.Pt(0, 0), image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y), 3.5)

	cherryArr := make([]image.Image, 0)
	coffeeArr := make([]image.Image, 0)
	for _, v := range cherryAnimArr {
		img, _ := mergi.Crop(cherry, v, image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y))
		cherryArr = append(cherryArr, img)
	}

	for _, v := range coffeeAnimArr {
		img, _ := mergi.Crop(coffee, v, image.Pt(coffee.Bounds().Max.X, coffee.Bounds().Max.Y))
		coffeeArr = append(coffeeArr, img)
	}

	for i := 0; i < len(cherryAnimArr); i++ {
		img, _ := mergi.Merge("TT", []image.Image{cherryArr[i], coffeeArr[i]})
		images = append(images, img)
	}

	gif, _ := mergi.Animate(images, 2)
	mergi.Export(loader.NewAnimationExporter(gif, "examples/easing/res/crop_merge_ease.gif"))
}
