package note_test

import (
	"context"
	"errors"
	"keep-remind-app/businesses/label"
	labelDomain "keep-remind-app/businesses/label"
	_labelMock "keep-remind-app/businesses/label/mocks"
	"keep-remind-app/businesses/note"
	_noteMock "keep-remind-app/businesses/note/mocks"
	ocrDomain "keep-remind-app/businesses/ocr"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errCase         = errors.New("error_case")
	ocrUsecase      ocrDomain.OCRUsecase
	labelUsecase    labelDomain.LabelUsecase
	labelRepository _labelMock.LabelRepository
)

func TestMain(m *testing.M) {
	ocrUsecase = ocrDomain.NewOCRUsecase()
	labelUsecase = labelDomain.NewLabelUsecase(&labelRepository)
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := []note.NoteDomain{
			{
				ID:     1,
				UserID: 1,
				Title:  "test",
				Note:   "test",
			},
		}
		param := note.NoteParameter{}
		noteRepository.On("FindAll", context.Background(), &param).Return(noteData, nil).Once()

		result, err := noteUsecase.FindAll(context.Background(), &param)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := []note.NoteDomain{
			{
				ID:     1,
				UserID: 1,
				Title:  "test",
				Note:   "test",
			},
		}
		param := note.NoteParameter{}
		noteRepository.On("FindAll", context.Background(), &param).Return(noteData, errCase).Once()
		_, err := noteUsecase.FindAll(context.Background(), &param)
		assert.NotNil(t, err)

	})
}

func TestFindAllPagination(t *testing.T) {
	t.Run("Find All Pagination | Valid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := []note.NoteDomain{
			{
				ID:     1,
				UserID: 1,
				Title:  "test",
				Note:   "test",
			},
		}
		param := note.NoteParameter{}
		noteRepository.On("FindAllPagination", context.Background(), &param).Return(noteData, 10, nil).Once()

		result, _, err := noteUsecase.FindAllPagination(context.Background(), &param)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All Pagination | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := []note.NoteDomain{
			{
				ID:     1,
				UserID: 1,
				Title:  "test",
				Note:   "test",
			},
		}
		param := note.NoteParameter{}
		noteRepository.On("FindAllPagination", context.Background(), &param).Return(noteData, 10, errCase).Once()

		_, _, err := noteUsecase.FindAllPagination(context.Background(), &param)
		assert.NotNil(t, err)
	})
}

func TestFindOne(t *testing.T) {
	t.Run("Find One | Valid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		param := note.NoteParameter{}
		noteRepository.On("FindOne", context.Background(), &param).Return(noteData, nil).Once()

		result, err := noteUsecase.FindOne(context.Background(), &param)
		assert.Nil(t, err)
		assert.Equal(t, 1, result.ID)
	})

	t.Run("Find One | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		param := note.NoteParameter{}
		noteRepository.On("FindOne", context.Background(), &param).Return(noteData, errCase).Once()

		_, err := noteUsecase.FindOne(context.Background(), &param)
		assert.NotNil(t, err)
	})
}
func TestAdd(t *testing.T) {
	t.Run("Add | Valid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		noteRepository.On("Add", ctx, &noteData).Return(1, nil).Once()
		result, err := noteUsecase.Add(ctx, &noteData)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Add | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
			Labels: []note.LabelDomain{
				{
					ID: 0,
				},
			},
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		noteRepository.On("Add", ctx, &noteData).Return(0, errCase).Once()
		labelRepository.On("FindOne", ctx, &label.LabelParameter{}).Return(label.LabelDomain{
			ID:     1,
			Name:   "label",
			UserID: 1,
		}, nil)
		result, err := noteUsecase.Add(ctx, &noteData)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("Add | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		noteRepository.On("Add", ctx, &noteData).Return(0, errCase).Once()
		result, err := noteUsecase.Add(ctx, &noteData)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("Add | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 0)
		noteRepository.On("Add", ctx, &noteData).Return(0, errCase).Once()
		result, err := noteUsecase.Add(ctx, &noteData)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

// func TestAddWithImageBytes(t *testing.T) {
// 	t.Run("TestAddWithImageBytes | Valid", func(t *testing.T) {
// 		var (
// 			noteRepository _noteMock.NoteRepository
// 			noteUsecase    note.NoteUsecase
// 		)
// 		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
// 		bytes, _ := ioutil.ReadFile("../../files/test_note.png")
// 		noteData := note.NoteDomain{
// 			ID:     1,
// 			UserID: 1,
// 			Title:  "test",
// 		}
// 		ctx := context.WithValue(context.Background(), "user_id", 1)
// 		noteRepository.On("Add", ctx, &noteData).Return(1, nil).Once()
// 		_, err := noteUsecase.AddWithImageBytes(ctx, "test", bytes)
// 		assert.Nil(t, err)
// 		// assert.Equal(t, 1, result)
// 	})
// }

func TestEdit(t *testing.T) {
	t.Run("Edit | Valid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		noteRepository.On("Edit", ctx, &noteData).Return(nil).Once()
		err := noteUsecase.Edit(ctx, &noteData)
		assert.Nil(t, err)
	})

	t.Run("Edit | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 1)
		noteRepository.On("Edit", ctx, &noteData).Return(errCase).Once()
		err := noteUsecase.Edit(ctx, &noteData)
		assert.NotNil(t, err)
	})

	t.Run("Edit | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		noteData := note.NoteDomain{
			ID:     1,
			UserID: 1,
			Title:  "test",
			Note:   "test",
		}
		ctx := context.WithValue(context.Background(), "user_id", 0)
		noteRepository.On("Edit", ctx, &noteData).Return(errCase).Once()
		err := noteUsecase.Edit(ctx, &noteData)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		ctx := context.WithValue(context.Background(), "user_id", 1)
		noteRepository.On("Delete", ctx, 1).Return(nil).Once()
		err := noteUsecase.Delete(ctx, 1)
		assert.Nil(t, err)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		var (
			noteRepository _noteMock.NoteRepository
			noteUsecase    note.NoteUsecase
		)
		noteUsecase = note.NewNoteUsecase(&noteRepository, ocrUsecase, labelUsecase)
		ctx := context.WithValue(context.Background(), "user_id", 1)
		noteRepository.On("Delete", ctx, 1).Return(errCase).Once()
		err := noteUsecase.Delete(ctx, 1)
		assert.NotNil(t, err)
	})
}
