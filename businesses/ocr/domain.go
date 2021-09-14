package ocr

import "context"

type OCRUsecase interface {
	GetImageTextFromImagePath(ctx context.Context, path string) (string, error)
	GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (string, error)
}
