package transition

import (
	"image"
	"github.com/noelyahan/mergi"
	"image/color"
)

type Transition []image.Image

func Image(img1, img2 image.Image, trans Transition, mask color.RGBA, start, end, speed float64) []image.Image {
	frames := make([]image.Image, 0)
	var msk image.Image
	for i := start; i <= end; i += speed {
		c := int(i)
		msk, _ = mergi.Resize(trans[c], uint(img1.Bounds().Max.X), uint(img1.Bounds().Max.Y))
		msk, _ = mergi.Mask(msk, img2, mask)
		res, _ := mergi.Watermark(msk, img1, image.ZP)
		frames = append(frames, res)
	}
	return frames
}

func Images(imgs1, imgs2, trans Transition, mask color.RGBA, start, end, speed float64) []image.Image {
	frames := make([]image.Image, 0)
	var msk image.Image
	for i := start; i <= end; i += speed {
		c := int(i)
		msk, _ = mergi.Resize(trans[c], uint(imgs1[c].Bounds().Max.X), uint(imgs1[c].Bounds().Max.Y))
		msk, _ = mergi.Mask(msk, imgs2[c], mask)
		msk, _ = mergi.Watermark(msk, imgs1[c], image.ZP)
		frames = append(frames, msk)
	}
	return frames
}
