package note

import "context"

type noteUsecase struct {
	repository Repository
}

func NewUsecase(repository Repository) Usecase {
	return &noteUsecase{
		repository: repository,
	}
}

func (uc noteUsecase) Add(ctx context.Context, data *Domain) (Domain, error) {
	result, err := uc.repository.Add(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	return result, err
}
