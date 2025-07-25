// Code generated by MockGen. DO NOT EDIT.
// Source: get_account.go
//
// Generated by this command:
//
//	mockgen -source=get_account.go -destination=../test/mock/usecase/get_account_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	usecase "backend/usecase"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGetAccount is a mock of GetAccount interface.
type MockGetAccount struct {
	ctrl     *gomock.Controller
	recorder *MockGetAccountMockRecorder
	isgomock struct{}
}

// MockGetAccountMockRecorder is the mock recorder for MockGetAccount.
type MockGetAccountMockRecorder struct {
	mock *MockGetAccount
}

// NewMockGetAccount creates a new mock instance.
func NewMockGetAccount(ctrl *gomock.Controller) *MockGetAccount {
	mock := &MockGetAccount{ctrl: ctrl}
	mock.recorder = &MockGetAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetAccount) EXPECT() *MockGetAccountMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetAccount) Execute(ctx context.Context, input usecase.GetAccountInput) (usecase.GetAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, input)
	ret0, _ := ret[0].(usecase.GetAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetAccountMockRecorder) Execute(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetAccount)(nil).Execute), ctx, input)
}
