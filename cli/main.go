package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	_book_handler_echo "github.com/yanadhiwiranata/go-test-clean-arch/book/delivery/http/echo"
	_book_cache_repository "github.com/yanadhiwiranata/go-test-clean-arch/book/repository/cache"
	_book_usecase "github.com/yanadhiwiranata/go-test-clean-arch/book/usecase"
	_booking_handler_echo "github.com/yanadhiwiranata/go-test-clean-arch/booking/delivery/http/echo"
	_booking_cache_repository "github.com/yanadhiwiranata/go-test-clean-arch/booking/repository/cache"
	_booking_usecase "github.com/yanadhiwiranata/go-test-clean-arch/booking/usecase"
	_domain_helper "github.com/yanadhiwiranata/go-test-clean-arch/domain/helper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	r := defaultEchoServer()
	err := http.ListenAndServe(viper.GetString("server.address"), r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("done")
}

func defaultEchoServer() *echo.Echo {
	r := echo.New()
	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	bookRepo := _book_cache_repository.NewCacheBookRepository()
	bookUsecase := _book_usecase.NewBookUsecase(bookRepo, timeoutContext)
	_book_handler_echo.NewBookHandler(r, bookUsecase)

	bookingRepo := _booking_cache_repository.NewCacheBookingRepository(bookRepo)
	bookingUsecase := _booking_usecase.NewBookingUsecase(bookingRepo, bookRepo, timeoutContext)
	_booking_handler_echo.NewBookingHandler(r, bookingUsecase)

	r.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	r.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			response := _domain_helper.ResponseError{
				Message: "Bad Gateway",
			}
			c.JSON(http.StatusBadGateway, response)
		},
		Timeout: 5000 * time.Millisecond,
	}))

	r.HTTPErrorHandler = customHTTPErrorHandler
	return r
}

func customHTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Error(report)

	response := _domain_helper.ResponseError{
		Message: "Internal Server Error",
	}
	c.JSON(http.StatusInternalServerError, response)
}
