package cache_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	_repo "github.com/yanadhiwiranata/go-test-clean-arch/book/repository/cache"
)

const cacheName = "book"

func TestIndex(t *testing.T) {
	bRepo := _repo.NewCacheBookRepository()
	books, _ := bRepo.Index(context.Background())
	assert.NotEmpty(t, books)
	assert.NotEmpty(t, books[0].Authors)
	assert.NotEmpty(t, books[0].Subjects)
	assert.NotEmpty(t, books[0].ID)
}

func TestFilterBySubject(t *testing.T) {
	bRepo := _repo.NewCacheBookRepository()
	books, _ := bRepo.FilterBySubject(context.Background(), "Bibliography")
	assert.NotEmpty(t, books)

	books, _ = bRepo.FilterBySubject(context.Background(), "Bibliography Empty")
	assert.Empty(t, books)
}

func TestFilterByID(t *testing.T) {
	bRepo := _repo.NewCacheBookRepository()
	books, _ := bRepo.FilterByID(context.Background(), "/works/OL362427W")
	assert.NotEmpty(t, books)

	books, _ = bRepo.FilterByID(context.Background(), "ad123")
	assert.Empty(t, books)
}
