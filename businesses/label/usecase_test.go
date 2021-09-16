package label_test

import (
	"context"
	"errors"
	"keep-remind-app/businesses/label"
	_labelMock "keep-remind-app/businesses/label/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errCase = errors.New("error_case")
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
			labelUsecase    label.LabelUsecase
		)
		labelUsecase = label.NewLabelUsecase(&labelRepository)
		labelData := []label.LabelDomain{
			{
				ID:     1,
				UserID: 1,
				Name:   "Label 1",
			},
			{
				ID:     2,
				UserID: 1,
				Name:   "Label 2",
			},
		}
		param := label.LabelParameter{}
		labelRepository.On("FindAll", context.Background(), &param).Return(labelData, nil).Once()

		result, err := labelUsecase.FindAll(context.Background(), &label.LabelParameter{})
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
			labelUsecase    label.LabelUsecase
		)
		labelUsecase = label.NewLabelUsecase(&labelRepository)
		labelData := []label.LabelDomain{
			{
				ID:     1,
				UserID: 1,
				Name:   "Label 1",
			},
			{
				ID:     2,
				UserID: 1,
				Name:   "Label 2",
			},
		}
		param := label.LabelParameter{}
		labelRepository.On("FindAll", context.Background(), &param).Return(labelData, errCase).Once()

		_, err := labelUsecase.FindAll(context.Background(), &label.LabelParameter{})
		assert.NotNil(t, err)

	})
}

func TestFindAllPagination(t *testing.T) {

	t.Run("Find All Pagination | Valid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
			labelUsecase    label.LabelUsecase
		)
		labelUsecase = label.NewLabelUsecase(&labelRepository)
		labelData := []label.LabelDomain{
			{
				ID:     1,
				UserID: 1,
				Name:   "Label 1",
			},
			{
				ID:     2,
				UserID: 1,
				Name:   "Label 2",
			},
		}
		param := label.LabelParameter{}
		labelRepository.On("FindAllPagination", context.Background(), &param).Return(labelData, 2, nil).Once()
		_, total, err := labelUsecase.FindAllPagination(context.Background(), &label.LabelParameter{})
		assert.Nil(t, err)
		assert.Equal(t, 2, total.TotalData)
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
			labelUsecase    label.LabelUsecase
		)
		labelUsecase = label.NewLabelUsecase(&labelRepository)
		labelData := []label.LabelDomain{
			{
				ID:     1,
				UserID: 1,
				Name:   "Label 1",
			},
			{
				ID:     2,
				UserID: 1,
				Name:   "Label 2",
			},
		}
		param := label.LabelParameter{}
		labelRepository.On("FindAllPagination", context.Background(), &param).Return(labelData, 2, errCase).Once()
		_, _, err := labelUsecase.FindAllPagination(context.Background(), &label.LabelParameter{})
		assert.NotNil(t, err)
	})
}

func TestFindOne(t *testing.T) {
	t.Run("Find One | Valid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
			labelUsecase    label.LabelUsecase
		)
		labelUsecase = label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			ID:     1,
			UserID: 1,
			Name:   "Label 1",
		}
		param := label.LabelParameter{}
		labelRepository.On("FindOne", context.Background(), &param).Return(labelData, nil).Once()
		result, err := labelUsecase.FindOne(context.Background(), &label.LabelParameter{})
		assert.Nil(t, err)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, 1, result.UserID)
		assert.Equal(t, "Label 1", result.Name)
	})

	t.Run("Find One | InValid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
			labelUsecase    label.LabelUsecase
		)
		labelUsecase = label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			ID:     0,
			UserID: 0,
			Name:   "",
		}
		param := label.LabelParameter{}
		labelRepository.On("FindOne", context.Background(), &param).Return(labelData, errCase).Once()
		_, err := labelUsecase.FindOne(context.Background(), &label.LabelParameter{})
		assert.NotNil(t, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Test Add | Valid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
		)
		labelUsecase := label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			UserID: 1,
			Name:   "Label 1",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		labelRepository.On("Add", ctx, &labelData).Return(1, nil).Once()
		result, err := labelUsecase.Add(ctx, &labelData)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Add | InValid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
		)
		labelUsecase := label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			UserID: 1,
			Name:   "Label 1",
		}
		ctx := context.WithValue(context.Background(), "user_id", 0)
		labelRepository.On("Add", ctx, &labelData).Return(1, errCase).Once()
		_, err := labelUsecase.Add(ctx, &labelData)
		assert.NotNil(t, err)
	})
	t.Run("Test Add | InValid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
		)
		labelUsecase := label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			UserID: 1,
			Name:   "Label 1",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		labelRepository.On("Add", ctx, &labelData).Return(1, errCase).Once()
		_, err := labelUsecase.Add(ctx, &labelData)
		assert.NotNil(t, err)
	})
}

func TestEdit(t *testing.T) {
	t.Run("Test Edit | Valid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
		)
		labelUsecase := label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			UserID: 1,
			Name:   "Label 1",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		labelRepository.On("Edit", ctx, &labelData).Return(nil).Once()
		err := labelUsecase.Edit(ctx, &labelData)
		assert.Nil(t, err)
	})

	t.Run("Test Edit | InValid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
		)
		labelUsecase := label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			UserID: 1,
			Name:   "Label 1",
		}
		ctx := context.WithValue(context.Background(), "user_id", 0)
		labelRepository.On("Edit", ctx, &labelData).Return(errCase).Once()
		err := labelUsecase.Edit(ctx, &labelData)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Delete | Valid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
		)
		labelUsecase := label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			UserID: 1,
			Name:   "Label 1",
		}
		labelRepository.On("Delete", context.Background(), 0).
			Return(nil).Once()

		err := labelUsecase.Delete(context.Background(), labelData.ID)
		assert.Nil(t, err)
	})

	t.Run("Test Delete | InValid", func(t *testing.T) {
		var (
			labelRepository _labelMock.LabelRepository
		)
		labelUsecase := label.NewLabelUsecase(&labelRepository)
		labelData := label.LabelDomain{
			UserID: 1,
			Name:   "Label 1",
		}
		labelRepository.On("Delete", context.Background(), 0).Return(errCase).Once()
		err := labelUsecase.Delete(context.Background(), labelData.ID)
		assert.NotNil(t, err)
	})
}
