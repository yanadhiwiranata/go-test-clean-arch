package domain

import (
	"context"
	"time"
)

type Booking struct {
	ID       int       `json:"id"`
	BookID   string    `json:"book_id"`
	Quantity int       `json:"quantity"`
	BookAt   time.Time `json:"book_at"`
	ReturnAt time.Time `json:"return_at"`
}

type BookingRepository interface {
	CountCurrentBooking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time) (int, error)
	Booking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time, quantity int) (Booking, error)
	FilterBooking(ctx context.Context, bookAt time.Time, returnAt time.Time) []Booking
}

type BookingUsecase interface {
	Booking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time, quantity int) (Booking, error)
	Index(ctx context.Context, bookAt time.Time, returnAt time.Time) ([]Booking, error)
}
