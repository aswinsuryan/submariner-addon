// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloud.go

// Package fake is a generated GoMock package.
package fake

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	events "github.com/openshift/library-go/pkg/operator/events"
	v1alpha1 "github.com/stolostron/submariner-addon/pkg/apis/submarinerconfig/v1alpha1"
	cloud "github.com/stolostron/submariner-addon/pkg/cloud"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// CleanUpSubmarinerClusterEnv mocks base method.
func (m *MockProvider) CleanUpSubmarinerClusterEnv() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUpSubmarinerClusterEnv")
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUpSubmarinerClusterEnv indicates an expected call of CleanUpSubmarinerClusterEnv.
func (mr *MockProviderMockRecorder) CleanUpSubmarinerClusterEnv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUpSubmarinerClusterEnv", reflect.TypeOf((*MockProvider)(nil).CleanUpSubmarinerClusterEnv))
}

// PrepareSubmarinerClusterEnv mocks base method.
func (m *MockProvider) PrepareSubmarinerClusterEnv() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareSubmarinerClusterEnv")
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareSubmarinerClusterEnv indicates an expected call of PrepareSubmarinerClusterEnv.
func (mr *MockProviderMockRecorder) PrepareSubmarinerClusterEnv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareSubmarinerClusterEnv", reflect.TypeOf((*MockProvider)(nil).PrepareSubmarinerClusterEnv))
}

// MockProviderFactory is a mock of ProviderFactory interface.
type MockProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockProviderFactoryMockRecorder
}

// MockProviderFactoryMockRecorder is the mock recorder for MockProviderFactory.
type MockProviderFactoryMockRecorder struct {
	mock *MockProviderFactory
}

// NewMockProviderFactory creates a new mock instance.
func NewMockProviderFactory(ctrl *gomock.Controller) *MockProviderFactory {
	mock := &MockProviderFactory{ctrl: ctrl}
	mock.recorder = &MockProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderFactory) EXPECT() *MockProviderFactoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProviderFactory) Get(managedClusterInfo *v1alpha1.ManagedClusterInfo, config *v1alpha1.SubmarinerConfig, eventsRecorder events.Recorder) (cloud.Provider, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", managedClusterInfo, config, eventsRecorder)
	ret0, _ := ret[0].(cloud.Provider)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockProviderFactoryMockRecorder) Get(managedClusterInfo, config, eventsRecorder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProviderFactory)(nil).Get), managedClusterInfo, config, eventsRecorder)
}
