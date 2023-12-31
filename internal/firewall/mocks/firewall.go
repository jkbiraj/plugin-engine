// Code generated by MockGen. DO NOT EDIT.
// Source: firewall.go

// Package mocks is a generated GoMock package.
package mocks

import (
	firewall "pluggin-engine/internal/firewall"
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

// GetFirewallData mocks base method.
func (m *MockService) GetFirewallData() (firewall.MockFirewallData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewallData")
	ret0, _ := ret[0].(firewall.MockFirewallData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewallData indicates an expected call of GetFirewallData.
func (mr *MockServiceMockRecorder) GetFirewallData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewallData", reflect.TypeOf((*MockService)(nil).GetFirewallData))
}
