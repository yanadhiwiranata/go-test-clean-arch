// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/yanadhiwiranata/go-test-clean-arch/domain"

	time "time"
)

// BookingUsecase is an autogenerated mock type for the BookingUsecase type
type BookingUsecase struct {
	mock.Mock
}

// Booking provides a mock function with given fields: ctx, bookID, bookAt, returnAt, quantity
func (_m *BookingUsecase) Booking(ctx context.Context, bookID string, bookAt time.Time, returnAt time.Time, quantity int) (domain.Booking, error) {
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

// Index provides a mock function with given fields: ctx, bookAt, returnAt
func (_m *BookingUsecase) Index(ctx context.Context, bookAt time.Time, returnAt time.Time) ([]domain.Booking, error) {
	ret := _m.Called(ctx, bookAt, returnAt)

	var r0 []domain.Booking
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []domain.Booking); ok {
		r0 = rf(ctx, bookAt, returnAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, bookAt, returnAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookingUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookingUsecase creates a new instance of BookingUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookingUsecase(t mockConstructorTestingTNewBookingUsecase) *BookingUsecase {
	mock := &BookingUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}