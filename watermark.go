package mergi

import (
	"errors"
	"image"
	"image/draw"
	"image/color"
)

// Watermark uses go standard image.Image to get the watermark image and original image that want to watermark,
//
// the position of the watermark has to provide in image.Point then it'll returns the watermarked image output
//
// for more watermark examples https://github.com/noelyahan/mergi/examples/watermark
func Watermark(watermark, original image.Image, p image.Point) (image.Image, error) {
	if watermark == nil {
		msg := "Mergi found a error watermark image"
		return nil, errors.New(msg)
	}
	if original == nil {
		msg := "Mergi found a error original image"
		return nil, errors.New(msg)
	}
	oBounds := original.Bounds()
	oRGBA := image.NewRGBA(oBounds)

	draw.Draw(oRGBA, oBounds, original, image.ZP, draw.Src)
	draw.Draw(oRGBA, watermark.Bounds().Add(p), watermark, image.ZP, draw.Over)
	return oRGBA, nil
}

// Opacity uses go standard image.Image to change the alpha channel of the given image,
//
// the floating point alpha amount has to provide with the given image and it'll return the opacity image
//
// for more opacity examples https://github.com/noelyahan/mergi/examples/opacity
func Opacity(img image.Image, alpha float64) (image.Image, error) {
	mapValues := func(value, start1, stop1, start2, stop2 float64) int {
		return int(start2 + (stop2-start2)*((value-start1)/(stop1-start1)))
	}

	if alpha < 0 {
		alpha = 0
	}else if alpha > 1 {
		alpha = 1
	}

	bounds := img.Bounds()
	mask := image.NewAlpha(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			r := mapValues(alpha, 1, 0, 0, 255)
			mask.SetAlpha(x, y, color.Alpha{uint8(255 - r)})
		}
	}

	mskWatermark := image.NewRGBA(bounds)
	draw.DrawMask(mskWatermark, bounds, img, image.ZP, mask, image.ZP, draw.Over)
	return mskWatermark, nil
}

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