// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import os "os"

// DirectoryClient is an autogenerated mock type for the DirectoryClient type
type DirectoryClient struct {
	mock.Mock
}

// Dir provides a mock function with given fields:
func (_m *DirectoryClient) Dir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Mkdir provides a mock function with given fields: perm
func (_m *DirectoryClient) Mkdir(perm os.FileMode) error {
	ret := _m.Called(perm)

	var r0 error
	if rf, ok := ret.Get(0).(func(os.FileMode) error); ok {
		r0 = rf(perm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadDir provides a mock function with given fields:
func (_m *DirectoryClient) ReadDir() ([]os.FileInfo, error) {
	ret := _m.Called()

	var r0 []os.FileInfo
	if rf, ok := ret.Get(0).(func() []os.FileInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stat provides a mock function with given fields:
func (_m *DirectoryClient) Stat() (os.FileInfo, error) {
	ret := _m.Called()

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func() os.FileInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
