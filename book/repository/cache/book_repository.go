package cache

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
)

type CacheBookRepository struct {
	c         *cache.Cache
	cacheName string
}

func NewCacheBookRepository() domain.BookRepository {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &CacheBookRepository{
		c:         c,
		cacheName: "book",
	}
}

func (o *CacheBookRepository) fetchBook(url string) []domain.Book {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	type Author struct {
		ID   string `json:"id"`
		Key  string `json:"key"`
		Name string `json:"name"`
	}

	type Work struct {
		ID           string   `json:"id"`
		Key          string   `json:"key"`
		Title        string   `json:"title"`
		EditionCount int      `json:"edition_count"`
		CoverId      int      `json:"cover_id"`
		Subjects     []string `json:"subjects"`
		Subject      []string `json:"subject"`
		Authors      []Author `json:"authors"`
	}

	type SubjectResponse struct {
		Key         string `json:"key"`
		Name        string `json:"name"`
		SubjectType string `json:"subject_type"`
		WorkCount   int    `json:"work_count"`
		Works       []Work `json:"works"`
	}

	var subject SubjectResponse
	err = json.Unmarshal(responseData, &subject)
	if err != nil {
		log.Fatal(err)
	}

	books := []domain.Book{}
	for i, work := range subject.Works {
		subject.Works[i].ID = work.Key
		subject.Works[i].Subjects = work.Subject
		for i2, author := range subject.Works[i].Authors {
			subject.Works[i].Authors[i2].ID = author.Key
		}
	}
	byteData, _ := json.Marshal(subject.Works)
	json.Unmarshal(byteData, &books)

	return books
}

func (o *CacheBookRepository) Index(ctx context.Context) ([]domain.Book, error) {
	data, found := o.c.Get(o.cacheName)
	books := []domain.Book{}
	if found {
		byteData, _ := json.Marshal(data)
		json.Unmarshal(byteData, &books)
	} else {
		books = o.fetchBook("http://openlibrary.org/subjects/love.json?published_in=1500-1600")
		o.c.Set(o.cacheName, books, cache.DefaultExpiration)
	}

	if len(books) < 1 {
		return []domain.Book{}, domain.ErrNotFound
	}

	return books, nil
}
