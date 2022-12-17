package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	"github.com/yanadhiwiranata/go-test-clean-arch/util"
)

type CacheBookRepository struct {
	c         *cache.Cache
	cacheName string
}

func NewCacheBookingRepository() domain.BookingRepository {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &CacheBookRepository{c: c, cacheName: "booking"}
}

func (s *CacheBookRepository) CountCurrentBooking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time) (int, error) {
	bookings := s.AllBooking()
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

func (s *CacheBookRepository) Booking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time, quantity int) (domain.Booking, error) {
	if quantity < 1 {
		return domain.Booking{}, domain.ErrBadParamInput
	}

	// TODO lock or add sharing sequence to keep ID increment correctly when scale up(it will be ok if using database sequence)
	bookings := s.AllBooking()
	newBooking := domain.Booking{
		ID:       len(bookings) + 1,
		BookID:   bookID,
		BookAt:   bookAt,
		ReturnAt: returnAt,
		Quantity: quantity,
	}
	bookings = append(bookings, newBooking)
	return newBooking, nil
}

func (s *CacheBookRepository) AllBooking() []domain.Booking {
	data, found := s.c.Get(s.cacheName)
	bookings := []domain.Booking{}
	if found {
		byteData, _ := json.Marshal(data)
		json.Unmarshal(byteData, &bookings)
	}
	return bookings
}

func (s *CacheBookRepository) FilterBooking(ctx context.Context, bookAt time.Time, returnAt time.Time) []domain.Booking {
	all_bookings := s.AllBooking()
	bookings := []domain.Booking{}

	for _, book := range all_bookings {
		if util.SameDay(bookAt, book.BookAt) || util.SameDay(returnAt, book.ReturnAt) {

		} else if bookAt.After(book.BookAt) && bookAt.After(book.ReturnAt) {
			continue
		} else if returnAt.Before(book.BookAt) && returnAt.Before(book.ReturnAt) {
			continue
		}
		bookings = append(bookings, book)
	}

	return bookings
}
