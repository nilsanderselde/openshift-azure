// Code generated by MockGen. DO NOT EDIT.
// Source: hash.go

// Package mock_cluster is a generated GoMock package.
package mock_cluster

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/openshift/openshift-azure/pkg/api"
)

// MockHasher is a mock of Hasher interface
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// HashWorkerScaleSet mocks base method
func (m *MockHasher) HashWorkerScaleSet(arg0 *api.OpenShiftManagedCluster, arg1 *api.AgentPoolProfile) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashWorkerScaleSet", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashWorkerScaleSet indicates an expected call of HashWorkerScaleSet
func (mr *MockHasherMockRecorder) HashWorkerScaleSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashWorkerScaleSet", reflect.TypeOf((*MockHasher)(nil).HashWorkerScaleSet), arg0, arg1)
}

// HashMasterScaleSet mocks base method
func (m *MockHasher) HashMasterScaleSet(cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile, instanceID int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashMasterScaleSet", cs, app, instanceID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashMasterScaleSet indicates an expected call of HashMasterScaleSet
func (mr *MockHasherMockRecorder) HashMasterScaleSet(cs, app, instanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashMasterScaleSet", reflect.TypeOf((*MockHasher)(nil).HashMasterScaleSet), cs, app, instanceID)
}

// HashSyncPod mocks base method
func (m *MockHasher) HashSyncPod(cs *api.OpenShiftManagedCluster) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashSyncPod", cs)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashSyncPod indicates an expected call of HashSyncPod
func (mr *MockHasherMockRecorder) HashSyncPod(cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashSyncPod", reflect.TypeOf((*MockHasher)(nil).HashSyncPod), cs)
}
