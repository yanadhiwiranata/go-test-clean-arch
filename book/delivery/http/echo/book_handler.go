package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
}

func NewBookHandler(e *echo.Echo) {
	handler := &BookHandler{}
	bookGroup := e.Group("/books")
	bookGroup.GET("/test_response", handler.TestResponse)
}

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (a *BookHandler) TestResponse(c echo.Context) error {
	return c.String(http.StatusOK, "book index")
}
