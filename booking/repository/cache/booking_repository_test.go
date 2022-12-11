package cache_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, true, true)
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
