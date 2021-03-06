// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import kv "chainspace.io/chainspace-go/kv"
import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *Service) Get(key []byte) ([]byte, error) {
	ret := _m.Called(key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPrefix provides a mock function with given fields: prefix
func (_m *Service) GetByPrefix(prefix []byte) ([]kv.ObjectEntry, error) {
	ret := _m.Called(prefix)

	var r0 []kv.ObjectEntry
	if rf, ok := ret.Get(0).(func([]byte) []kv.ObjectEntry); ok {
		r0 = rf(prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kv.ObjectEntry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, value
func (_m *Service) Set(key []byte, value []byte) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
