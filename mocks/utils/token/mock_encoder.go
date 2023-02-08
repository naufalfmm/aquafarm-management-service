// Code generated by MockGen. DO NOT EDIT.
// Source: utils/token/encoder.go

// Package mock_token is a generated GoMock package.
package mock_token

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	token "github.com/naufalfmm/aquafarm-management-service/utils/token"
)

// MockEncoder is a mock of Encoder interface.
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder.
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance.
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// EncodeToken mocks base method.
func (m *MockEncoder) EncodeToken(claims token.Claims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncodeToken", claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncodeToken indicates an expected call of EncodeToken.
func (mr *MockEncoderMockRecorder) EncodeToken(claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeToken", reflect.TypeOf((*MockEncoder)(nil).EncodeToken), claims)
}
