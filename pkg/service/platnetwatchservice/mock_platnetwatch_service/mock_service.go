// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_platnetwatch_service is a generated GoMock package.
package mock_platnetwatch_service

import (
	context "context"
	reflect "reflect"

	platnetwatchservice "github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
	gomock "github.com/golang/mock/gomock"
)

// MockIPlanetwatchService is a mock of IPlanetwatchService interface.
type MockIPlanetwatchService struct {
	ctrl     *gomock.Controller
	recorder *MockIPlanetwatchServiceMockRecorder
}

// MockIPlanetwatchServiceMockRecorder is the mock recorder for MockIPlanetwatchService.
type MockIPlanetwatchServiceMockRecorder struct {
	mock *MockIPlanetwatchService
}

// NewMockIPlanetwatchService creates a new mock instance.
func NewMockIPlanetwatchService(ctrl *gomock.Controller) *MockIPlanetwatchService {
	mock := &MockIPlanetwatchService{ctrl: ctrl}
	mock.recorder = &MockIPlanetwatchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPlanetwatchService) EXPECT() *MockIPlanetwatchServiceMockRecorder {
	return m.recorder
}

// GetIncome mocks base method.
func (m *MockIPlanetwatchService) GetIncome(ctx context.Context) ([]*platnetwatchservice.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncome", ctx)
	ret0, _ := ret[0].([]*platnetwatchservice.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncome indicates an expected call of GetIncome.
func (mr *MockIPlanetwatchServiceMockRecorder) GetIncome(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncome", reflect.TypeOf((*MockIPlanetwatchService)(nil).GetIncome), ctx)
}
