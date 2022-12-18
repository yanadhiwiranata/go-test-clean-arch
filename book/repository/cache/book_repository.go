package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (o *CacheBookRepository) fetchBook(ctx context.Context, url string) ([]domain.Book, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return []domain.Book{}, err
	}

	ctx1, cancel := context.WithTimeout(ctx, time.Millisecond*300000)
	defer cancel()

	request = request.WithContext(ctx1)
	transport := new(http.Transport)
	client := new(http.Client)
	client.Transport = transport

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return []domain.Book{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return []domain.Book{}, err
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
		fmt.Println(err)
		return []domain.Book{}, err
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

	return books, nil
}

func (o *CacheBookRepository) Index(ctx context.Context) ([]domain.Book, error) {
	data, found := o.c.Get(o.cacheName)
	books := []domain.Book{}
	var err error
	if found {
		byteData, _ := json.Marshal(data)
		json.Unmarshal(byteData, &books)
	} else {
		books, err = o.fetchBook(ctx, "http://openlibrary.org/subjects/love.json?published_in=1500-1600")
		if err != nil {
			return books, err
		}
		o.c.Set(o.cacheName, books, cache.DefaultExpiration)
	}

	return books, nil
}

func (o *CacheBookRepository) FilterBySubject(ctx context.Context, subject string) ([]domain.Book, error) {
	books, err := o.Index(ctx)
	if err != nil {
		return []domain.Book{}, err
	}

	var filterredBook []domain.Book
	for _, book := range books {
		for _, s := range book.Subjects {
			if s == subject {
				filterredBook = append(filterredBook, book)
				break
			}
		}
	}

	return filterredBook, nil
}

func (o *CacheBookRepository) FilterByID(ctx context.Context, id string) (domain.Book, error) {
	books, err := o.Index(ctx)

	if err != nil {
		return domain.Book{}, err
	}

	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}

	return domain.Book{}, nil
}
