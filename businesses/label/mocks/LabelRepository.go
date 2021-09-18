// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	label "keep-remind-app/businesses/label"

	mock "github.com/stretchr/testify/mock"
)

// LabelRepository is an autogenerated mock type for the LabelRepository type
type LabelRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, data
func (_m *LabelRepository) Add(ctx context.Context, data *label.LabelDomain) (int, error) {
	ret := _m.Called(ctx, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *label.LabelDomain) int); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *label.LabelDomain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *LabelRepository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: ctx, data
func (_m *LabelRepository) Edit(ctx context.Context, data *label.LabelDomain) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *label.LabelDomain) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, param
func (_m *LabelRepository) FindAll(ctx context.Context, param *label.LabelParameter) ([]label.LabelDomain, error) {
	ret := _m.Called(ctx, param)

	var r0 []label.LabelDomain
	if rf, ok := ret.Get(0).(func(context.Context, *label.LabelParameter) []label.LabelDomain); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]label.LabelDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *label.LabelParameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllPagination provides a mock function with given fields: ctx, param
func (_m *LabelRepository) FindAllPagination(ctx context.Context, param *label.LabelParameter) ([]label.LabelDomain, int, error) {
	ret := _m.Called(ctx, param)

	var r0 []label.LabelDomain
	if rf, ok := ret.Get(0).(func(context.Context, *label.LabelParameter) []label.LabelDomain); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]label.LabelDomain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, *label.LabelParameter) int); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *label.LabelParameter) error); ok {
		r2 = rf(ctx, param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindOne provides a mock function with given fields: ctx, param
func (_m *LabelRepository) FindOne(ctx context.Context, param *label.LabelParameter) (label.LabelDomain, error) {
	ret := _m.Called(ctx, param)

	var r0 label.LabelDomain
	if rf, ok := ret.Get(0).(func(context.Context, *label.LabelParameter) label.LabelDomain); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(label.LabelDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *label.LabelParameter) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
