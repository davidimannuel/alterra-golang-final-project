package ocr

import (
	"github.com/otiai10/gosseract/v2"
)

type OCR struct {
	Client *gosseract.Client
}

func InitOCR() *gosseract.Client {
	return gosseract.NewClient()
}

func (ocr *OCR) GetImageTextFromImagePath(path string) (string, error) {
	ocr.Client.SetImage(path)
	return ocr.Client.Text()
}

func (ocr *OCR) GetImageTextFromImageBytes(bytes []byte) (string, error) {
	ocr.Client.SetImageFromBytes(bytes)
	return ocr.Client.Text()
}
