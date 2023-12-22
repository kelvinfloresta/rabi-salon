// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	database "rabi-salon/frameworks/database"

	mock "github.com/stretchr/testify/mock"

	user_gateway "rabi-salon/frameworks/database/gateways/user_gateway"
)

// UserGateway is an autogenerated mock type for the UserGateway type
type UserGateway struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *UserGateway) Create(input user_gateway.CreateInput) (string, error) {
	ret := _m.Called(input)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(user_gateway.CreateInput) (string, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(user_gateway.CreateInput) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(user_gateway.CreateInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *UserGateway) Delete(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *UserGateway) GetByID(id string) (*user_gateway.GetByIDOutput, error) {
	ret := _m.Called(id)

	var r0 *user_gateway.GetByIDOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user_gateway.GetByIDOutput, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *user_gateway.GetByIDOutput); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user_gateway.GetByIDOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Paginate provides a mock function with given fields: filter, paginate
func (_m *UserGateway) Paginate(filter user_gateway.PaginateFilter, paginate database.PaginateInput) (*user_gateway.PaginateOutput, error) {
	ret := _m.Called(filter, paginate)

	var r0 *user_gateway.PaginateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(user_gateway.PaginateFilter, database.PaginateInput) (*user_gateway.PaginateOutput, error)); ok {
		return rf(filter, paginate)
	}
	if rf, ok := ret.Get(0).(func(user_gateway.PaginateFilter, database.PaginateInput) *user_gateway.PaginateOutput); ok {
		r0 = rf(filter, paginate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user_gateway.PaginateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(user_gateway.PaginateFilter, database.PaginateInput) error); ok {
		r1 = rf(filter, paginate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: filter, values
func (_m *UserGateway) Patch(filter user_gateway.PatchFilter, values user_gateway.PatchValues) (bool, error) {
	ret := _m.Called(filter, values)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(user_gateway.PatchFilter, user_gateway.PatchValues) (bool, error)); ok {
		return rf(filter, values)
	}
	if rf, ok := ret.Get(0).(func(user_gateway.PatchFilter, user_gateway.PatchValues) bool); ok {
		r0 = rf(filter, values)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(user_gateway.PatchFilter, user_gateway.PatchValues) error); ok {
		r1 = rf(filter, values)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserGateway creates a new instance of UserGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserGateway(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserGateway {
	mock := &UserGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
