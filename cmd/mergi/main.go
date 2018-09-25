package main

import (
	"flag"
	"fmt"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"image"
	"image/gif"
	"os"
	"strings"
)

type mergiResult struct {
	image     image.Image
	animation gif.GIF
}

const (
	version       = "1.0.0"
	flagImage     = "image"
	flagCrop      = "crop"
	flagResize    = "resize"
	flagWatermark = "watermark"
	flagFinal     = "final"
	sprite        = "sprite"
	smooth        = "smooth"
)

func main() {
	var iFlags arrFlags
	var cFlags arrFlags
	var rFlags arrFlags
	var wFlags arrFlags

	const (
		templateMsg = "Enter a merge template string ex: TBTBTB"
		imagesMsg   = "Enter images that want to merge ex: /path/img1 or url"
		resizeMsg   = "Enter resize width and height of the output ex: 100 200"
		cropMsg     = "Enter crop points and height and width ex: x y w h"
		wmMsg       = "Enter watermark image and points to place it, [-r w h] is optional  ex: /path/img -r w h x y"
		animMsg     = "Enter animation type=[sprite, slide] and the delay to get mergi gif animation ex: smooth 10"
		outMsg      = "Enter image outputs file ex: out.png or out.jpg"
		finalMsg    = "Enter true if you want to process the final output"
	)

	flag.Usage = setFlagUsage()

	template := flag.String("t", "T", templateMsg)
	animation := flag.String("a", "", animMsg)
	out := flag.String("o", "out.png", outMsg)

	flag.Var(&iFlags, "i", imagesMsg)
	flag.Var(&rFlags, "r", resizeMsg)
	flag.Var(&cFlags, "c", cropMsg)
	flag.Var(&wFlags, "w", wmMsg)

	var final string
	flag.StringVar(&final, "f", "false", finalMsg)

	flag.Parse()

	if len(*animation) > 0 && *out == "out.png" {
		*out = "out.gif"
	}
	// input validation + data preparation
	paths := getFilePaths(iFlags)
	tokens := strings.Split(*template, "")

	newPaths := getPreProcessTemplatePaths(tokens, paths, iFlags)

	if len(newPaths) != len(paths) {
		paths = newPaths
	}

	// pre-load images
	imgs, err := getImagesFromPaths(paths)
	if err != nil {
		fmt.Println(err)
		return
	}

	// validate the job and images
	tasks := getFlagOrder(os.Args)
	jobs := getJobMap(tasks)

	res := processFlaggedJobs(imgs, *template, *animation, jobs, cFlags, rFlags, wFlags)

	if res.image != nil {
		mergi.Export(loader.NewFileExporter(res.image, *out))
	} else {
		mergi.Export(loader.NewAnimationExporter(res.animation, *out))
	}

	msg := fmt.Sprintf("mergi success !, final image saved to %s", *out)
	fmt.Println(msg)
}

func setFlagUsage() func() {
	return func() {
		banner := `
 ╔╦╗╔═╗╦═╗╔═╗╦
 ║║║║╣ ╠╦╝║ ╦║
 ╩ ╩╚═╝╩╚═╚═╝╩
 let's go & make imaging fun
 http://mergi.io
 version %s

`
		fmt.Fprintf(os.Stderr, banner, version)
		flag.PrintDefaults()
	}
}
