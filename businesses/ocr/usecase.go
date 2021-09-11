package ocr

import "context"

type ocrUsecase struct {
	repository Repository
}

func NewUsecase(repository Repository) Usecase {
	return &ocrUsecase{
		repository: repository,
	}
}

func (uc ocrUsecase) GetImageTextFromImagePath(ctx context.Context, path string) (res Domain, err error) {
	res, err = uc.repository.GetImageTextFromImagePath(ctx, path)
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc ocrUsecase) GetImageTextFromImageBytes(ctx context.Context, bytes []byte) (res Domain, err error) {
	res, err = uc.repository.GetImageTextFromImageBytes(ctx, bytes)
	if err != nil {
		return res, err
	}
	return res, err
}
