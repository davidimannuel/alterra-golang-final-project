package ocr_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/otiai10/gosseract/v2"
)

func TestOCR(t *testing.T) {
	client := gosseract.NewClient()
	defer client.Close()
	// client.SetImage("test.png")
	imgBytes, _ := ioutil.ReadFile("test.png")
	client.SetImageFromBytes(imgBytes)
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
