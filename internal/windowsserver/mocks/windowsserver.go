// Code generated by MockGen. DO NOT EDIT.
// Source: windowsserver.go

// Package mocks is a generated GoMock package.
package mocks

import (
	windowsserver "pluggin-engine/internal/windowsserver"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetWindowsServerData mocks base method.
func (m *MockService) GetWindowsServerData() (windowsserver.MockWindowsServerData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWindowsServerData")
	ret0, _ := ret[0].(windowsserver.MockWindowsServerData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWindowsServerData indicates an expected call of GetWindowsServerData.
func (mr *MockServiceMockRecorder) GetWindowsServerData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWindowsServerData", reflect.TypeOf((*MockService)(nil).GetWindowsServerData))
}
