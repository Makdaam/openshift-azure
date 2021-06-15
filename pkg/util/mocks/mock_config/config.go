// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/openshift/openshift-azure/pkg/api"
	plugin "github.com/openshift/openshift-azure/pkg/api/plugin"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockInterface) Generate(template *plugin.Config, setVersionFields bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", template, setVersionFields)
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockInterfaceMockRecorder) Generate(template, setVersionFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockInterface)(nil).Generate), template, setVersionFields)
}

// GenerateStartup mocks base method.
func (m *MockInterface) GenerateStartup() (*api.OpenShiftManagedCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateStartup")
	ret0, _ := ret[0].(*api.OpenShiftManagedCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateStartup indicates an expected call of GenerateStartup.
func (mr *MockInterfaceMockRecorder) GenerateStartup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateStartup", reflect.TypeOf((*MockInterface)(nil).GenerateStartup))
}

// InvalidateCertificates mocks base method.
func (m *MockInterface) InvalidateCertificates() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateCertificates")
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateCertificates indicates an expected call of InvalidateCertificates.
func (mr *MockInterfaceMockRecorder) InvalidateCertificates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateCertificates", reflect.TypeOf((*MockInterface)(nil).InvalidateCertificates))
}

// InvalidateSecrets mocks base method.
func (m *MockInterface) InvalidateSecrets() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateSecrets")
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateSecrets indicates an expected call of InvalidateSecrets.
func (mr *MockInterfaceMockRecorder) InvalidateSecrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateSecrets", reflect.TypeOf((*MockInterface)(nil).InvalidateSecrets))
}
