// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_bitkub is a generated GoMock package.
package mock_bitkub

import (
	context "context"
	reflect "reflect"

	bitkub "github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
	gomock "github.com/golang/mock/gomock"
)

// MockIBitkub is a mock of IBitkub interface.
type MockIBitkub struct {
	ctrl     *gomock.Controller
	recorder *MockIBitkubMockRecorder
}

// MockIBitkubMockRecorder is the mock recorder for MockIBitkub.
type MockIBitkubMockRecorder struct {
	mock *MockIBitkub
}

// NewMockIBitkub creates a new mock instance.
func NewMockIBitkub(ctrl *gomock.Controller) *MockIBitkub {
	mock := &MockIBitkub{ctrl: ctrl}
	mock.recorder = &MockIBitkubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBitkub) EXPECT() *MockIBitkubMockRecorder {
	return m.recorder
}

// GetTricker mocks base method.
func (m *MockIBitkub) GetTricker(ctx context.Context) (bitkub.GetTrickerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTricker", ctx)
	ret0, _ := ret[0].(bitkub.GetTrickerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTricker indicates an expected call of GetTricker.
func (mr *MockIBitkubMockRecorder) GetTricker(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTricker", reflect.TypeOf((*MockIBitkub)(nil).GetTricker), ctx)
}

// GetWallet mocks base method.
func (m *MockIBitkub) GetWallet(ctx context.Context) (bitkub.GetWalletResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", ctx)
	ret0, _ := ret[0].(bitkub.GetWalletResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockIBitkubMockRecorder) GetWallet(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockIBitkub)(nil).GetWallet), ctx)
}