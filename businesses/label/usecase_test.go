package label_test

import (
	"context"
	"errors"
	"keep-remind-app/businesses"
	"keep-remind-app/businesses/label"
	_labelMock "keep-remind-app/businesses/label/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	labelRepository _labelMock.LabelRepository
	labelUsecase    label.LabelUsecase
	labelDomain     label.LabelDomain
	errCase         = errors.New("error_case")
)

func TestMain(m *testing.M) {
	labelUsecase = label.NewLabelUsecase(&labelRepository)
	m.Run()
}

func TestFindAll(t *testing.T) {
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
	t.Run("Find All | Valid", func(t *testing.T) {
		labelRepository.On("FindAll", mock.AnythingOfType("context.Context"), mock.AnythingOfType("&label.LabelParameter{}")).
			Return(labelData, nil).Once()

		result, err := labelUsecase.FindAll(context.Background(), &label.LabelParameter{})
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		labelRepository.On("FindAll", mock.AnythingOfType("context.Context"), mock.AnythingOfType("&label.LabelParameter{}")).
			Return(labelData, errCase).Once()

		_, err := labelUsecase.FindAll(context.Background(), &label.LabelParameter{})
		assert.NotNil(t, err)

	})
}

func TestFindAllPagination(t *testing.T) {
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

	t.Run("Find All Pagination | Valid", func(t *testing.T) {
		labelRepository.On("FindAllPagination", mock.AnythingOfType("context.Context"), mock.AnythingOfType("&label.LabelParameter{}")).
			Return(labelData, nil, businesses.Pagination{Page: 1, PerPage: 10, TotalData: 20}).Once()

		result, pagination, err := labelUsecase.FindAllPagination(context.Background(), &label.LabelParameter{})
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, pagination.TotalData, 1)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		labelRepository.On("FindAllPagination", mock.AnythingOfType("context.Context"), mock.AnythingOfType("&label.LabelParameter{}")).
			Return(labelData, errCase, businesses.Pagination{Page: 1, PerPage: 10, TotalData: 20}).Once()

		_, _, err := labelUsecase.FindAllPagination(context.Background(), &label.LabelParameter{})
		assert.NotNil(t, err)

	})
}

func TestFindOne(t *testing.T) {
	labelData := label.LabelDomain{
		ID:     1,
		UserID: 1,
		Name:   "Label 1",
	}

	t.Run("Find One | Valid", func(t *testing.T) {
		labelRepository.On("FindOne", mock.AnythingOfType("context.Context"), mock.AnythingOfType("&label.LabelParameter{}")).
			Return(labelData, nil).Once()

		result, err := labelUsecase.FindOne(context.Background(), &label.LabelParameter{})
		assert.Nil(t, err)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, 1, result.UserID)
		assert.Equal(t, "Label 1", result.Name)
	})

	t.Run("Find One | InValid", func(t *testing.T) {
		labelRepository.On("FindOne", mock.AnythingOfType("context.Context"), mock.AnythingOfType("&label.LabelParameter{}")).
			Return([]label.LabelDomain{labelDomain}, errCase).Once()

		_, err := labelUsecase.FindOne(context.Background(), &label.LabelParameter{})
		assert.NotNil(t, err)
	})
}

func TestAdd(t *testing.T) {
	labelData := label.LabelDomain{
		ID:     1,
		UserID: 1,
		Name:   "Label 1",
	}

	t.Run("Test Add | Valid", func(t *testing.T) {
		labelRepository.On("Add", mock.AnythingOfType("context.Context"), mock.AnythingOfType("labelData")).
			Return(labelData, nil).Once()

		result, err := labelUsecase.Add(context.Background(), &labelData)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Add | InValid", func(t *testing.T) {
		labelRepository.On("Add", mock.AnythingOfType("context.Context"), mock.AnythingOfType("labelData")).
			Return(labelData, errCase).Once()

		_, err := labelUsecase.Add(context.Background(), &labelData)
		assert.NotNil(t, err)
	})
}

func TestEdit(t *testing.T) {
	labelData := label.LabelDomain{
		ID:     1,
		UserID: 1,
		Name:   "Label 1",
	}

	t.Run("Test Edit | Valid", func(t *testing.T) {
		labelRepository.On("Edit", mock.AnythingOfType("context.Context"), mock.AnythingOfType("labelData")).
			Return(nil).Once()

		err := labelUsecase.Edit(context.Background(), &labelData)
		assert.Nil(t, err)
	})

	t.Run("Test Edit | InValid", func(t *testing.T) {
		labelRepository.On("Edit", mock.AnythingOfType("context.Context"), mock.AnythingOfType("labelData")).
			Return(errCase).Once()

		err := labelUsecase.Edit(context.Background(), &labelData)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	labelData := label.LabelDomain{
		ID:     1,
		UserID: 1,
		Name:   "Label 1",
	}

	t.Run("Test Delete | Valid", func(t *testing.T) {
		labelRepository.On("Delete", mock.AnythingOfType("context.Context"), mock.AnythingOfType("labelData")).
			Return(nil).Once()

		err := labelUsecase.Delete(context.Background(), labelData.ID)
		assert.Nil(t, err)
	})

	t.Run("Test Delete | InValid", func(t *testing.T) {
		labelRepository.On("Delete", mock.AnythingOfType("context.Context"), mock.AnythingOfType("labelData")).
			Return(errCase).Once()

		err := labelUsecase.Delete(context.Background(), labelData.ID)
		assert.NotNil(t, err)
	})
}
