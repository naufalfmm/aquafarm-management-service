// Code generated by MockGen. DO NOT EDIT.
// Source: persistents/repositories/endpoints/index.go

// Package mock_endpointsRepositories is a generated GoMock package.
package mock_endpointsRepositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/naufalfmm/aquafarm-management-service/model/dao"
)

// MockRepositories is a mock of Repositories interface.
type MockRepositories struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoriesMockRecorder
}

// MockRepositoriesMockRecorder is the mock recorder for MockRepositories.
type MockRepositoriesMockRecorder struct {
	mock *MockRepositories
}

// NewMockRepositories creates a new mock instance.
func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	mock := &MockRepositories{ctrl: ctrl}
	mock.recorder = &MockRepositoriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositories) EXPECT() *MockRepositoriesMockRecorder {
	return m.recorder
}

// BulkCreate mocks base method.
func (m *MockRepositories) BulkCreate(ctx context.Context, data dao.Endpoints) (dao.Endpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", ctx, data)
	ret0, _ := ret[0].(dao.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockRepositoriesMockRecorder) BulkCreate(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockRepositories)(nil).BulkCreate), ctx, data)
}

// BulkDeleteByIDs mocks base method.
func (m *MockRepositories) BulkDeleteByIDs(ctx context.Context, ids []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteByIDs", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteByIDs indicates an expected call of BulkDeleteByIDs.
func (mr *MockRepositoriesMockRecorder) BulkDeleteByIDs(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteByIDs", reflect.TypeOf((*MockRepositories)(nil).BulkDeleteByIDs), ctx, ids)
}

// GetAll mocks base method.
func (m *MockRepositories) GetAll(ctx context.Context) (dao.Endpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(dao.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoriesMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepositories)(nil).GetAll), ctx)
}

// GetByMethodPath mocks base method.
func (m *MockRepositories) GetByMethodPath(ctx context.Context, method, path string) (dao.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMethodPath", ctx, method, path)
	ret0, _ := ret[0].(dao.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMethodPath indicates an expected call of GetByMethodPath.
func (mr *MockRepositoriesMockRecorder) GetByMethodPath(ctx, method, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMethodPath", reflect.TypeOf((*MockRepositories)(nil).GetByMethodPath), ctx, method, path)
}
