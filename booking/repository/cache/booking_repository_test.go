package cache_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	_cacheBookRepository "github.com/yanadhiwiranata/go-test-clean-arch/booking/repository/cache"
	"github.com/yanadhiwiranata/go-test-clean-arch/domain"
)

func TestCountCurrentBooking(t *testing.T) {
	type testCase struct {
		name     string
		bookAt   time.Time
		returnAt time.Time
		quantity int
	}

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)
	the_day_after_tomorrow := now.AddDate(0, 0, 2)

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

	bookingRepository := _cacheBookRepository.NewCacheBookingRepository()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			patches := gomonkey.ApplyMethod(reflect.TypeOf(bookingRepository), "AllBooking", func(m *_cacheBookRepository.CacheBookRepository) []domain.Booking {
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
		name string
	}

	tcs := []testCase{
		{name: "booking success"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, true, true)
		})

	}
}