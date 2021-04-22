// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient/network (interfaces: VirtualNetworksClient,VirtualNetworksPeeringsClient,PublicIPAddressesClient)

// Package mock_network is a generated GoMock package.
package mock_network

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-06-01/network"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockVirtualNetworksClient is a mock of VirtualNetworksClient interface
type MockVirtualNetworksClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworksClientMockRecorder
}

// MockVirtualNetworksClientMockRecorder is the mock recorder for MockVirtualNetworksClient
type MockVirtualNetworksClientMockRecorder struct {
	mock *MockVirtualNetworksClient
}

// NewMockVirtualNetworksClient creates a new mock instance
func NewMockVirtualNetworksClient(ctrl *gomock.Controller) *MockVirtualNetworksClient {
	mock := &MockVirtualNetworksClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualNetworksClient) EXPECT() *MockVirtualNetworksClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockVirtualNetworksClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockVirtualNetworksClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualNetworksClient)(nil).Client))
}

// CreateOrUpdate mocks base method
func (m *MockVirtualNetworksClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 network.VirtualNetwork) (network.VirtualNetworksCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetworksCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockVirtualNetworksClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualNetworksClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockVirtualNetworksClient) Delete(arg0 context.Context, arg1, arg2 string) (network.VirtualNetworksDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.VirtualNetworksDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualNetworksClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualNetworksClient)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockVirtualNetworksClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVirtualNetworksClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualNetworksClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockVirtualNetworksClient) List(arg0 context.Context, arg1 string) (network.VirtualNetworkListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(network.VirtualNetworkListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualNetworksClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualNetworksClient)(nil).List), arg0, arg1)
}

// MockVirtualNetworksPeeringsClient is a mock of VirtualNetworksPeeringsClient interface
type MockVirtualNetworksPeeringsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworksPeeringsClientMockRecorder
}

// MockVirtualNetworksPeeringsClientMockRecorder is the mock recorder for MockVirtualNetworksPeeringsClient
type MockVirtualNetworksPeeringsClientMockRecorder struct {
	mock *MockVirtualNetworksPeeringsClient
}

// NewMockVirtualNetworksPeeringsClient creates a new mock instance
func NewMockVirtualNetworksPeeringsClient(ctrl *gomock.Controller) *MockVirtualNetworksPeeringsClient {
	mock := &MockVirtualNetworksPeeringsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworksPeeringsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualNetworksPeeringsClient) EXPECT() *MockVirtualNetworksPeeringsClientMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockVirtualNetworksPeeringsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockVirtualNetworksPeeringsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualNetworksPeeringsClient)(nil).Client))
}

// Delete mocks base method
func (m *MockVirtualNetworksPeeringsClient) Delete(arg0 context.Context, arg1, arg2, arg3 string) (network.VirtualNetworkPeeringsDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetworkPeeringsDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualNetworksPeeringsClientMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualNetworksPeeringsClient)(nil).Delete), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockVirtualNetworksPeeringsClient) List(arg0 context.Context, arg1, arg2 string) (network.VirtualNetworkPeeringListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.VirtualNetworkPeeringListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualNetworksPeeringsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualNetworksPeeringsClient)(nil).List), arg0, arg1, arg2)
}

// MockPublicIPAddressesClient is a mock of PublicIPAddressesClient interface
type MockPublicIPAddressesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPublicIPAddressesClientMockRecorder
}

// MockPublicIPAddressesClientMockRecorder is the mock recorder for MockPublicIPAddressesClient
type MockPublicIPAddressesClientMockRecorder struct {
	mock *MockPublicIPAddressesClient
}

// NewMockPublicIPAddressesClient creates a new mock instance
func NewMockPublicIPAddressesClient(ctrl *gomock.Controller) *MockPublicIPAddressesClient {
	mock := &MockPublicIPAddressesClient{ctrl: ctrl}
	mock.recorder = &MockPublicIPAddressesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublicIPAddressesClient) EXPECT() *MockPublicIPAddressesClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPublicIPAddressesClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.PublicIPAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.PublicIPAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPublicIPAddressesClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// ListVirtualMachineScaleSetPublicIPAddressesComplete mocks base method
func (m *MockPublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesComplete(arg0 context.Context, arg1, arg2 string) (network.PublicIPAddressListResultIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVirtualMachineScaleSetPublicIPAddressesComplete", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.PublicIPAddressListResultIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualMachineScaleSetPublicIPAddressesComplete indicates an expected call of ListVirtualMachineScaleSetPublicIPAddressesComplete
func (mr *MockPublicIPAddressesClientMockRecorder) ListVirtualMachineScaleSetPublicIPAddressesComplete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualMachineScaleSetPublicIPAddressesComplete", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).ListVirtualMachineScaleSetPublicIPAddressesComplete), arg0, arg1, arg2)
}
