package ocr

import "context"

type Domain struct {
	ImageText string
}

type Usecase interface {
	GetImageTextFromImagePath(ctx context.Context, path string) (Domain, error)
	GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (Domain, error)
}
