package domain

import "context"

type Book struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	EditionCount int      `json:"edition_count"`
	Subjects     []string `json:"subjects"`
	Authors      []Author `json:"authors"`
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BookRepository interface {
	Index(ctx context.Context) ([]Book, error)
	FilterBySubject(ctx context.Context, subject string) ([]Book, error)
	FilterByID(ctx context.Context, id string) (Book, error)
}

type BookUsecase interface {
	Index(ctx context.Context, subject string) ([]Book, error)
}
