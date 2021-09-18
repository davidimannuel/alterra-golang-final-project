package ocr

import (
	"github.com/otiai10/gosseract/v2"
)

func GetImageTextFromImagePath(path string) (string, error) {
	ocr := gosseract.NewClient()
	defer ocr.Close()
	ocr.SetImage(path)
	return ocr.Text()
}

func GetImageTextFromImageBytes(bytes []byte) (string, error) {
	ocr := gosseract.NewClient()
	defer ocr.Close()
	ocr.SetImageFromBytes(bytes)
	return ocr.Text()
}
