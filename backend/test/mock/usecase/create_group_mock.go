// Code generated by MockGen. DO NOT EDIT.
// Source: create_group.go
//
// Generated by this command:
//
//	mockgen -source=create_group.go -destination=../test/mock/usecase/create_group_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	usecase "backend/usecase"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCreateGroup is a mock of CreateGroup interface.
type MockCreateGroup struct {
	ctrl     *gomock.Controller
	recorder *MockCreateGroupMockRecorder
	isgomock struct{}
}

// MockCreateGroupMockRecorder is the mock recorder for MockCreateGroup.
type MockCreateGroupMockRecorder struct {
	mock *MockCreateGroup
}

// NewMockCreateGroup creates a new mock instance.
func NewMockCreateGroup(ctrl *gomock.Controller) *MockCreateGroup {
	mock := &MockCreateGroup{ctrl: ctrl}
	mock.recorder = &MockCreateGroupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateGroup) EXPECT() *MockCreateGroupMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCreateGroup) Execute(ctx context.Context, in usecase.CreateGroupInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockCreateGroupMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateGroup)(nil).Execute), ctx, in)
}
