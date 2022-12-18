package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	"github.com/yanadhiwiranata/go-test-clean-arch/util"
)

type CacheBookingRepository struct {
	c              *cache.Cache
	bookRepository domain.BookRepository
	cacheName      string
}

func NewCacheBookingRepository(bookRepository domain.BookRepository) domain.BookingRepository {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &CacheBookingRepository{c: c, cacheName: "booking", bookRepository: bookRepository}
}

func (s *CacheBookingRepository) CountCurrentBooking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time) (int, error) {
	bookings := s.AllBooking(ctx)
	count := 0

	for _, book := range bookings {
		if book.BookID == bookID {
			if bookAt.After(book.BookAt) && bookAt.After(book.ReturnAt) {
				continue
			} else if returnAt.Before(book.BookAt) && returnAt.Before(book.ReturnAt) {
				continue
			}
			count += book.Quantity
		}

	}
	return count, nil
}

func (s *CacheBookingRepository) Booking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time, quantity int) (domain.Booking, error) {
	if quantity < 1 {
		return domain.Booking{}, domain.ErrBadParamInput
	}

	// TODO lock or add sharing sequence to keep ID increment correctly when scale up(it will be ok if using database sequence)
	bookings := s.AllBooking(ctx)
	newBooking := domain.Booking{
		ID:       len(bookings) + 1,
		BookID:   bookID,
		BookAt:   bookAt,
		ReturnAt: returnAt,
		Quantity: quantity,
	}
	bookings = append(bookings, newBooking)
	s.c.Set(s.cacheName, bookings, cache.DefaultExpiration)
	return newBooking, nil
}

func (s *CacheBookingRepository) AllBooking(ctx context.Context) []domain.Booking {
	data, found := s.c.Get(s.cacheName)
	bookings := []domain.Booking{}
	if found {
		byteData, _ := json.Marshal(data)
		json.Unmarshal(byteData, &bookings)
	}

	if len(bookings) > 0 {
		for i, booking := range bookings {
			book, err := s.bookRepository.FilterByID(ctx, booking.BookID)
			if err != nil {
				bookings[i].Book = book
			}
		}
	}

	return bookings
}

func (s *CacheBookingRepository) FilterBooking(ctx context.Context, startAt time.Time, endAt time.Time) []domain.Booking {
	all_bookings := s.AllBooking(ctx)
	bookings := []domain.Booking{}

	for _, booking := range all_bookings {
		if util.SameDay(startAt, booking.BookAt) || util.SameDay(endAt, booking.ReturnAt) {

		} else if startAt.After(booking.BookAt) && startAt.After(booking.ReturnAt) {
			continue
		} else if endAt.Before(booking.BookAt) && endAt.Before(booking.ReturnAt) {
			continue
		}
		bookings = append(bookings, booking)
	}

	return bookings
}
