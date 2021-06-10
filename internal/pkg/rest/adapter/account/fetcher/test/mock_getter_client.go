// Code generated by MockGen. DO NOT EDIT.
// Source: getter.go

// Package mock_getter is a generated GoMock package.
package test

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockRestGetter is a mock of RestGetter interface
type MockRestGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRestGetterMockRecorder
}

// MockRestGetterMockRecorder is the mock recorder for MockRestGetter
type MockRestGetterMockRecorder struct {
	mock *MockRestGetter
}

// NewMockRestGetter creates a new mock instance
func NewMockRestGetter(ctrl *gomock.Controller) *MockRestGetter {
	mock := &MockRestGetter{ctrl: ctrl}
	mock.recorder = &MockRestGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRestGetter) EXPECT() *MockRestGetterMockRecorder {
	return m.recorder
}

// MakeGetRequest mocks base method
func (m *MockRestGetter) MakeGetRequest(url string) ([]byte, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeGetRequest", url)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MakeGetRequest indicates an expected call of MakeGetRequest
func (mr *MockRestGetterMockRecorder) MakeGetRequest(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeGetRequest", reflect.TypeOf((*MockRestGetter)(nil).MakeGetRequest), url)
}