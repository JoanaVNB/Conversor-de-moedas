// Code generated by MockGen. DO NOT EDIT.
// Source: delete.go

// Package handlers is a generated GoMock package.
package handlers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeleteHandler is a mock of DeleteHandler interface.
type MockDeleteHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteHandlerMockRecorder
}

// MockDeleteHandlerMockRecorder is the mock recorder for MockDeleteHandler.
type MockDeleteHandlerMockRecorder struct {
	mock *MockDeleteHandler
}

// NewMockDeleteHandler creates a new mock instance.
func NewMockDeleteHandler(ctrl *gomock.Controller) *MockDeleteHandler {
	mock := &MockDeleteHandler{ctrl: ctrl}
	mock.recorder = &MockDeleteHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteHandler) EXPECT() *MockDeleteHandlerMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockDeleteHandler) Execute(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockDeleteHandlerMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDeleteHandler)(nil).Execute), arg0)
}
