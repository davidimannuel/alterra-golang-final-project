package ocr

import (
	"context"
	"keep-remind-app/businesses/ocr"

	"github.com/otiai10/gosseract/v2"
)

type ocrRepository struct {
	Client *gosseract.Client
}

func NewOCRRepository(client *gosseract.Client) ocr.Repository {
	return &ocrRepository{
		Client: client,
	}
}

func (repo *ocrRepository) GetImageTextFromImagePath(ctx context.Context, path string) (ocr.Domain, error) {
	repo.Client.SetImage(path)
	text, err := repo.Client.Text()
	if err != nil {
		return ocr.Domain{}, err
	}
	return ocr.Domain{ImageText: text}, nil
}
func (repo *ocrRepository) GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (ocr.Domain, error) {
	repo.Client.SetImageFromBytes(bytes)
	text, err := repo.Client.Text()
	if err != nil {
		return ocr.Domain{}, err
	}
	repo.Client.Close()
	return ocr.Domain{ImageText: text}, nil
}
