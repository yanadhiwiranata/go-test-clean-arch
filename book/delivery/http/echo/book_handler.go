package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	_domain_helper "github.com/yanadhiwiranata/go-test-clean-arch/domain/helper"
)

type BookHandler struct {
	BookUsecase domain.BookUsecase
}

func NewBookHandler(e *echo.Echo, bookUsecase domain.BookUsecase) {
	handler := &BookHandler{
		BookUsecase: bookUsecase,
	}
	bookGroup := e.Group("/books")
	bookGroup.GET("", handler.Index)
	bookGroup.GET("/test_response", handler.TestResponse)
}

func (s *BookHandler) TestResponse(c echo.Context) error {
	return c.String(http.StatusOK, "book index")
}

func (s *BookHandler) Index(c echo.Context) error {
	subject := c.QueryParam("subject")
	books, err := s.BookUsecase.Index(c.Request().Context(), subject)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _domain_helper.ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}
