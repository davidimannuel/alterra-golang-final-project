// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	telegramUser "keep-remind-app/businesses/telegramUser"

	mock "github.com/stretchr/testify/mock"
)

// TelegramUserRepository is an autogenerated mock type for the TelegramUserRepository type
type TelegramUserRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, data
func (_m *TelegramUserRepository) Add(ctx context.Context, data *telegramUser.TelegramUserDomain) (int, error) {
	ret := _m.Called(ctx, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *telegramUser.TelegramUserDomain) int); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *telegramUser.TelegramUserDomain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, data
func (_m *TelegramUserRepository) Delete(ctx context.Context, data *telegramUser.TelegramUserDomain) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *telegramUser.TelegramUserDomain) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditStatus provides a mock function with given fields: ctx, data
func (_m *TelegramUserRepository) EditStatus(ctx context.Context, data *telegramUser.TelegramUserDomain) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *telegramUser.TelegramUserDomain) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, param
func (_m *TelegramUserRepository) FindAll(ctx context.Context, param *telegramUser.TelegramUserParameter) ([]telegramUser.TelegramUserDomain, error) {
	ret := _m.Called(ctx, param)

	var r0 []telegramUser.TelegramUserDomain
	if rf, ok := ret.Get(0).(func(context.Context, *telegramUser.TelegramUserParameter) []telegramUser.TelegramUserDomain); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]telegramUser.TelegramUserDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *telegramUser.TelegramUserParameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllPagination provides a mock function with given fields: ctx, param
func (_m *TelegramUserRepository) FindAllPagination(ctx context.Context, param *telegramUser.TelegramUserParameter) ([]telegramUser.TelegramUserDomain, int, error) {
	ret := _m.Called(ctx, param)

	var r0 []telegramUser.TelegramUserDomain
	if rf, ok := ret.Get(0).(func(context.Context, *telegramUser.TelegramUserParameter) []telegramUser.TelegramUserDomain); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]telegramUser.TelegramUserDomain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, *telegramUser.TelegramUserParameter) int); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *telegramUser.TelegramUserParameter) error); ok {
		r2 = rf(ctx, param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindOne provides a mock function with given fields: ctx, param
func (_m *TelegramUserRepository) FindOne(ctx context.Context, param *telegramUser.TelegramUserParameter) (telegramUser.TelegramUserDomain, error) {
	ret := _m.Called(ctx, param)

	var r0 telegramUser.TelegramUserDomain
	if rf, ok := ret.Get(0).(func(context.Context, *telegramUser.TelegramUserParameter) telegramUser.TelegramUserDomain); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(telegramUser.TelegramUserDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *telegramUser.TelegramUserParameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
