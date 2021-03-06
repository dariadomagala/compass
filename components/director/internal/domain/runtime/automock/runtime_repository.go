// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import labelfilter "github.com/kyma-incubator/compass/components/director/internal/labelfilter"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// RuntimeRepository is an autogenerated mock type for the RuntimeRepository type
type RuntimeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: item
func (_m *RuntimeRepository) Create(item *model.Runtime) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Runtime) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: item
func (_m *RuntimeRepository) Delete(item *model.Runtime) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Runtime) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *RuntimeRepository) GetByID(id string) (*model.Runtime, error) {
	ret := _m.Called(id)

	var r0 *model.Runtime
	if rf, ok := ret.Get(0).(func(string) *model.Runtime); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Runtime)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: filter, pageSize, cursor
func (_m *RuntimeRepository) List(filter []*labelfilter.LabelFilter, pageSize *int, cursor *string) (*model.RuntimePage, error) {
	ret := _m.Called(filter, pageSize, cursor)

	var r0 *model.RuntimePage
	if rf, ok := ret.Get(0).(func([]*labelfilter.LabelFilter, *int, *string) *model.RuntimePage); ok {
		r0 = rf(filter, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RuntimePage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*labelfilter.LabelFilter, *int, *string) error); ok {
		r1 = rf(filter, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: item
func (_m *RuntimeRepository) Update(item *model.Runtime) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Runtime) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
