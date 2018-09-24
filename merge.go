package mergi

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"strings"
)

const (
	top    = "T"
	bottom = "B"
)

// Merge uses to merge multiple images according to given templates
//
// ex: "TT" - top, top - horizontal merge (2 images)
//
// ex: "TB" - top, bottom - vertical merge (2 images)
//
// for more merge techniques https://github.com/noelyahan/mergi/examples/merge
func Merge(template string, imgs []image.Image) (image.Image, error) {
	// return a nil
	tokens := strings.Split(template, "")
	if len(tokens) != len(imgs) {
		msg := fmt.Sprintf("Mergi founds a template=[%d] images=[%d] missmatch\n",
			len(tokens), len(imgs))
		return nil, errors.New(msg)
	}

	for i, v := range imgs {
		if v == nil {
			msg := fmt.Sprintf("Mergi founds a error image=[%v] on [%d]\n",
				v, i)
			return nil, errors.New(msg)
		}
	}
	first := imgs[0]
	yCount := getYAxisCount(template)
	xCount := 0
	// to get default height calculate first yCount imag Y
	height := 0
	width := 0
	for i, tmp := range imgs {
		if tmp == nil {
			msg := fmt.Sprintf("Mergi founds a error imageh=[%v]\n",
				tmp)
			return nil, errors.New(msg)
		}
		if i <= yCount {
			height += tmp.Bounds().Max.Y
		}
		if tokens[i] == top {
			width += tmp.Bounds().Max.X
		}
	}

	b := image.Rect(0, 0, width, height)
	resImage := image.NewRGBA(b)
	// T drawing
	draw.Draw(resImage, b, first, image.ZP, draw.Src)

	// t - top
	// b - bottom
	prev := top
	currX := first.Bounds().Max.X
	currY := first.Bounds().Max.Y
	isNotFirstColumn := false

	for i, t := range tokens {

		if i == 0 {
			// Because initial image will be always in top
			continue
		}

		first = imgs[i]

		if t == top {
			isNotFirstColumn = true
			currY = 0
			if xCount >= 1 {
				if len(imgs) > (i - 1) {
					currX += imgs[i-1].Bounds().Max.X
				}
			}
			xCount++
			offset := image.Pt(currX, 0)
			draw.Draw(resImage, first.Bounds().Add(offset), first, image.ZP, draw.Over)
		} else if t == bottom {

			offset := image.Pt(0, currY)

			if prev == top && isNotFirstColumn {
				offset = image.Pt(currX, currY)
			} else if isNotFirstColumn {
				offset = image.Pt(currX, currY)
			}

			if prev == top && !isNotFirstColumn {
				offset = image.Pt(0, currY)
			}

			if i >= yCount {
				isNotFirstColumn = true
			}

			draw.Draw(resImage, first.Bounds().Add(offset), first, image.ZP, draw.Over)
		}
		prev = t
		currY += first.Bounds().Max.Y
	}
	return resImage, nil
}

func getYAxisCount(template string) int {
	tCount := 0
	bCount := 0
	arr := strings.Split(template, "")

	for _, v := range arr {
		if tCount == 2 {
			break
		}
		if v == top {
			tCount++
		} else if v == bottom {
			bCount++
		}
	}
	return bCount
}
