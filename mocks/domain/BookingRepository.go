// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/yanadhiwiranata/go-test-clean-arch/domain"

	time "time"
)

// BookingRepository is an autogenerated mock type for the BookingRepository type
type BookingRepository struct {
	mock.Mock
}

// Booking provides a mock function with given fields: ctx, bookID, bookAt, returnAt, quantity
func (_m *BookingRepository) Booking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time, quantity int) (domain.Booking, error) {
	ret := _m.Called(ctx, bookID, bookAt, returnAt, quantity)

	var r0 domain.Booking
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, time.Time, int) domain.Booking); ok {
		r0 = rf(ctx, bookID, bookAt, returnAt, quantity)
	} else {
		r0 = ret.Get(0).(domain.Booking)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time, time.Time, int) error); ok {
		r1 = rf(ctx, bookID, bookAt, returnAt, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountCurrentBooking provides a mock function with given fields: ctx, bookID, bookAt, returnAt
func (_m *BookingRepository) CountCurrentBooking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time) (int, error) {
	ret := _m.Called(ctx, bookID, bookAt, returnAt)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, time.Time) int); ok {
		r0 = rf(ctx, bookID, bookAt, returnAt)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time, time.Time) error); ok {
		r1 = rf(ctx, bookID, bookAt, returnAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterBooking provides a mock function with given fields: ctx, bookAt, returnAt
func (_m *BookingRepository) FilterBooking(ctx context.Context, bookAt time.Time, returnAt time.Time) []domain.Booking {
	ret := _m.Called(ctx, bookAt, returnAt)

	var r0 []domain.Booking
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []domain.Booking); ok {
		r0 = rf(ctx, bookAt, returnAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Booking)
		}
	}

	return r0
}

type mockConstructorTestingTNewBookingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookingRepository creates a new instance of BookingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookingRepository(t mockConstructorTestingTNewBookingRepository) *BookingRepository {
	mock := &BookingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
