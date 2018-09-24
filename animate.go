package mergi

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
)

// Animate uses go standard image.Image type array and delay input to returns a gif animation
//
// this type can be exported via mergi.Export function
//
// for more animate examples https://github.com/noelyahan/mergi/examples/animate
func Animate(imgs []image.Image, delay int) (gif.GIF, error) {
	for i, v := range imgs {
		if v == nil {
			msg := fmt.Sprintf("Mergi found a error image=[%v] on %d", v, i)
			return gif.GIF{}, errors.New(msg)
		}
	}
	delays := make([]int, 0)

	for i := 0; i < len(imgs); i++ {
		delays = append(delays, delay)
	}

	images := encodeImgPaletted(&imgs)

	return gif.GIF{
		Image: images,
		Delay: delays,
	}, nil
}

// Source Ref: https://github.com/ritchie46/GOPHY/blob/master/img2gif/img2gif.go#L38
// Thanks ritchie46
func encodeImgPaletted(images *[]image.Image) []*image.Paletted {
	// Gif options
	opt := gif.Options{}
	g := []*image.Paletted{}

	for _, im := range *images {
		b := bytes.Buffer{}
		// Write img2gif file to buffer.
		err := gif.Encode(&b, im, &opt)

		if err != nil {
			println(err)
		}
		// Decode img2gif file from buffer to img.
		img, err := gif.Decode(&b)

		if err != nil {
			println(err)
		}

		// Cast img.
		i, ok := img.(*image.Paletted)
		if ok {
			g = append(g, i)
		}
	}
	return g
}
