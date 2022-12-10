package usecase

import (
	"context"

	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
)

type BookUsecase struct {
	bookRepository domain.BookRepository
}

func NewBookUsecase(bookRepository domain.BookRepository) domain.BookUsecase {
	return &BookUsecase{
		bookRepository: bookRepository,
	}
}

func (a *BookUsecase) Index(ctx context.Context, subject string) ([]domain.Book, error) {
	books, err := a.bookRepository.FilterBySubject(ctx, subject)
	if err != nil {
		return books, err
	}

	if len(books) < 1 {
		return books, domain.ErrNotFound
	}

	return books, nil
}
