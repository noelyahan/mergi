package main

import (
	"fmt"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"image"
	"net/url"
	"strconv"
	"strings"
)

type arrFlags []string

func (i *arrFlags) String() string {
	return ""
}

func (i *arrFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func getFlagOrder(arr []string) []string {
	res := make([]string, 0)
	imgCount := -1
	appendInt := func(i int, s string) string {
		return strconv.Itoa(i) + "#" + s
	}
	for _, v := range arr {
		v = strings.Replace(v, "-", "", -1)
		if v == "c" || v == flagCrop {
			res = append(res, appendInt(imgCount, flagCrop))
		} else if v == "r" || v == flagResize {
			res = append(res, appendInt(imgCount, flagResize))
		} else if v == "w" || v == flagWatermark {
			res = append(res, appendInt(imgCount, flagWatermark))
		} else if v == "f" || v == flagFinal {
			imgCount++
			res = append(res, appendInt(imgCount, flagFinal))
			imgCount++
		} else if v == "i" || v == flagImage {
			imgCount++
		}
	}
	return res
}

func getWatermarkImageXY(wmStr string) (image.Image, int, int) {
	wmImgXY := strings.Split(wmStr, " ")

	if len(wmImgXY) == 3 {
		x, _ := strconv.Atoi(wmImgXY[1])
		y, _ := strconv.Atoi(wmImgXY[2])

		var img image.Image
		var err error
		if isValidURL(wmImgXY[0]) {
			img, err = mergi.Import(loader.NewURLImporter(wmImgXY[0]))
		} else {
			img, err = mergi.Import(loader.NewFileImporter(wmImgXY[0]))
		}

		if err != nil {
			return nil, 0, 0
		}

		return img, x, y
	}
	x, _ := strconv.Atoi(wmImgXY[4])
	y, _ := strconv.Atoi(wmImgXY[5])

	w, _ := strconv.Atoi(wmImgXY[2])
	h, _ := strconv.Atoi(wmImgXY[3])

	var img image.Image
	var err error
	if isValidURL(wmImgXY[0]) {
		img, err = mergi.Import(loader.NewURLImporter(wmImgXY[0]))
	} else {
		img, err = mergi.Import(loader.NewFileImporter(wmImgXY[0]))
	}

	if err != nil {
		return nil, 0, 0
	}
	img, err = mergi.Resize(img, uint(w), uint(h))
	if err != nil {
		fmt.Printf("Mergi cannot complete the resize: [%v]", err)
	}
	return img, x, y
}

func getResizeWH(rFlags arrFlags, resizeStr string) (uint, uint) {
	resizeWH := strings.Split(resizeStr, " ")

	if len(rFlags) != 0 {
		if len(resizeWH) == 2 {
			rw, _ := strconv.Atoi(resizeWH[0])
			rh, _ := strconv.Atoi(resizeWH[1])
			return uint(rw), uint(rh)
		}
	}
	fmt.Println("mergi wants you to enter width and height", resizeStr, ":)")
	return 0, 0
}

func getCropPoints(cFlags arrFlags, cropStr string) (image.Point, image.Point) {
	cropXYXY := strings.Split(cropStr, " ")
	if len(cropXYXY) == 4 {
		x1, _ := strconv.Atoi(cropXYXY[0])
		y1, _ := strconv.Atoi(cropXYXY[1])
		x2, _ := strconv.Atoi(cropXYXY[2])
		y2, _ := strconv.Atoi(cropXYXY[3])
		p1 := image.Pt(x1, y1)
		p2 := image.Pt(x2, y2)
		return p1, p2
	}
	fmt.Println("mergi wants you to enter point X and point Y along with width and height to crop", cFlags[0], ":)")
	return image.Pt(-1, -1), image.Pt(-1, -1)
}

func getAnimationParams(animStr string) (string, int) {
	splitTypeDelay := strings.Split(animStr, " ")
	if len(splitTypeDelay) == 2 {
		animType := splitTypeDelay[0]
		delay, _ := strconv.Atoi(splitTypeDelay[1])
		return animType, delay
	}
	return "", 0
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true
}
