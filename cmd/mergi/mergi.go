package main

import (
	"errors"
	"fmt"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
	"image/gif"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getJobMap(flags []string) map[int][]string {
	m := make(map[int][]string, 0)
	for _, v := range flags {
		tags := strings.Split(v, "#")
		if len(tags) != 2 {
			return m
		}
		key, _ := strconv.Atoi(tags[0])
		val := tags[1]
		if val == flagFinal {
			continue
		}
		arr, ok := m[key]
		if !ok {
			tmp := make([]string, 0)
			tmp = append(tmp, val)
			m[key] = tmp
		} else {
			arr = append(arr, val)
			m[key] = arr
		}
	}
	return m
}

func getFilePaths(arr arrFlags) []string {
	var paths []string
	if len(arr) == 1 {
		paths = strings.Split(arr[0], " ")
	} else {
		paths = make([]string, 0)
		for _, p := range arr {
			paths = append(paths, p)
		}
	}
	return paths
}

func getPreProcessTemplatePaths(tokens, paths []string, arr arrFlags) []string {
	if len(tokens) != len(paths) {
		msg := fmt.Sprintf("mergi found a image template missmatch template size: %d image size: %d",
			len(tokens), len(paths))
		fmt.Println(msg)
		os.Exit(1)
		diff := len(tokens) - len(paths)
		for i := 0; i < diff; i++ {
			paths = append(paths, arr[0])
		}
	}
	return paths
}

func getImagesFromPaths(paths []string) ([]image.Image, error) {
	imgs := make([]image.Image, 0)
	var img image.Image
	var err error
	for _, p := range paths {
		if isValidURL(p) {
			img, err = mergi.Import(io.NewURLImporter(p))
		} else {
			img, err = mergi.Import(io.NewFileImporter(p))
		}

		if err != nil {
			msg := fmt.Sprintf("mergi expects jpg or png [%s] sorry ! :)", p)
			return imgs, errors.New(msg)
		}
		imgs = append(imgs, img)
	}
	return imgs, nil
}

func processFlaggedJobs(imgs []image.Image, template, animation string, jobs map[int][]string, cFlags, rFlags, wFlags arrFlags) mergiResult {
	cfCount := 0
	rfCount := 0
	wfCount := 0
	// process the large job also
	// sort jobs
	var keys []int
	for k := range jobs {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	tmpImgs := imgs
	for _, k := range keys {
		v := jobs[k]
		if k >= len(imgs) {
			break
		}
		for _, task := range v {
			if task == flagCrop {
				p1, p2 := getCropPoints(cFlags, cFlags[cfCount])
				timg, err := mergi.Crop(tmpImgs[k], p1, p2)
				if err != nil {
					fmt.Printf("Mergi cannot complete the crop: [%v]", err)
					continue
				}
				tmpImgs[k] = timg
				cfCount++
			} else if task == flagResize {
				w, h := getResizeWH(rFlags, rFlags[rfCount])
				timg, err := mergi.Resize(tmpImgs[k], w, h)
				if err != nil {
					fmt.Printf("Mergi cannot complete the resize: [%v]", err)
					continue
				}
				tmpImgs[k] = timg
				rfCount++
			} else if task == flagWatermark {
				img, x, y := getWatermarkImageXY(wFlags[wfCount])
				timg, err := mergi.Watermark(img, tmpImgs[k], image.Pt(x, y))
				if err != nil {
					fmt.Printf("Mergi cannot complete the watermark: [%v]", err)
					continue
				}
				tmpImgs[k] = timg
				wfCount++
			}
		}
	}

	var res image.Image
	var gif gif.GIF
	var err error

	skip := false
	if animation != "" {
		skip = true
		animType, delay := getAnimationParams(animation)
		if animType == sprite {
			gif, err = mergi.Animate(imgs, delay)
			if err != nil {
				fmt.Printf("Mergi cannot complete the animate: [%v]", err)
			}
		} else if animType == smooth {
			temp := ""
			tmpImgs = append(tmpImgs, tmpImgs[0])
			for i := 0; i < len(tmpImgs); i++ {
				temp += "T"
			}
			res, err := mergi.Merge(temp, tmpImgs)
			if err != nil {
				fmt.Printf("Mergi cannot complete the watermark: [%v]", err)
			}

			tmpSlides := make([]image.Image, 0)
			w := res.Bounds().Max.X / len(tmpImgs)
			size := image.Pt(w, res.Bounds().Max.Y)
			for x := 0; x < res.Bounds().Max.X-w; x += delay {
				croped, err := mergi.Crop(res, image.Pt(x, 0), size)
				if err != nil {
					fmt.Printf("Mergi cannot complete the crop: [%v]", err)
					continue
				}
				tmpSlides = append(tmpSlides, croped)
			}
			gif, err = mergi.Animate(tmpSlides, delay)
			if err != nil {
				fmt.Printf("Mergi cannot complete the animate: [%v]", err)
			}
		}
	} else {
		res, err = mergi.Merge(template, tmpImgs)
		if err != nil {
			fmt.Printf("Mergi cannot complete the watermark: [%v]", err)
		}
	}

	if skip {
		return mergiResult{res, gif}
	}

	// post-processing
	tasks, ok := jobs[len(imgs)+1]
	if ok {
		for _, task := range tasks {
			if task == flagCrop {
				p1, p2 := getCropPoints(cFlags, cFlags[cfCount])
				var err error
				res, err = mergi.Crop(res, p1, p2)
				if err != nil {
					fmt.Printf("Mergi cannot complete the crop: [%v]", err)
					continue
				}
				cfCount++
			} else if task == flagResize {
				w, h := getResizeWH(rFlags, rFlags[rfCount])
				res, err = mergi.Resize(res, w, h)
				if err != nil {
					fmt.Printf("Mergi cannot complete the resize: [%v]", err)
					continue
				}
				rfCount++
			} else if task == flagWatermark {
				img, x, y := getWatermarkImageXY(wFlags[wfCount])
				res, err = mergi.Watermark(img, res, image.Pt(x, y))
				if err != nil {
					fmt.Printf("Mergi cannot complete the watermark: [%v]", err)
					continue
				}
				wfCount++
			}
		}
	}
	return mergiResult{res, gif}
}
