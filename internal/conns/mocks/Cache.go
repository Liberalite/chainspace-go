// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import service "chainspace.io/chainspace-go/service"
import time "time"

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Cache) Close() {
	_m.Called()
}

// WriteRequest provides a mock function with given fields: nodeID, msg, timeout, ack, cb
func (_m *Cache) WriteRequest(nodeID uint64, msg *service.Message, timeout time.Duration, ack bool, cb func(uint64, *service.Message)) (uint64, error) {
	ret := _m.Called(nodeID, msg, timeout, ack, cb)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64, *service.Message, time.Duration, bool, func(uint64, *service.Message)) uint64); ok {
		r0 = rf(nodeID, msg, timeout, ack, cb)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, *service.Message, time.Duration, bool, func(uint64, *service.Message)) error); ok {
		r1 = rf(nodeID, msg, timeout, ack, cb)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
