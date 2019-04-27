package mergi

import (
	"image"
	"image/color"
)

func Transit(imgs1, imgs2, trans []image.Image, mask color.RGBA, start, end, speed float64) []image.Image {
	// validation imgs1, imgs2, trans must be same length || imgs1, imgs2 must be length 1
	frames := make([]image.Image, 0)
	var msk image.Image
	isSingleImage := len(imgs1) == 1 && len(imgs2) == 1
	for i := start; i <= end; i += speed {
		c := int(i)
		msk, _ = Resize(trans[c], uint(imgs1[0].Bounds().Max.X), uint(imgs1[0].Bounds().Max.Y))
		if isSingleImage {
			c = 0
		}
		msk, _ = Mask(msk, imgs2[c], mask)
		msk, _ = Watermark(msk, imgs1[c], image.ZP)
		frames = append(frames, msk)
	}
	return frames
}
