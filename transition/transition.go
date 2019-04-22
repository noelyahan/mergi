package transition

import (
	"image"
	"github.com/noelyahan/mergi"
)

func mapValues(value, start1, stop1, start2, stop2 float64) float64 {
	return start2 + (stop2-start2)*((value-start1)/(stop1-start1))
}

func FadeIn(img1, img2 image.Image, time, speed float64) []image.Image {
	frames := make([]image.Image, 0)
	for i := time; i > 0; i -= speed {
		v := mapValues(i, time, 0, 1, 0)
		res, _ := mergi.Opacity(img1, v)
		frames = append(frames, res)
	}

	for i := 0.0; i < time; i += speed {
		v := mapValues(i, 0, time, 0, 1)
		res, _ := mergi.Opacity(img2, v)
		frames = append(frames, res)
	}
	return frames
}

func FadeOut(img1, img2 image.Image, time, speed float64) []image.Image {
	frames := make([]image.Image, 0)

	for i := 0.0; i < time; i += speed {
		v := mapValues(i, 0, time, 0, 1)
		res, _ := mergi.Opacity(img2, v)
		frames = append(frames, res)
	}

	for i := time; i > 0; i -= speed {
		v := mapValues(i, time, 0, 1, 0)
		res, _ := mergi.Opacity(img1, v)
		frames = append(frames, res)
	}

	return frames
}

func Cover(img1, img2 image.Image, time, speed float64) []image.Image {
	frames := make([]image.Image, 0)

	for i := 0.0; i < time; i += speed {
		// left x
		lx := mapValues(i, 0, time, 0, float64(img1.Bounds().Max.X) + time)
		// right x
		rx := float64(img1.Bounds().Max.X) - lx

		// crop image 1
		cimg1, _ := mergi.Crop(img1, image.Pt(0, 0), image.Pt(int(lx), img1.Bounds().Max.Y))
		// crop image 2
		cimg2, _ := mergi.Crop(img2, image.Pt(0, 0), image.Pt(int(rx), img2.Bounds().Max.Y))

		// merged crop image 1 + crop image 2
		mimg, _ := mergi.Merge("TT", []image.Image{cimg2, cimg1})

		// add to transition frames
		frames = append(frames, mimg)
	}

	return frames
}

func Split(img1, img2 image.Image, time, speed float64) []image.Image {
	frames := make([]image.Image, 0)

	// img1 -> nature
	part := float64(img2.Bounds().Max.X) / 2

	// crop image 1 for left cover
	cimg1, _ := mergi.Crop(img1, image.ZP, image.Pt(int(part), img1.Bounds().Max.Y))

	// crop image 2 for left cover
	cimg2, _ := mergi.Crop(img1, image.Pt(int(part), 0), image.Pt(img1.Bounds().Max.X, img1.Bounds().Max.Y))


	for i := 0.0; i < time; i += speed {
		// left x
		lx := mapValues(i, 0, time, 0, part)

		// new crop image 3 for left cover movement
		cimg3, _ := mergi.Crop(cimg1, image.Pt(int(lx), 0), image.Pt(cimg1.Bounds().Max.X, cimg1.Bounds().Max.Y))

		// watermark image 1 for left
		wm1, _ := mergi.Watermark(cimg3, img2, image.ZP)

		// watermark image 2 for right
		wm2, _ := mergi.Watermark(cimg2, wm1, image.Pt(int(part + lx), 0))

		// add to transition frames
		frames = append(frames, wm2)
	}

	return frames
}