// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_elrond is a generated GoMock package.
package mock_elrond

import (
	context "context"
	reflect "reflect"

	elrond "github.com/dacharat/my-crypto-assets/pkg/external/elrond"
	gomock "github.com/golang/mock/gomock"
)

// MockIElrond is a mock of IElrond interface.
type MockIElrond struct {
	ctrl     *gomock.Controller
	recorder *MockIElrondMockRecorder
}

// MockIElrondMockRecorder is the mock recorder for MockIElrond.
type MockIElrondMockRecorder struct {
	mock *MockIElrond
}

// NewMockIElrond creates a new mock instance.
func NewMockIElrond(ctrl *gomock.Controller) *MockIElrond {
	mock := &MockIElrond{ctrl: ctrl}
	mock.recorder = &MockIElrondMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIElrond) EXPECT() *MockIElrondMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockIElrond) GetAccount(ctx context.Context, address string) (elrond.GetAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, address)
	ret0, _ := ret[0].(elrond.GetAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockIElrondMockRecorder) GetAccount(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockIElrond)(nil).GetAccount), ctx, address)
}

// GetAccountDelegation mocks base method.
func (m *MockIElrond) GetAccountDelegation(ctx context.Context, address string) ([]elrond.GetAccountDelegationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountDelegation", ctx, address)
	ret0, _ := ret[0].([]elrond.GetAccountDelegationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountDelegation indicates an expected call of GetAccountDelegation.
func (mr *MockIElrondMockRecorder) GetAccountDelegation(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDelegation", reflect.TypeOf((*MockIElrond)(nil).GetAccountDelegation), ctx, address)
}

// GetAccountNfts mocks base method.
func (m *MockIElrond) GetAccountNfts(ctx context.Context, address string) ([]elrond.GetAccountNftResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountNfts", ctx, address)
	ret0, _ := ret[0].([]elrond.GetAccountNftResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountNfts indicates an expected call of GetAccountNfts.
func (mr *MockIElrondMockRecorder) GetAccountNfts(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountNfts", reflect.TypeOf((*MockIElrond)(nil).GetAccountNfts), ctx, address)
}
