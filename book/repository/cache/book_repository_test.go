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
