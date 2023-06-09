// Code generated by MockGen. DO NOT EDIT.
// Source: create.go

// Package handlers is a generated GoMock package.
package handlers

import (
	domain "exchange/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCreateHandler is a mock of CreateHandler interface.
type MockCreateHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCreateHandlerMockRecorder
}

// MockCreateHandlerMockRecorder is the mock recorder for MockCreateHandler.
type MockCreateHandlerMockRecorder struct {
	mock *MockCreateHandler
}

// NewMockCreateHandler creates a new mock instance.
func NewMockCreateHandler(ctrl *gomock.Controller) *MockCreateHandler {
	mock := &MockCreateHandler{ctrl: ctrl}
	mock.recorder = &MockCreateHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateHandler) EXPECT() *MockCreateHandlerMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCreateHandler) Execute(arg0 float64, arg1, arg2 string, arg3 float64) (domain.Exchange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(domain.Exchange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCreateHandlerMockRecorder) Execute(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateHandler)(nil).Execute), arg0, arg1, arg2, arg3)
}
