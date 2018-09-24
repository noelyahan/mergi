package mergi

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
)

// Resize uses go standard image.Image, unsigned int for width and height that want to resize
// returns the resize image output
//
// for more resize examples https://github.com/noelyahan/mergi/examples/resize
func Resize(img image.Image, w, h uint) (image.Image, error) {
	if img == nil {
		return nil, errors.New("Mergi found a error image on Resize")
	}
	return resize.Resize(w, h, img, resize.Lanczos3), nil
}
