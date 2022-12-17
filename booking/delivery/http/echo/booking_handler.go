package echo

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	"github.com/yanadhiwiranata/go-test-clean-arch/util"
)

type BookingHandler struct {
	BookingUsecase domain.BookingUsecase
}

func NewBookingHandler(e *echo.Echo, bookingUsecase domain.BookingUsecase) {
	handler := &BookingHandler{
		BookingUsecase: bookingUsecase,
	}
	bookGroup := e.Group("/bookings")
	bookGroup.GET("", handler.Index)
	bookGroup.POST("", handler.Booking)
}

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type BookingPostRequest struct {
	ID       int                      `json:"id"`
	BookID   string                   `json:"book_id"`
	Quantity int                      `json:"quantity"`
	BookAt   domain.CustomRequestTime `json:"book_at"`
	ReturnAt domain.CustomRequestTime `json:"return_at"`
}

type BookingResponse struct {
	ID       int                      `json:"id"`
	BookID   string                   `json:"book_id"`
	Quantity int                      `json:"quantity"`
	BookAt   domain.CustomResposeTime `json:"book_at"`
	ReturnAt domain.CustomResposeTime `json:"return_at"`
}

func (s *BookingHandler) Index(c echo.Context) error {
	startAt, err1 := util.GetTimestampFromUnixString(c.QueryParam("startAt"))
	endAt, err2 := util.GetTimestampFromUnixString(c.QueryParam("endAt"))

	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: domain.ErrBadParamInput.Error()})
	}

	bookings, err := s.BookingUsecase.Index(c.Request().Context(), startAt, endAt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	var response []BookingResponse
	byteData, _ := json.Marshal(bookings)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *BookingHandler) Booking(c echo.Context) error {
	var request BookingPostRequest
	err := c.Bind(&request)
	if err != nil {
		return c.String(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	bookings, err := s.BookingUsecase.Booking(c.Request().Context(), request.BookID, request.BookAt.Time, request.ReturnAt.Time, request.Quantity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	var response BookingResponse
	byteData, _ := json.Marshal(bookings)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, response)
}
