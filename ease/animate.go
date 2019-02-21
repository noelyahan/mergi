package ease

import (
	"math"
	"image"
)

func mapVlaues(value, start1, stop1, start2, stop2 float64) float64 {
	return start2 + (stop2-start2)*((value-start1)/(stop1-start1))
}

func AnimatePoints(move Move, from, to image.Point, fps float64) (arr []image.Point) {
	myMin := math.MaxFloat64
	myArr := make([]float64, 0)
	mxArr := make([]float64, 0)
	for x := float64(to.X); x > 0.0; x -= fps {
		t := mapVlaues(x, float64(from.X), float64(to.X), 0, 1)
		my := move(t)

		if my < myMin {
			myMin = my
		}
		myArr = append(myArr, my)
		mxArr = append(mxArr, x)
	}
	for i, my := range myArr {
		v := mapVlaues(my, myMin, 1, float64(from.Y), float64(to.Y))
		arr = append(arr, image.Pt(int(mxArr[i]), int(v)))
	}

	return
}
