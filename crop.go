package mergi

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
)

var errCrop = errors.New("Mergi found a error image on Crop")
var errCropBound = errors.New("Mergi expects more than 0 value for bounds")

// Crop uses go standard image.Image, the starting X, Y position as go standard image.Point crop width and height as image.Point
// returns the crop image output
//
// for more crop examples https://github.com/noelyahan/mergi/examples/crop
func Crop(img image.Image, p1 image.Point, p2 image.Point) (image.Image, error) {
	if img == nil {
		return nil, errCrop
	}
	if p1.X < 0 || p1.Y < 0 || p2.X < 0 || p2.Y < 0 {
		fmt.Println()
		return nil, errCropBound
	}
	b := image.Rect(0, 0, p2.X, p2.Y)
	resImg := image.NewRGBA(b)
	draw.Draw(resImg, b, img, p1, draw.Src)
	return resImg, nil
}
