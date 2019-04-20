package transition

import (
	"image"
	"github.com/noelyahan/mergi"
)

func mapValues(value, start1, stop1, start2, stop2 float64) float64 {
	return start2 + (stop2-start2)*((value-start1)/(stop1-start1))
}

func FadeIn(img1, img2 image.Image, time, speed float64) []image.Image {
	imgs := make([]image.Image, 0)
	for i := time; i > 0; i -= speed {
		v := mapValues(i, time, 0, 1, 0)
		res, _ := mergi.Opacity(img1, v)
		imgs = append(imgs, res)
	}

	for i := 0.0; i < time; i += speed {
		v := mapValues(i, 0, time, 0, 1)
		res, _ := mergi.Opacity(img2, v)
		imgs = append(imgs, res)
	}
	return imgs
}

func FadeOut(img1, img2 image.Image, time, speed float64) []image.Image {
	imgs := make([]image.Image, 0)

	for i := 0.0; i < time; i += speed {
		v := mapValues(i, 0, time, 0, 1)
		res, _ := mergi.Opacity(img2, v)
		imgs = append(imgs, res)
	}

	for i := time; i > 0; i -= speed {
		v := mapValues(i, time, 0, 1, 0)
		res, _ := mergi.Opacity(img1, v)
		imgs = append(imgs, res)
	}

	return imgs
}