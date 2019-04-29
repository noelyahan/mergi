package mergi_test

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/eximp"
	"image"
	"log"
)

// This example shows how to use the basic mergi crop
// Read the byte[] from the file and decode to a go standard image file
// and saving the import_export image file to disk
func ExampleCrop_croping() {
	// Get the image content by passing image path url or file path
	img, err := mergi.Import(eximp.NewFileImporter("testdata/tiger-2320819_960_720.jpg"))
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	// Now let's use the mergi's crop API

	// Set where to start the crop point
	cropStartPoint := image.Pt(0, 0)

	// Set crop width and height as point
	cropSize := image.Pt(495, 600)

	resultImage, err := mergi.Crop(img, cropStartPoint, cropSize)
	if err != nil {
		log.Fatalf("Mergi crop fails due to [%v]", err)
	}

	// Let's save the image
	err = mergi.Export(eximp.NewFileExporter(resultImage, "result.jpg"))
	if err != nil {
		log.Fatalf("failed to save: %s", err)
	}
}

// This example shows how to use the basic mergi merge for 2 images
// Read the byte[] from the file and decode to a go standard image file
// and saving the import_export image file to disk
func ExampleMerge_merging() {
	// Get the image content by passing image path url or file path
	img, err := mergi.Import(eximp.NewFileImporter("testdata/evraiki-2514543_240_180.jpg"))
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	// Now let's use the mergi's merge API

	// Lets define a template to merge 2 images horizontally
	template := "TT"
	images := []image.Image{img, img}

	resultImage, err := mergi.Merge(template, images)
	if err != nil {
		log.Fatalf("failed to merge: %s", err)
	}

	// Let's save the image
	err = mergi.Export(eximp.NewFileExporter(resultImage, "result.jpg"))
	if err != nil {
		log.Fatalf("failed to save: %s", err)
	}
}

// This example shows how to use the basic mergi resize
// Read the byte[] from the file and decode to a go standard image file
// and saving the import_export image file to disk
func ExampleResize_resizing() {
	// Get the image content by passing image path url or file path
	img, err := mergi.Import(eximp.NewFileImporter("testdata/evraiki-2514543_240_180.jpg"))
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	// Now let's use the mergi's resize API

	// Lets resize double of the given image's size
	width := uint(img.Bounds().Max.X * 2)
	height := uint(img.Bounds().Max.Y * 2)

	resultImage, err := mergi.Resize(img, width, height)
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}

	// Let's save the image
	err = mergi.Export(eximp.NewFileExporter(resultImage, "result.jpg"))
	if err != nil {
		log.Fatalf("failed to save: %s", err)
	}
}

// This example shows how to use the basic mergi watermark
// Read the byte[] from the file and decode to a go standard image file
// and saving the import_export image file to disk
func ExampleWatermark_watermarking() {
	// Get the image content by passing image path url or file path
	imgOriginal, err := mergi.Import(eximp.NewFileImporter("testdata/tiger-2320819_960_720.jpg"))
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	// Get the image content by passing image path url or file path
	imgWatermark, err := mergi.Import(eximp.NewFileImporter("./testdata/mergi_logo_watermark.png"))
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	// Now let's use the mergi's watermark API

	// Let's position the watermark left top corner
	p := image.Pt(0, 0)

	resultImage, err := mergi.Watermark(imgWatermark, imgOriginal, p)
	if err != nil {
		log.Fatalf("failed to watermark image: %s", err)
	}

	// Let's save the image
	err = mergi.Export(eximp.NewFileExporter(resultImage, "result.jpg"))
	if err != nil {
		log.Fatalf("failed to save: %s", err)
	}
}

// This example shows how to use the mergi Import
// Mergi provides the simplified api to load images via URL resource/file path
// Mergi.Import will return go standard image.Image
func ExampleImport_importing() {
	// Get the image content by passing image path url or file path
	imgFromLocal, err := mergi.Import(eximp.NewFileImporter("testdata/tiger-2320819_960_720.jpg"))
	if err != nil {
		log.Fatalf("failed to load via local: %s", err)
	}
	// Get the image content by passing image path url or file path
	// Reference: https://pixabay.com/en/woman-old-senior-female-elderly-1031000/
	imageFromURL, err := mergi.Import(eximp.NewURLImporter("https://cdn.pixabay.com/photo/2015/11/07/11/17/woman-1031000__340.jpg"))
	if err != nil {
		log.Fatalf("failed to load via url: %s", err)
	}

	// Now you can use this returned standard go image type inside Mergi APIS
	log.Println(imgFromLocal.Bounds(), imageFromURL.Bounds())
}

// This example shows how to use the mergi Export
// Mergi provides the simplified api to export the final result
// Mergi expects Exporter interface type
// Mergi supports base64, file, animation exporters
func ExampleExport_exporting() {
	// Get the image content by passing image path url or file path
	imgFromLocal, err := mergi.Import(eximp.NewFileImporter("testdata/tiger-2320819_960_720.jpg"))
	if err != nil {
		log.Fatalf("failed to load via local: %s", err)
	}
	// Get the image content by passing image path url or file path
	// Reference: https://pixabay.com/en/woman-old-senior-female-elderly-1031000/
	imageFromURL, err := mergi.Import(eximp.NewURLImporter("https://cdn.pixabay.com/photo/2015/11/07/11/17/woman-1031000__340.jpg"))
	if err != nil {
		log.Fatalf("failed to load via url: %s", err)
	}

	// Now you can use this returned standard go image type inside Mergi APIS
	log.Println(imgFromLocal.Bounds(), imageFromURL.Bounds())
}

// This example shows how to use the mergi Animate
// Mergi provides the simplified api to animate any given image array result
func ExampleAnimate_animating() {
	// Get the image content by passing image path url or file path
	img1, err := mergi.Import(eximp.NewFileImporter("testdata/avocado-3210885_960_720.jpg"))
	if err != nil {
		log.Fatalf("failed to load via local: %s", err)
	}
	img2, err := mergi.Import(eximp.NewFileImporter("testdata/cherry-3074284_960_720.jpg"))
	if err != nil {
		log.Fatalf("failed to load via local: %s", err)
	}

	imgFrames := []image.Image{img1, img2}
	// Image array and delay should be passed here
	animation, err := mergi.Animate(imgFrames, 20)
	if err != nil {
		log.Fatalf("failed to create animation: %s", err)
	}

	// Export final result via animation exporter
	mergi.Export(eximp.NewAnimationExporter(animation, "result.gif"))
}
