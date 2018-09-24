package mergi

import (
	"errors"
	"image"
	"image/draw"
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
	b := original.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, original, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(p), watermark, image.ZP, draw.Over)
	return m, nil
}
