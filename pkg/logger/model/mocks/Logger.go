// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	log "log"

	model "github.com/goLogOverCoat/pkg/logger/model"
	mock "github.com/stretchr/testify/mock"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: msg
func (_m *Logger) Debug(msg string) {
	_m.Called(msg)
}

// Debugf provides a mock function with given fields: format, args
func (_m *Logger) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: msg
func (_m *Logger) Error(msg string) {
	_m.Called(msg)
}

// Errorf provides a mock function with given fields: format, args
func (_m *Logger) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: msg
func (_m *Logger) Info(msg string) {
	_m.Called(msg)
}

// Infof provides a mock function with given fields: format, args
func (_m *Logger) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// ToStdLogger provides a mock function with given fields:
func (_m *Logger) ToStdLogger() *log.Logger {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ToStdLogger")
	}

	var r0 *log.Logger
	if rf, ok := ret.Get(0).(func() *log.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*log.Logger)
		}
	}

	return r0
}

// Warn provides a mock function with given fields: msg
func (_m *Logger) Warn(msg string) {
	_m.Called(msg)
}

// Warnf provides a mock function with given fields: format, args
func (_m *Logger) Warnf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WithError provides a mock function with given fields: err
func (_m *Logger) WithError(err error) model.Logger {
	ret := _m.Called(err)

	if len(ret) == 0 {
		panic("no return value specified for WithError")
	}

	var r0 model.Logger
	if rf, ok := ret.Get(0).(func(error) model.Logger); ok {
		r0 = rf(err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Logger)
		}
	}

	return r0
}

// WithField provides a mock function with given fields: key, value
func (_m *Logger) WithField(key string, value interface{}) model.Logger {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for WithField")
	}

	var r0 model.Logger
	if rf, ok := ret.Get(0).(func(string, interface{}) model.Logger); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Logger)
		}
	}

	return r0
}

// WithFields provides a mock function with given fields: fields
func (_m *Logger) WithFields(fields model.Fields) model.Logger {
	ret := _m.Called(fields)

	if len(ret) == 0 {
		panic("no return value specified for WithFields")
	}

	var r0 model.Logger
	if rf, ok := ret.Get(0).(func(model.Fields) model.Logger); ok {
		r0 = rf(fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Logger)
		}
	}

	return r0
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
