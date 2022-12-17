package echo_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	_booking_handler_echo "github.com/yanadhiwiranata/go-test-clean-arch/booking/delivery/http/echo"
	_booking_usecase "github.com/yanadhiwiranata/go-test-clean-arch/booking/usecase"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	mocks "github.com/yanadhiwiranata/go-test-clean-arch/mocks/domain"
)

func TestBooking(t *testing.T) {

	assert.Equal(t, true, true)
}

func TestIndex(t *testing.T) {
	e := echo.New()
	now := time.Now().Truncate(time.Second)
	startAt := now
	startAtString := strconv.Itoa(int(startAt.Unix()))
	endAt := now
	endAtString := strconv.Itoa(int(endAt.Unix()))
	path := "/books"
	req, err := http.NewRequest(echo.GET, path, strings.NewReader(""))
	assert.NoError(t, err)

	q := req.URL.Query()
	q.Add("startAt", startAtString)
	q.Add("endAt", endAtString)
	req.URL.RawQuery = q.Encode()

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(path)
	c.SetParamNames("startAt")
	c.SetParamValues(startAtString)
	c.SetParamNames("endAt")
	c.SetParamValues(endAtString)

	mockAuthor1 := domain.Author{ID: "/authors/OL9388A", Name: "Yan"}
	mockBooks := []domain.Book{
		{
			ID:           "/works/OL362427W",
			Title:        "title 1",
			EditionCount: 20,
			Authors:      []domain.Author{mockAuthor1},
			Subjects:     []string{"asd", "sad"},
		},
	}

	mockBookings := []domain.Booking{
		{
			ID:       1,
			BookID:   mockBooks[0].ID,
			Quantity: 10,
			BookAt:   now,
			ReturnAt: now,
		},
	}

	mockBookRepo := new(mocks.BookRepository)

	type testCase struct {
		name     string
		startAt  time.Time
		endAt    time.Time
		bookings []domain.Booking
		err      error
		status   int
	}

	tcs := []testCase{
		{name: "Show bookings", startAt: now, endAt: now, bookings: mockBookings, err: nil, status: http.StatusOK},
		{name: "Show empty bookings", startAt: now, endAt: now, bookings: []domain.Booking{}, err: nil, status: http.StatusOK},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockBookingRepo := new(mocks.BookingRepository)
			mockBookingRepo.On("FilterBooking", mock.Anything, tc.startAt, tc.endAt).Return(tc.bookings)
			bookingUsecase := _booking_usecase.NewBookingUsecase(mockBookingRepo, mockBookRepo)
			handler := _booking_handler_echo.BookingHandler{
				BookingUsecase: bookingUsecase,
			}
			err = handler.Index(c)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.status, rec.Code)

			res := rec.Result()
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err)

			if len(tc.bookings) >= 0 {
				var response []_booking_handler_echo.BookingPostRequest
				err = json.Unmarshal(body, &response)
				assert.Equal(t, len(tc.bookings), len(response))
			}
		})
	}

	assert.Equal(t, true, true)
}
