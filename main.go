package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_book_handler_echo "github.com/yanadhiwiranata/go-test-clean-arch/book/delivery/http/echo"
)

func main() {
	r := defaultEchoServer()
	http.ListenAndServe(":4000", r)
}

func defaultEchoServer() *echo.Echo {
	r := echo.New()
	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	_book_handler_echo.NewBookHandler(r)
	return r
}
