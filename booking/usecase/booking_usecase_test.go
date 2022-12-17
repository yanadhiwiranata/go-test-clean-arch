package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	_bookingUsecase "github.com/yanadhiwiranata/go-test-clean-arch/booking/usecase"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	mocks "github.com/yanadhiwiranata/go-test-clean-arch/mocks/domain"
	"github.com/yanadhiwiranata/go-test-clean-arch/util"
)

func TestBooking(t *testing.T) {

	yesterday, now, tomorrow, the_day_after_tomorrow := util.GenerateSampleTestTime()

	mockBook := domain.Book{
		ID:           "works/OL8193420W",
		Title:        "title 1",
		EditionCount: 20,
		Authors:      []domain.Author{},
	}

	mockBooking := domain.Booking{
		ID:       1,
		BookID:   mockBook.ID,
		Quantity: mockBook.EditionCount / 2,
		BookAt:   tomorrow,
		ReturnAt: the_day_after_tomorrow,
	}

	type testCase struct {
		name                 string
		bookID               string
		bookEditionCount     int
		requestBookAt        time.Time
		requestReturnAt      time.Time
		requestQuantity      int
		responseFilterBook   domain.Book
		responseFilterError  error
		responseQuantity     int
		responseBooking      domain.Booking
		responseBookingError error
		responseError        error
	}

	tcs := []testCase{
		{name: "Booking success", bookID: mockBook.ID, bookEditionCount: 10, requestBookAt: now, requestReturnAt: tomorrow, requestQuantity: 10, responseQuantity: 0, responseFilterBook: mockBook, responseFilterError: nil, responseBooking: mockBooking, responseBookingError: nil, responseError: nil},
		{name: "Failed booking because no book found", bookID: "asdasda", bookEditionCount: 10, requestBookAt: now, requestReturnAt: now, requestQuantity: 10, responseQuantity: 0, responseFilterBook: domain.Book{}, responseFilterError: domain.ErrNotFound, responseBooking: domain.Booking{}, responseBookingError: nil, responseError: domain.ErrNotFound},
		{name: "Failed booking because no stock", bookID: mockBook.ID, bookEditionCount: 10, requestBookAt: now, requestReturnAt: now, requestQuantity: 10, responseQuantity: 10, responseFilterBook: mockBook, responseFilterError: nil, responseBooking: domain.Booking{}, responseBookingError: nil, responseError: domain.ErrNotFound},
		{name: "Failed booking because request quantity > available stock", bookID: mockBook.ID, bookEditionCount: 10, requestBookAt: now, requestReturnAt: now, requestQuantity: 20, responseQuantity: 0, responseFilterBook: mockBook, responseFilterError: nil, responseBooking: domain.Booking{}, responseBookingError: nil, responseError: domain.ErrBadParamInput},
		{name: "Failed booking because bookAt is yesterday", bookID: mockBook.ID, bookEditionCount: 10, requestBookAt: yesterday, requestReturnAt: tomorrow, requestQuantity: 10, responseQuantity: 0, responseFilterBook: mockBook, responseFilterError: nil, responseBooking: mockBooking, responseBookingError: nil, responseError: domain.ErrBadParamInput},
		{name: "Failed booking because bookAt is later than returnAt", bookID: mockBook.ID, bookEditionCount: 10, requestBookAt: tomorrow, requestReturnAt: now, requestQuantity: 10, responseQuantity: 0, responseFilterBook: mockBook, responseFilterError: nil, responseBooking: mockBooking, responseBookingError: nil, responseError: domain.ErrBadParamInput},
	}

	mockBookingRepo := new(mocks.BookingRepository)
	mockBookRepo := new(mocks.BookRepository)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if tc.responseFilterError == nil {
				tc.responseFilterBook.EditionCount = tc.bookEditionCount
			}

			if tc.responseBookingError == nil {
				tc.responseBooking.Quantity = tc.requestQuantity
			}
			mockBooking.Quantity = tc.requestQuantity
			mockBooking.BookAt = tc.requestBookAt
			mockBooking.ReturnAt = tc.requestReturnAt
			mockBookRepo.On("FilterByID", mock.Anything, tc.bookID).Return(tc.responseFilterBook, tc.responseFilterError)
			mockBookingRepo.On("CountCurrentBooking", mock.Anything, tc.bookID, tc.requestBookAt, tc.requestReturnAt).Return(tc.responseQuantity, nil)
			mockBookingRepo.On("Booking", mock.Anything, tc.bookID, tc.requestBookAt, tc.requestReturnAt, tc.requestQuantity).Return(tc.responseBooking, tc.responseBookingError)

			bookingUsecase := _bookingUsecase.NewBookingUsecase(mockBookingRepo, mockBookRepo)
			_, err := bookingUsecase.Booking(context.Background(), tc.bookID, tc.requestBookAt, tc.requestReturnAt, tc.requestQuantity)
			if tc.responseError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	type testCase struct {
		name            string
		startAt         time.Time
		endAt           time.Time
		bookings        []domain.Booking
		bookingQuantity int
	}

	_, now, tomorrow, the_day_after_tomorrow := util.GenerateSampleTestTime()

	mockBook := domain.Book{
		ID:           "works/OL8193420W",
		Title:        "title 1",
		EditionCount: 20,
		Authors:      []domain.Author{},
	}

	allDayBooking := []domain.Booking{
		{
			ID:       10,
			BookID:   mockBook.ID,
			Quantity: 1,
			BookAt:   now,
			ReturnAt: the_day_after_tomorrow,
		},
	}

	todayBooking := []domain.Booking{
		{
			ID:       1,
			BookID:   mockBook.ID,
			Quantity: 1,
			BookAt:   now,
			ReturnAt: now,
		},
	}

	tomorrowBooking := []domain.Booking{
		{
			ID:       2,
			BookID:   mockBook.ID,
			Quantity: 1,
			BookAt:   tomorrow,
			ReturnAt: tomorrow,
		},
		{
			ID:       3,
			BookID:   mockBook.ID,
			Quantity: 1,
			BookAt:   tomorrow,
			ReturnAt: tomorrow,
		},
	}

	theDayAfterTomorrowBooking := []domain.Booking{
		{
			ID:       4,
			BookID:   mockBook.ID,
			Quantity: 1,
			BookAt:   the_day_after_tomorrow,
			ReturnAt: the_day_after_tomorrow,
		},
		{
			ID:       5,
			BookID:   mockBook.ID,
			Quantity: 1,
			BookAt:   the_day_after_tomorrow,
			ReturnAt: the_day_after_tomorrow,
		},
		{
			ID:       6,
			BookID:   mockBook.ID,
			Quantity: 1,
			BookAt:   the_day_after_tomorrow,
			ReturnAt: the_day_after_tomorrow,
		},
	}

	mockBookings := []domain.Booking{}
	mockBookings = append(mockBookings, allDayBooking...)
	mockBookings = append(mockBookings, todayBooking...)
	mockBookings = append(mockBookings, tomorrowBooking...)
	mockBookings = append(mockBookings, theDayAfterTomorrowBooking...)

	tcs := []testCase{
		{name: "show today book", startAt: now, endAt: now, bookings: append(todayBooking, allDayBooking...), bookingQuantity: len(append(todayBooking, allDayBooking...))},
		{name: "show tomorrow book", startAt: tomorrow, endAt: tomorrow, bookings: append(tomorrowBooking, allDayBooking...), bookingQuantity: len(append(tomorrowBooking, allDayBooking...))},
		{name: "show the day after tomorrow book", startAt: the_day_after_tomorrow, endAt: the_day_after_tomorrow, bookings: append(theDayAfterTomorrowBooking, allDayBooking...), bookingQuantity: len(append(theDayAfterTomorrowBooking, allDayBooking...))},
	}

	mockBookingRepo := new(mocks.BookingRepository)
	mockBookRepo := new(mocks.BookRepository)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockBookingRepo.On("FilterBooking", mock.Anything, tc.startAt, tc.endAt).Return(tc.bookings)
			bookingUsecase := _bookingUsecase.NewBookingUsecase(mockBookingRepo, mockBookRepo)
			bookings, _ := bookingUsecase.Index(context.Background(), tc.startAt, tc.endAt)
			assert.Equal(t, tc.bookingQuantity, len(bookings))
		})
	}

}
