package echo_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	_bookHandler "github.com/yanadhiwiranata/go-test-clean-arch/book/delivery/http/echo"
	_bookUsecase "github.com/yanadhiwiranata/go-test-clean-arch/book/usecase"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	mocks "github.com/yanadhiwiranata/go-test-clean-arch/mocks/domain"
)

func TestResponse(t *testing.T) {
	e := echo.New()
	path := "/books/test_response"
	req, err := http.NewRequest(echo.GET, path, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(path)

	handler := &_bookHandler.BookHandler{
		BookUsecase: nil,
	}

	err = handler.TestResponse(c)

	res := rec.Result()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, body, []byte("book index"))
	assert.NoError(t, nil)
}

func TestIndex(t *testing.T) {
	e := echo.New()
	subject := "Bibliography"
	path := "/books"
	req, err := http.NewRequest(echo.GET, path, strings.NewReader(""))
	assert.NoError(t, err)

	q := req.URL.Query()
	q.Add("subject", subject)
	req.URL.RawQuery = q.Encode()

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(path)
	c.SetParamNames("subject")
	c.SetParamValues(subject)

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

	mockBookRepo := new(mocks.BookRepository)

	mockBookRepo.On("FilterBySubject", mock.Anything, subject).Return(mockBooks, nil)
	mockBookUsecase := &_bookUsecase.BookUsecase{BookRepository: mockBookRepo}

	handler := _bookHandler.BookHandler{
		BookUsecase: mockBookUsecase,
	}

	err = handler.Index(c)

	res := rec.Result()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, body)
	assert.NoError(t, nil)
}
