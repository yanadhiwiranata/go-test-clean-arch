package usecase

import (
	"context"
	"time"

	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	"github.com/yanadhiwiranata/go-test-clean-arch/util"
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
		if !util.SameDay(bookAt, now) {
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

func (s *BookingUsecase) Index(ctx context.Context, startAt time.Time, endAt time.Time) ([]domain.Booking, error) {
	if startAt.After(endAt) {
		return []domain.Booking{}, domain.ErrBadParamInput
	}

	bookings := s.bookingReposity.FilterBooking(ctx, startAt, endAt)

	return bookings, nil
}
