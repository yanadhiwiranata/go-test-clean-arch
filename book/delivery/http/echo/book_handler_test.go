package echo_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_bookHandler "github.com/yanadhiwiranata/go-test-clean-arch/book/delivery/http/echo"
)

func TestAdminIndex(t *testing.T) {
	e := echo.New()
	path := "/books/test_response"
	req, err := http.NewRequest(echo.GET, path, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(path)
	handler := _bookHandler.BookHandler{}

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
