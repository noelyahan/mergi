package ease

import (
	"math"
	"image"
	"github.com/noelyahan/mergi"
)

func mapValues(value, istart, istop, ostart, ostop float64) float64 {
	return ostart + (ostop-ostart)*((value-istart)/(istop-istart))
}

func Ease(value, start, end float64, ease EaseType) float64 {
	t := mapValues(value, start, end, 0.1, 1)
	v := ease(t)
	return mapValues(v, 0, 1, start, end)
}

func Ease1(img1, img2 image.Image, from, to image.Point, speed float64, ease EaseType) []image.Image {
	myMin := math.MaxFloat64
	myArr := make([]float64, 0)
	mxArr := make([]float64, 0)
	frames := make([]image.Image, 0)
	for x := float64(from.X); x < float64(to.X); x += speed {
		t := mapValues(x, float64(from.X), float64(to.X), 0, 1)
		my := ease(t)
		if my < myMin {
			myMin = my
		}
		myArr = append(myArr, my)
		mxArr = append(mxArr, x)
	}
	for i, my := range myArr {
		y := mapValues(my, myMin, 1, float64(from.Y), float64(to.Y))
		img, _ := mergi.Watermark(img2, img1, image.Pt(int(mxArr[i]), int(y)))
		frames = append(frames, img)
	}
	return frames
}

func AnimatePoints(move EaseType, from, to image.Point, fps float64) (arr []image.Point) {
	myMin := math.MaxFloat64
	myArr := make([]float64, 0)
	mxArr := make([]float64, 0)
	for x := float64(to.X); x > 0.0; x -= fps {
		t := mapValues(x, float64(from.X), float64(to.X), 0, 1)
		my := move(t)

		if my < myMin {
			myMin = my
		}
		myArr = append(myArr, my)
		mxArr = append(mxArr, x)
	}
	for i, my := range myArr {
		v := mapValues(my, myMin, 1, float64(from.Y), float64(to.Y))
		arr = append(arr, image.Pt(int(mxArr[i]), int(v)))
	}

	return
}
