package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
)

type BookHandler struct {
	BookUsecase domain.BookUsecase
}

func NewBookHandler(e *echo.Echo, bookUsecase domain.BookUsecase) {
	handler := &BookHandler{
		BookUsecase: bookUsecase,
	}
	bookGroup := e.Group("/books")
	bookGroup.GET("", handler.TestResponse)
	bookGroup.GET("/test_response", handler.TestResponse)
}

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (s *BookHandler) TestResponse(c echo.Context) error {
	return c.String(http.StatusOK, "book index")
}

func (s *BookHandler) Index(c echo.Context) error {
	subject := c.Param("subject")
	books, err := s.BookUsecase.Index(c.Request().Context(), subject)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}
