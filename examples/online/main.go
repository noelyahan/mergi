package main

import (
	"bytes"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"path/filepath"
	"strings"

	"github.com/dave/console"
	"github.com/dave/dropper"
	"github.com/dave/saver"
	"honnef.co/go/js/dom"
	"github.com/noelyahan/mergi"
	"golang.org/x/tools/go/analysis/passes/tests/testdata/src/divergent"
)

func main() {
	go run()
}

func run() {
	w := &console.Writer{}
	w.Message("Drag image here to compress")

	// initialise the drag+drop function with the github.com/dave/dropper package
	events := dropper.Initialise(dom.GetWindow().Document())

	// the dropper package creates a channel of events
	for ev := range events {
		switch ev := ev.(type) {
		case dropper.DropEvent:
			// accept a dropped file
			w.Message("Processing")
			// choose a filename
			name := strings.TrimSuffix(ev[0].Name(), filepath.Ext(ev[0].Name())) + ".jpg"

			// decode the image using the standard library image package
			img, _, err := image.Decode(ev[0].Reader())
			if err != nil {
				panic(err)
			}

			// do the mergi merge
			res, err := mergi.Merge("TBTB", []image.Image{img, img, img, img})
			if err != nil {
				panic(err)
			}

			// encode the image using the standard library jpeg package
			buf := &bytes.Buffer{}
			if err := jpeg.Encode(buf, res, &jpeg.Options{Quality: 50}); err != nil {
				panic(err)
			}

			// save the file as a browser download using the github.com/dave/saver package
			saver.Save(name, "image/jpeg", buf.Bytes())
			w.Message("Done")

		case dropper.EnterEvent:
			// drag event enters the page
			w.Message("Drop here")

		case dropper.LeaveEvent:
			// drag event exits the page
			w.Message("Drag image here to compress")

		}
	}
}
