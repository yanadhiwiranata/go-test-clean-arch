package usecase

import (
	"context"

	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
)

type BookUsecase struct {
	BookRepository domain.BookRepository
}

func NewBookUsecase(bookRepository domain.BookRepository) domain.BookUsecase {
	return &BookUsecase{
		BookRepository: bookRepository,
	}
}

func (a *BookUsecase) Index(ctx context.Context, subject string) ([]domain.Book, error) {
	var books []domain.Book
	var err error
	if len(subject) == 0 {
		books, err = a.BookRepository.Index(ctx)
	} else {
		books, err = a.BookRepository.FilterBySubject(ctx, subject)
	}

	if err != nil {
		return books, err
	}

	if len(books) < 1 {
		return books, domain.ErrNotFound
	}

	return books, nil
}
