package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	_bookUsecase "github.com/yanadhiwiranata/go-test-clean-arch/book/usecase"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	mocks "github.com/yanadhiwiranata/go-test-clean-arch/mocks/domain"
)

func TestIndexFilterBySubject(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	subject := "Bibliography"
	mockAuthor1 := domain.Author{ID: "/authors/OL9388A", Name: "Yan"}
	mockBooks := []domain.Book{
		{
			ID:           "/works/OL362427W",
			Title:        "title 1",
			EditionCount: 20,
			Authors:      []domain.Author{mockAuthor1},
			Subjects:     []string{"asd", subject},
		},
	}
	mockBookRepo.On("FilterBySubject", mock.Anything, subject).Return(mockBooks, nil)
	bookUsecase := _bookUsecase.NewBookUsecase(mockBookRepo)
	books, err := bookUsecase.Index(context.Background(), subject)
	assert.NoError(t, err)
	assert.NotEmpty(t, books)

	mockBookRepo = new(mocks.BookRepository)
	mockBookRepo.On("FilterBySubject", mock.Anything, subject).Return([]domain.Book{}, nil)
	bookUsecase = _bookUsecase.NewBookUsecase(mockBookRepo)
	books, err = bookUsecase.Index(context.Background(), subject)
	assert.Equal(t, domain.ErrNotFound, err)
	assert.Empty(t, books)
}
