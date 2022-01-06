// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_line is a generated GoMock package.
package mock_line

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	linebot "github.com/line/line-bot-sdk-go/v7/linebot"
)

// MockILine is a mock of ILine interface.
type MockILine struct {
	ctrl     *gomock.Controller
	recorder *MockILineMockRecorder
}

// MockILineMockRecorder is the mock recorder for MockILine.
type MockILineMockRecorder struct {
	mock *MockILine
}

// NewMockILine creates a new mock instance.
func NewMockILine(ctrl *gomock.Controller) *MockILine {
	mock := &MockILine{ctrl: ctrl}
	mock.recorder = &MockILineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILine) EXPECT() *MockILineMockRecorder {
	return m.recorder
}

// PushMessage mocks base method.
func (m *MockILine) PushMessage(ctx context.Context, container *linebot.BubbleContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushMessage", ctx, container)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushMessage indicates an expected call of PushMessage.
func (mr *MockILineMockRecorder) PushMessage(ctx, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMessage", reflect.TypeOf((*MockILine)(nil).PushMessage), ctx, container)
}

// ReplyTextMessage mocks base method.
func (m *MockILine) ReplyTextMessage(ctx context.Context, token, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyTextMessage", ctx, token, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyTextMessage indicates an expected call of ReplyTextMessage.
func (mr *MockILineMockRecorder) ReplyTextMessage(ctx, token, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyTextMessage", reflect.TypeOf((*MockILine)(nil).ReplyTextMessage), ctx, token, message)
}

// SendFlexMessage mocks base method.
func (m *MockILine) SendFlexMessage(ctx context.Context, token string, container *linebot.BubbleContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendFlexMessage", ctx, token, container)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendFlexMessage indicates an expected call of SendFlexMessage.
func (mr *MockILineMockRecorder) SendFlexMessage(ctx, token, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendFlexMessage", reflect.TypeOf((*MockILine)(nil).SendFlexMessage), ctx, token, container)
}
