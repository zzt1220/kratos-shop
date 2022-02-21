// Code generated by MockGen. DO NOT EDIT.
// Source: user/internal/biz (interfaces: AddressRepo)

// Package mrepo is a generated GoMock package.
package mrepo

import (
	context "context"
	reflect "reflect"
	biz "user/internal/biz"

	gomock "github.com/golang/mock/gomock"
)

// MockAddressRepo is a mock of AddressRepo interface.
type MockAddressRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAddressRepoMockRecorder
}

// MockAddressRepoMockRecorder is the mock recorder for MockAddressRepo.
type MockAddressRepoMockRecorder struct {
	mock *MockAddressRepo
}

// NewMockAddressRepo creates a new mock instance.
func NewMockAddressRepo(ctrl *gomock.Controller) *MockAddressRepo {
	mock := &MockAddressRepo{ctrl: ctrl}
	mock.recorder = &MockAddressRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddressRepo) EXPECT() *MockAddressRepoMockRecorder {
	return m.recorder
}

// AddressListByUid mocks base method.
func (m *MockAddressRepo) AddressListByUid(arg0 context.Context, arg1 int64) ([]*biz.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressListByUid", arg0, arg1)
	ret0, _ := ret[0].([]*biz.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddressListByUid indicates an expected call of AddressListByUid.
func (mr *MockAddressRepoMockRecorder) AddressListByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressListByUid", reflect.TypeOf((*MockAddressRepo)(nil).AddressListByUid), arg0, arg1)
}

// CreateAddress mocks base method.
func (m *MockAddressRepo) CreateAddress(arg0 context.Context, arg1 *biz.Address) (*biz.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddress", arg0, arg1)
	ret0, _ := ret[0].(*biz.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAddress indicates an expected call of CreateAddress.
func (mr *MockAddressRepoMockRecorder) CreateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockAddressRepo)(nil).CreateAddress), arg0, arg1)
}

// DefaultAddress mocks base method.
func (m *MockAddressRepo) DefaultAddress(arg0 context.Context, arg1 *biz.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DefaultAddress indicates an expected call of DefaultAddress.
func (mr *MockAddressRepoMockRecorder) DefaultAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultAddress", reflect.TypeOf((*MockAddressRepo)(nil).DefaultAddress), arg0, arg1)
}

// DeleteAddress mocks base method.
func (m *MockAddressRepo) DeleteAddress(arg0 context.Context, arg1 *biz.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockAddressRepoMockRecorder) DeleteAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockAddressRepo)(nil).DeleteAddress), arg0, arg1)
}

// UpdateAddress mocks base method.
func (m *MockAddressRepo) UpdateAddress(arg0 context.Context, arg1 *biz.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockAddressRepoMockRecorder) UpdateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockAddressRepo)(nil).UpdateAddress), arg0, arg1)
}
