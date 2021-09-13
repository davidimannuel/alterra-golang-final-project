package ocr

import "context"

type Usecase interface {
	GetImageTextFromImagePath(ctx context.Context, path string) (string, error)
	GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (string, error)
}
