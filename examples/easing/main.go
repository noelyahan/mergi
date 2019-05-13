package main

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/impexp"
	"image"
)

func main() {
	// Load background and the square images
	square, _ := mergi.Import(impexp.NewFileImporter("testdata/square.jpg"))
	bg, _ := mergi.Import(impexp.NewFileImporter("testdata/white_bg.jpg"))

	// Init images frames to add applied ease frames
	frames := make([]image.Image, 0)

	// Init the limts of the Ease
	to := bg.Bounds().Max.X - square.Bounds().Max.X
	posY := bg.Bounds().Max.Y/2 - square.Bounds().Max.Y/2
	speed := 4

	// Ease from 0 to width of background
	for i := 0; i < to; i += speed {
		// Apply Easeing function InBounce
		posX := mergi.Ease(float64(i), 0, float64(to), mergi.InBounce)
		img, _ := mergi.Watermark(square, bg, image.Pt(int(posX), posY))
		frames = append(frames, img)
	}

	// For preview example, save as a gif
	gif, _ := mergi.Animate(frames, 1)
	mergi.Export(impexp.NewAnimationExporter(gif, "examples/easing/res/out.gif"))
}