package main

import (
	"github.com/noelyahan/mergi"
	"image"
	"image/gif"
	"log"
	"github.com/noelyahan/mergi/io"
)

// This example will guide how to animate any given images using mergi.Animate API
// mergi.Animate only needs frames of go lang images and some delay to sleep
func main() {
	gifAnim := smoothMoveWithCrop()
	mergi.Export(io.NewAnimationExporter(gifAnim, "examples/animate/res/smooth.gif"))

	gifAnim = simpleSlideChange()
	mergi.Export(io.NewAnimationExporter(gifAnim, "examples/animate/res/slide.gif"))

	gifAnim = catFighterSpriteSheet()
	mergi.Export(io.NewAnimationExporter(gifAnim, "examples/animate/res/sprite.gif"))
}

func catFighterSpriteSheet() gif.GIF {
	img, err := mergi.Import(io.NewFileImporter("testdata/cat_fighter_sprite1.png"))
	if err != nil {
		log.Fatal(err)
	}
	imgs := make([]image.Image, 0)
	x := 0
	w := 50
	h := 50
	for i := 0; i < 10; i++ {
		res, _ := mergi.Crop(img, image.Pt(x, 0), image.Pt(w, h))
		x += w
		res, err = mergi.Resize(res, uint(w*3), uint(h*3))
		if err != nil {
			log.Fatal(err)
		}
		imgs = append(imgs, res)
	}

	anim, err := mergi.Animate(imgs, 10)
	if err != nil {
		log.Fatal(err)
	}
	return anim
}

func simpleSlideChange() gif.GIF {
	cherry, _ := mergi.Import(io.NewFileImporter("testdata/cherry-3074284_960_720.jpg"))
	grapes, _ := mergi.Import(io.NewFileImporter("testdata/grapes-2032838_960_720.jpg"))
	smoothie, _ := mergi.Import(io.NewFileImporter("testdata/smoothie-3193660_960_720.jpg"))
	w := uint(240)
	h := uint(180)
	cherryR, _ := mergi.Resize(cherry, w, h)
	grapesR, _ := mergi.Resize(grapes, w, h)
	smoothieR, _ := mergi.Resize(smoothie, w, h)
	imgs := []image.Image{cherryR, grapesR, smoothieR}

	anim, err := mergi.Animate(imgs, 50)
	if err != nil {
		log.Fatal(err)
	}
	return anim
}

func smoothMoveWithCrop() gif.GIF {
	cherry, _ := mergi.Import(io.NewFileImporter("testdata/cherry-3074284_960_720.jpg"))
	grapes, _ := mergi.Import(io.NewFileImporter("testdata/grapes-2032838_960_720.jpg"))
	smoothie, _ := mergi.Import(io.NewFileImporter("testdata/smoothie-3193660_960_720.jpg"))
	w := uint(240)
	h := uint(180)
	cherryR, _ := mergi.Resize(cherry, w, h)
	grapesR, _ := mergi.Resize(grapes, w, h)
	smoothieR, _ := mergi.Resize(smoothie, w, h)
	img, err := mergi.Merge("TTTT", []image.Image{cherryR, grapesR, smoothieR, cherryR})
	if err != nil {
		log.Fatal(err)
	}

	imgs := make([]image.Image, 0)
	b := img.Bounds().Max
	cropSize := image.Pt(int(w), int(h))
	for i := 0; i < b.X-int(w); i += 6 {
		tmp, _ := mergi.Crop(img, image.Pt(i, 0), cropSize)
		imgs = append(imgs, tmp)
	}

	anim, err := mergi.Animate(imgs, 5)
	if err != nil {
		log.Fatal(err)
	}
	return anim
}
