package main

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/eximp"
	"image"
	"github.com/noelyahan/mergi/ease"
)

func main() {
	mergiLogo, _ := mergi.Import(eximp.NewFileImporter("testdata/mergi_logo_watermark.png"))
	coffee, _ := mergi.Import(eximp.NewFileImporter("testdata/coffee-171653_960_720.jpg"))

	// scale down
	mergiLogoSmall, _ := mergi.Resize(mergiLogo, uint(mergiLogo.Bounds().Max.X/2), uint(mergiLogo.Bounds().Max.Y/2))

	cropAnimation(mergiLogoSmall, ease.InBounce)
	watermarkAnimation(coffee, mergiLogoSmall, ease.OutBounce)
	cropMergeAnimation(mergiLogoSmall, mergiLogoSmall, ease.InBounce, ease.InElastic)
}

func cropAnimation(cherry image.Image, move ease.EaseType) {
	images := make([]image.Image, 0)

	arr := ease.AnimatePoints(move, image.Pt(0, 0), image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y), 3.5)
	// animate with crop operation
	for _, v := range arr {
		img, _ := mergi.Crop(cherry, v, image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y))
		images = append(images, img)
	}

	gif, _ := mergi.Animate(images, 2)
	mergi.Export(eximp.NewAnimationExporter(gif, "examples/easing/res/crop_ease.gif"))
}

func watermarkAnimation(coffee, cherry image.Image, move ease.EaseType) {
	images := make([]image.Image, 0)
	arr := ease.AnimatePoints(move, image.Pt(0, 0), image.Pt(coffee.Bounds().Max.X-cherry.Bounds().Max.X, coffee.Bounds().Max.Y-cherry.Bounds().Max.Y), 3.5)
	// animate with watermark operation
	for _, v := range arr {
		img, _ := mergi.Watermark(cherry, coffee, v)
		images = append(images, img)
	}

	gif, _ := mergi.Animate(images, 2)
	mergi.Export(eximp.NewAnimationExporter(gif, "examples/easing/res/watermark_ease.gif"))
}

func cropMergeAnimation(coffee, cherry image.Image, move1, move2 ease.EaseType) {
	images := make([]image.Image, 0)

	// animate with resize, crop and merge operation
	cherryAnimArr := ease.AnimatePoints(move1, image.Pt(0, 0), image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y), 3.5)
	coffeeAnimArr := ease.AnimatePoints(move2, image.Pt(0, 0), image.Pt(cherry.Bounds().Max.X, cherry.Bounds().Max.Y), 3.5)

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
	mergi.Export(eximp.NewAnimationExporter(gif, "examples/easing/res/crop_merge_ease.gif"))
}
