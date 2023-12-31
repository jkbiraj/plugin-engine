// Code generated by MockGen. DO NOT EDIT.
// Source: router.go

// Package mocks is a generated GoMock package.
package mocks

import (
	router "pluggin-engine/internal/router"
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

// GetRouterData mocks base method.
func (m *MockService) GetRouterData() (router.MockRouterData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouterData")
	ret0, _ := ret[0].(router.MockRouterData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRouterData indicates an expected call of GetRouterData.
func (mr *MockServiceMockRecorder) GetRouterData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouterData", reflect.TypeOf((*MockService)(nil).GetRouterData))
}
