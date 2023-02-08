// Code generated by MockGen. DO NOT EDIT.
// Source: utils/token/claims.go

// Package mock_token is a generated GoMock package.
package mock_token

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	token "github.com/naufalfmm/aquafarm-management-service/utils/token"
)

// MockClaims is a mock of Claims interface.
type MockClaims struct {
	ctrl     *gomock.Controller
	recorder *MockClaimsMockRecorder
}

// MockClaimsMockRecorder is the mock recorder for MockClaims.
type MockClaimsMockRecorder struct {
	mock *MockClaims
}

// NewMockClaims creates a new mock instance.
func NewMockClaims(ctrl *gomock.Controller) *MockClaims {
	mock := &MockClaims{ctrl: ctrl}
	mock.recorder = &MockClaimsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClaims) EXPECT() *MockClaimsMockRecorder {
	return m.recorder
}

// SetExp mocks base method.
func (m *MockClaims) SetExp(exp int64) token.Claims {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetExp", exp)
	ret0, _ := ret[0].(token.Claims)
	return ret0
}

// SetExp indicates an expected call of SetExp.
func (mr *MockClaimsMockRecorder) SetExp(exp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExp", reflect.TypeOf((*MockClaims)(nil).SetExp), exp)
}

// Valid mocks base method.
func (m *MockClaims) Valid() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Valid")
	ret0, _ := ret[0].(error)
	return ret0
}

// Valid indicates an expected call of Valid.
func (mr *MockClaimsMockRecorder) Valid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Valid", reflect.TypeOf((*MockClaims)(nil).Valid))
}