// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/niqinge/order/dao (interfaces: OrderDAO)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/niqinge/order/model"
	reflect "reflect"
)

// MockOrderDAO is a mock of OrderDAO interface
type MockOrderDAO struct {
	ctrl     *gomock.Controller
	recorder *MockOrderDAOMockRecorder
}

// MockOrderDAOMockRecorder is the mock recorder for MockOrderDAO
type MockOrderDAOMockRecorder struct {
	mock *MockOrderDAO
}

// NewMockOrderDAO creates a new mock instance
func NewMockOrderDAO(ctrl *gomock.Controller) *MockOrderDAO {
	mock := &MockOrderDAO{ctrl: ctrl}
	mock.recorder = &MockOrderDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderDAO) EXPECT() *MockOrderDAOMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockOrderDAO) Create(arg0 *model.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockOrderDAOMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderDAO)(nil).Create), arg0)
}

// QueryByNo mocks base method
func (m *MockOrderDAO) QueryByNo(arg0 string) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryByNo", arg0)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryByNo indicates an expected call of QueryByNo
func (mr *MockOrderDAOMockRecorder) QueryByNo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByNo", reflect.TypeOf((*MockOrderDAO)(nil).QueryByNo), arg0)
}

// QueryList mocks base method
func (m *MockOrderDAO) QueryList(arg0, arg1 int) ([]*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryList", arg0, arg1)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryList indicates an expected call of QueryList
func (mr *MockOrderDAOMockRecorder) QueryList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryList", reflect.TypeOf((*MockOrderDAO)(nil).QueryList), arg0, arg1)
}

// UpdateByNo mocks base method
func (m *MockOrderDAO) UpdateByNo(arg0 string, arg1 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByNo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByNo indicates an expected call of UpdateByNo
func (mr *MockOrderDAOMockRecorder) UpdateByNo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByNo", reflect.TypeOf((*MockOrderDAO)(nil).UpdateByNo), arg0, arg1)
}