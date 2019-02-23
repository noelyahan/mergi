package loader

import (
	"image"
	"encoding/base64"
	"log"
	"bytes"
	"image/png"
	"strings"
	"image/jpeg"
	"errors"
	"fmt"
)

type b64Importer struct {
	data string
}

type b64Exporter struct {
	ext string
	img image.Image
	cb func(data string)
}

func NewBase64Importer(data string) Importer {
	return b64Importer{data}
}

func NewBase64Exporter(ext string, img image.Image, cb func(data string)) Exporter {
	return b64Exporter{ext, img, cb}
}

func (o b64Exporter) Export() (err error) {
	img := o.img
	if img == nil {
		return errors.New("Mergi found a invalid file ")
	}
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	if o.ext == "jpg" || o.ext == "jpeg" {
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	} else if o.ext == "png" {
		err = png.Encode(buf, img)
	}
	if err != nil {
		return errors.New("Sorry Mergi cannot encode the image")
	}

	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	if o.cb != nil {
		s := fmt.Sprintf("data:image/%s;base64,", o.ext)
		o.cb(s + str)
	}
	return
}

func (o b64Importer) Import() (image.Image, error) {
	ext := ""
	if strings.Contains(o.data, "png") {
		ext = "png"
	}else if strings.Contains(o.data, "jpeg") || strings.Contains(o.data, "jpg") {
		ext = "jpg"
	}
	imgStr := strings.Join(strings.Split(o.data, ",")[1:], "")
	decoded, err := base64.StdEncoding.DecodeString(imgStr)
	if err != nil {
		log.Printf("base64 decode error:", err)
		return nil, err
	}
	buf := bytes.NewReader(decoded)
	var img image.Image
	if ext == "png" {
		img, err = png.Decode(buf)
	}else if ext == "jpg" {
		img, err = jpeg.Decode(buf)
	}
	return img, err
}
