// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/cluster/kubeclient (interfaces: Interface)

// Package mock_kubeclient is a generated GoMock package.
package mock_kubeclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
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

// BackupCluster mocks base method.
func (m *MockInterface) BackupCluster(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BackupCluster indicates an expected call of BackupCluster.
func (mr *MockInterfaceMockRecorder) BackupCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupCluster", reflect.TypeOf((*MockInterface)(nil).BackupCluster), arg0, arg1)
}

// DeleteMaster mocks base method.
func (m *MockInterface) DeleteMaster(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMaster", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMaster indicates an expected call of DeleteMaster.
func (mr *MockInterfaceMockRecorder) DeleteMaster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMaster", reflect.TypeOf((*MockInterface)(nil).DeleteMaster), arg0)
}

// DrainAndDeleteWorker mocks base method.
func (m *MockInterface) DrainAndDeleteWorker(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DrainAndDeleteWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DrainAndDeleteWorker indicates an expected call of DrainAndDeleteWorker.
func (mr *MockInterfaceMockRecorder) DrainAndDeleteWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DrainAndDeleteWorker", reflect.TypeOf((*MockInterface)(nil).DrainAndDeleteWorker), arg0, arg1)
}

// EnsureSyncPod mocks base method.
func (m *MockInterface) EnsureSyncPod(arg0 context.Context, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureSyncPod", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureSyncPod indicates an expected call of EnsureSyncPod.
func (mr *MockInterfaceMockRecorder) EnsureSyncPod(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSyncPod", reflect.TypeOf((*MockInterface)(nil).EnsureSyncPod), arg0, arg1, arg2)
}

// GetControlPlanePods mocks base method.
func (m *MockInterface) GetControlPlanePods(arg0 context.Context) ([]v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlPlanePods", arg0)
	ret0, _ := ret[0].([]v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControlPlanePods indicates an expected call of GetControlPlanePods.
func (mr *MockInterfaceMockRecorder) GetControlPlanePods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlPlanePods", reflect.TypeOf((*MockInterface)(nil).GetControlPlanePods), arg0)
}

// RemoveSyncPod mocks base method.
func (m *MockInterface) RemoveSyncPod(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSyncPod", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSyncPod indicates an expected call of RemoveSyncPod.
func (mr *MockInterfaceMockRecorder) RemoveSyncPod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSyncPod", reflect.TypeOf((*MockInterface)(nil).RemoveSyncPod), arg0)
}

// RemoveValidatingWebhookConfiguration mocks base method.
func (m *MockInterface) RemoveValidatingWebhookConfiguration(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveValidatingWebhookConfiguration", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveValidatingWebhookConfiguration indicates an expected call of RemoveValidatingWebhookConfiguration.
func (mr *MockInterfaceMockRecorder) RemoveValidatingWebhookConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveValidatingWebhookConfiguration", reflect.TypeOf((*MockInterface)(nil).RemoveValidatingWebhookConfiguration), arg0)
}

// WaitForReadyMaster mocks base method.
func (m *MockInterface) WaitForReadyMaster(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForReadyMaster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForReadyMaster indicates an expected call of WaitForReadyMaster.
func (mr *MockInterfaceMockRecorder) WaitForReadyMaster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForReadyMaster", reflect.TypeOf((*MockInterface)(nil).WaitForReadyMaster), arg0, arg1)
}

// WaitForReadySyncPod mocks base method.
func (m *MockInterface) WaitForReadySyncPod(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForReadySyncPod", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForReadySyncPod indicates an expected call of WaitForReadySyncPod.
func (mr *MockInterfaceMockRecorder) WaitForReadySyncPod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForReadySyncPod", reflect.TypeOf((*MockInterface)(nil).WaitForReadySyncPod), arg0)
}

// WaitForReadyWorker mocks base method.
func (m *MockInterface) WaitForReadyWorker(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForReadyWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForReadyWorker indicates an expected call of WaitForReadyWorker.
func (mr *MockInterfaceMockRecorder) WaitForReadyWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForReadyWorker", reflect.TypeOf((*MockInterface)(nil).WaitForReadyWorker), arg0, arg1)
}
