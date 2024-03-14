// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: DataFederationQueryLimitLister,DataFederationQueryLimitDescriber,DataFederationQueryLimitCreator,DataFederationQueryLimitDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

// MockDataFederationQueryLimitLister is a mock of DataFederationQueryLimitLister interface.
type MockDataFederationQueryLimitLister struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationQueryLimitListerMockRecorder
}

// MockDataFederationQueryLimitListerMockRecorder is the mock recorder for MockDataFederationQueryLimitLister.
type MockDataFederationQueryLimitListerMockRecorder struct {
	mock *MockDataFederationQueryLimitLister
}

// NewMockDataFederationQueryLimitLister creates a new mock instance.
func NewMockDataFederationQueryLimitLister(ctrl *gomock.Controller) *MockDataFederationQueryLimitLister {
	mock := &MockDataFederationQueryLimitLister{ctrl: ctrl}
	mock.recorder = &MockDataFederationQueryLimitListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationQueryLimitLister) EXPECT() *MockDataFederationQueryLimitListerMockRecorder {
	return m.recorder
}

// DataFederationQueryLimits mocks base method.
func (m *MockDataFederationQueryLimitLister) DataFederationQueryLimits(arg0, arg1 string) ([]admin.DataFederationTenantQueryLimit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationQueryLimits", arg0, arg1)
	ret0, _ := ret[0].([]admin.DataFederationTenantQueryLimit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationQueryLimits indicates an expected call of DataFederationQueryLimits.
func (mr *MockDataFederationQueryLimitListerMockRecorder) DataFederationQueryLimits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationQueryLimits", reflect.TypeOf((*MockDataFederationQueryLimitLister)(nil).DataFederationQueryLimits), arg0, arg1)
}

// MockDataFederationQueryLimitDescriber is a mock of DataFederationQueryLimitDescriber interface.
type MockDataFederationQueryLimitDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationQueryLimitDescriberMockRecorder
}

// MockDataFederationQueryLimitDescriberMockRecorder is the mock recorder for MockDataFederationQueryLimitDescriber.
type MockDataFederationQueryLimitDescriberMockRecorder struct {
	mock *MockDataFederationQueryLimitDescriber
}

// NewMockDataFederationQueryLimitDescriber creates a new mock instance.
func NewMockDataFederationQueryLimitDescriber(ctrl *gomock.Controller) *MockDataFederationQueryLimitDescriber {
	mock := &MockDataFederationQueryLimitDescriber{ctrl: ctrl}
	mock.recorder = &MockDataFederationQueryLimitDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationQueryLimitDescriber) EXPECT() *MockDataFederationQueryLimitDescriberMockRecorder {
	return m.recorder
}

// DataFederationQueryLimit mocks base method.
func (m *MockDataFederationQueryLimitDescriber) DataFederationQueryLimit(arg0, arg1, arg2 string) (*admin.DataFederationTenantQueryLimit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationQueryLimit", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DataFederationTenantQueryLimit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationQueryLimit indicates an expected call of DataFederationQueryLimit.
func (mr *MockDataFederationQueryLimitDescriberMockRecorder) DataFederationQueryLimit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationQueryLimit", reflect.TypeOf((*MockDataFederationQueryLimitDescriber)(nil).DataFederationQueryLimit), arg0, arg1, arg2)
}

// MockDataFederationQueryLimitCreator is a mock of DataFederationQueryLimitCreator interface.
type MockDataFederationQueryLimitCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationQueryLimitCreatorMockRecorder
}

// MockDataFederationQueryLimitCreatorMockRecorder is the mock recorder for MockDataFederationQueryLimitCreator.
type MockDataFederationQueryLimitCreatorMockRecorder struct {
	mock *MockDataFederationQueryLimitCreator
}

// NewMockDataFederationQueryLimitCreator creates a new mock instance.
func NewMockDataFederationQueryLimitCreator(ctrl *gomock.Controller) *MockDataFederationQueryLimitCreator {
	mock := &MockDataFederationQueryLimitCreator{ctrl: ctrl}
	mock.recorder = &MockDataFederationQueryLimitCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationQueryLimitCreator) EXPECT() *MockDataFederationQueryLimitCreatorMockRecorder {
	return m.recorder
}

// CreateDataFederationQueryLimit mocks base method.
func (m *MockDataFederationQueryLimitCreator) CreateDataFederationQueryLimit(arg0, arg1, arg2 string, arg3 *admin.DataFederationTenantQueryLimit) (*admin.DataFederationTenantQueryLimit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataFederationQueryLimit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*admin.DataFederationTenantQueryLimit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataFederationQueryLimit indicates an expected call of CreateDataFederationQueryLimit.
func (mr *MockDataFederationQueryLimitCreatorMockRecorder) CreateDataFederationQueryLimit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataFederationQueryLimit", reflect.TypeOf((*MockDataFederationQueryLimitCreator)(nil).CreateDataFederationQueryLimit), arg0, arg1, arg2, arg3)
}

// MockDataFederationQueryLimitDeleter is a mock of DataFederationQueryLimitDeleter interface.
type MockDataFederationQueryLimitDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDataFederationQueryLimitDeleterMockRecorder
}

// MockDataFederationQueryLimitDeleterMockRecorder is the mock recorder for MockDataFederationQueryLimitDeleter.
type MockDataFederationQueryLimitDeleterMockRecorder struct {
	mock *MockDataFederationQueryLimitDeleter
}

// NewMockDataFederationQueryLimitDeleter creates a new mock instance.
func NewMockDataFederationQueryLimitDeleter(ctrl *gomock.Controller) *MockDataFederationQueryLimitDeleter {
	mock := &MockDataFederationQueryLimitDeleter{ctrl: ctrl}
	mock.recorder = &MockDataFederationQueryLimitDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataFederationQueryLimitDeleter) EXPECT() *MockDataFederationQueryLimitDeleterMockRecorder {
	return m.recorder
}

// DeleteDataFederationQueryLimit mocks base method.
func (m *MockDataFederationQueryLimitDeleter) DeleteDataFederationQueryLimit(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataFederationQueryLimit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataFederationQueryLimit indicates an expected call of DeleteDataFederationQueryLimit.
func (mr *MockDataFederationQueryLimitDeleterMockRecorder) DeleteDataFederationQueryLimit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataFederationQueryLimit", reflect.TypeOf((*MockDataFederationQueryLimitDeleter)(nil).DeleteDataFederationQueryLimit), arg0, arg1, arg2)
}
