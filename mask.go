package mergi

import (
	"image/color"
	"image"
	"image/draw"
)

// MaskBlack provides black masking
var MaskBlack = color.RGBA{0, 0, 0, 0}

// MaskWhite provides white masking
var MaskWhite = color.RGBA{255, 255, 255, 0}
// Mask uses go standard image.Image to get the masked image with original image,
//
// Mask simply match the given color and matched with mask image and original to replace alpha values 0, 255
//
// for more mask examples https://github.com/noelyahan/mergi/examples/mask
func Mask(maskImg, original image.Image, maskColor color.RGBA) (image.Image, error) {
	bounds := original.Bounds()
	maskB := image.NewAlpha(bounds)
	R, G, B, _ := maskColor.RGBA()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			c := maskImg.At(x, y)
			r, g, b, _ := c.RGBA()
			if r == R && g == G && b == B {
				maskB.SetAlpha(x, y, color.Alpha{uint8(0)})
			}else {
				maskB.SetAlpha(x, y, color.Alpha{uint8(255)})
			}
		}
	}
	res := image.NewRGBA(bounds)
	draw.DrawMask(res, bounds, original, image.ZP, maskB, image.ZP, draw.Over)
	return res, nil
}
