package cache_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	_cacheBookingRepository "github.com/yanadhiwiranata/go-test-clean-arch/booking/repository/cache"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
	mocks "github.com/yanadhiwiranata/go-test-clean-arch/mocks/domain"
	"github.com/yanadhiwiranata/go-test-clean-arch/util"
)

func TestCountCurrentBooking(t *testing.T) {
	type testCase struct {
		name     string
		bookAt   time.Time
		returnAt time.Time
		quantity int
	}

	yesterday, now, tomorrow, the_day_after_tomorrow := util.GenerateSampleTestTime()

	fullQuantity := 20
	halfQuantity := fullQuantity / 2

	tcs := []testCase{
		{name: "invalid for yesterday request", bookAt: yesterday, returnAt: yesterday, quantity: 0},
		{name: "today empty booking", bookAt: now, returnAt: now, quantity: 0},
		{name: "tomorrow half booking", bookAt: tomorrow, returnAt: tomorrow, quantity: halfQuantity},
		{name: "the day after tomorrow full booking", bookAt: the_day_after_tomorrow, returnAt: the_day_after_tomorrow, quantity: fullQuantity},
	}

	mockBooks := []domain.Book{
		{
			ID:           "works/OL8193420W",
			Title:        "title 1",
			EditionCount: 20,
			Authors:      []domain.Author{},
		},
	}

	mockBookings := []domain.Booking{
		{
			ID:       1,
			BookID:   mockBooks[0].ID,
			Quantity: halfQuantity,
			BookAt:   tomorrow,
			ReturnAt: the_day_after_tomorrow,
		},
		{
			ID:       2,
			BookID:   mockBooks[0].ID,
			Quantity: halfQuantity,
			BookAt:   the_day_after_tomorrow,
			ReturnAt: the_day_after_tomorrow,
		},
	}

	mockBookRepository := new(mocks.BookRepository)
	bookingRepository := _cacheBookingRepository.NewCacheBookingRepository(mockBookRepository)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			patches := gomonkey.ApplyMethod(reflect.TypeOf(bookingRepository), "AllBooking", func(m *_cacheBookingRepository.CacheBookingRepository, ctx context.Context) []domain.Booking {
				return mockBookings
			})
			defer patches.Reset()

			count, err := bookingRepository.CountCurrentBooking(context.Background(), mockBooks[0].ID, tc.bookAt, tc.returnAt)
			assert.NoError(t, err)
			assert.Equal(t, tc.quantity, count)
		})
	}
}

func TestBooking(t *testing.T) {
	type testCase struct {
		name     string
		bookID   string
		bookAt   time.Time
		returnAt time.Time
		quantity int
		success  bool
	}

	mockBook := domain.Book{
		ID:           "works/OL8193420W",
		Title:        "title 1",
		EditionCount: 20,
		Authors:      []domain.Author{},
	}

	_, now, tomorrow, the_day_after_tomorrow := util.GenerateSampleTestTime()

	halfQuantity := mockBook.EditionCount / 2

	mockBookings := []domain.Booking{
		{
			ID:       1,
			BookID:   mockBook.ID,
			Quantity: halfQuantity,
			BookAt:   tomorrow,
			ReturnAt: the_day_after_tomorrow,
		},
	}
	tcs := []testCase{
		{name: "booking 1 quantity book", bookAt: now, returnAt: now, quantity: 1, success: true},
		{name: "booking 0 quantity book", bookID: mockBook.ID, bookAt: now, returnAt: now, quantity: 0, success: false},
		{name: "booking -1 quantity book", bookID: mockBook.ID, bookAt: now, returnAt: now, quantity: -1, success: false},
	}

	mockBookRepository := new(mocks.BookRepository)
	bookingRepository := _cacheBookingRepository.NewCacheBookingRepository(mockBookRepository)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			patches := gomonkey.ApplyMethod(reflect.TypeOf(bookingRepository), "AllBooking", func(m *_cacheBookingRepository.CacheBookingRepository, ctx context.Context) []domain.Booking {
				return mockBookings
			})
			defer patches.Reset()

			booking, err := bookingRepository.Booking(context.Background(), mockBook.ID, tc.bookAt, tc.returnAt, tc.quantity)
			if tc.success {
				assert.NoError(t, err)
				assert.Equal(t, len(mockBookings)+1, booking.ID)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestFilterBookingByTime(t *testing.T) {
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

	type testCase struct {
		name            string
		bookAt          time.Time
		returnAt        time.Time
		bookingQuantity int
	}

	tcs := []testCase{
		{name: "show today book", bookAt: now, returnAt: now, bookingQuantity: 2},
		{name: "show tomorrow book", bookAt: tomorrow, returnAt: tomorrow, bookingQuantity: 3},
		{name: "show the day after tomorrow book", bookAt: the_day_after_tomorrow, returnAt: the_day_after_tomorrow, bookingQuantity: 4},
	}

	mockBookRepository := new(mocks.BookRepository)
	bookingRepository := _cacheBookingRepository.NewCacheBookingRepository(mockBookRepository)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			patches := gomonkey.ApplyMethod(reflect.TypeOf(bookingRepository), "AllBooking", func(m *_cacheBookingRepository.CacheBookingRepository, ctx context.Context) []domain.Booking {
				return mockBookings
			})
			defer patches.Reset()

			bookings := bookingRepository.FilterBooking(context.Background(), tc.bookAt, tc.returnAt)
			assert.Equal(t, tc.bookingQuantity, len(bookings))
		})
	}
}
