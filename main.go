package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_book_handler_echo "github.com/yanadhiwiranata/go-test-clean-arch/book/delivery/http/echo"
	_book_cache_repository "github.com/yanadhiwiranata/go-test-clean-arch/book/repository/cache"
	_book_usecase "github.com/yanadhiwiranata/go-test-clean-arch/book/usecase"
	_booking_handler_echo "github.com/yanadhiwiranata/go-test-clean-arch/booking/delivery/http/echo"
	_booking_cache_repository "github.com/yanadhiwiranata/go-test-clean-arch/booking/repository/cache"
	_booking_usecase "github.com/yanadhiwiranata/go-test-clean-arch/booking/usecase"
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
	bRepo := _book_cache_repository.NewCacheBookRepository()
	bUsecase := _book_usecase.NewBookUsecase(bRepo)
	_book_handler_echo.NewBookHandler(r, bUsecase)

	bookingRepo := _booking_cache_repository.NewCacheBookingRepository()
	bookingUsecase := _booking_usecase.NewBookingUsecase(bookingRepo, bRepo)
	_booking_handler_echo.NewBookingHandler(r, bookingUsecase)
	return r
}
