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
)

func TestBooking(t *testing.T) {

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)
	the_day_after_tomorrow := now.AddDate(0, 0, 2)

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
