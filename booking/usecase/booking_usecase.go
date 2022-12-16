package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
)

type BookingUsecase struct {
	bookingReposity domain.BookingRepository
	bookRepository  domain.BookRepository
}

func NewBookingUsecase(bookingReposity domain.BookingRepository, bookRepository domain.BookRepository) *BookingUsecase {
	return &BookingUsecase{
		bookingReposity: bookingReposity,
		bookRepository:  bookRepository,
	}
}

func (s *BookingUsecase) Booking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time, quantity int) (domain.Booking, error) {
	now := time.Now()
	if bookAt.Before(now) {
		fmt.Println("date yesterday")
		if !s.sameDay(bookAt, now) {
			fmt.Println("date not valid")
			return domain.Booking{}, domain.ErrBadParamInput
		}
	}

	if bookAt.After(returnAt) {
		return domain.Booking{}, domain.ErrBadParamInput
	}

	book, err := s.bookRepository.FilterByID(ctx, bookID)
	if err != nil {
		return domain.Booking{}, err
	}

	if book.ID == "" {
		return domain.Booking{}, domain.ErrNotFound
	}

	currentBookingQuantity, err := s.bookingReposity.CountCurrentBooking(ctx, book.ID, bookAt, returnAt)
	if err != nil {
		return domain.Booking{}, err
	}

	availableQuantity := book.EditionCount - currentBookingQuantity

	if (availableQuantity - quantity) < 0 {
		return domain.Booking{}, domain.ErrBadParamInput
	}

	booking, err := s.bookingReposity.Booking(ctx, bookID, bookAt, returnAt, quantity)
	if err != nil {
		return domain.Booking{}, err
	}

	if booking.ID <= 0 {
		return domain.Booking{}, domain.ErrInternalServerError
	}

	return booking, nil
}

// common method will be move to lib later
func (s *BookingUsecase) sameDay(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
