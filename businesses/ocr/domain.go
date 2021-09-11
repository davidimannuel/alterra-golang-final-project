package ocr

import "context"

type Domain struct {
	ImageText string
}

type Repository interface {
	GetImageTextFromImagePath(ctx context.Context, path string) (Domain, error)
	GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (Domain, error)
}

type Usecase interface {
	GetImageTextFromImagePath(ctx context.Context, path string) (Domain, error)
	GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (Domain, error)
}
