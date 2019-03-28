// Code generated by MockGen. DO NOT EDIT.
// Source: sync.go

// Package mock_sync is a generated GoMock package.
package mock_sync

import (
	context "context"
	reflect "reflect"

	sync "github.com/cloudwan/gohan/sync"
	gomock "github.com/golang/mock/gomock"
)

// MockSync is a mock of Sync interface
type MockSync struct {
	ctrl     *gomock.Controller
	recorder *MockSyncMockRecorder
}

// MockSyncMockRecorder is the mock recorder for MockSync
type MockSyncMockRecorder struct {
	mock *MockSync
}

// NewMockSync creates a new mock instance
func NewMockSync(ctrl *gomock.Controller) *MockSync {
	mock := &MockSync{ctrl: ctrl}
	mock.recorder = &MockSyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSync) EXPECT() *MockSyncMockRecorder {
	return m.recorder
}

// HasLock mocks base method
func (m *MockSync) HasLock(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasLock", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasLock indicates an expected call of HasLock
func (mr *MockSyncMockRecorder) HasLock(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasLock", reflect.TypeOf((*MockSync)(nil).HasLock), path)
}

// Lock mocks base method
func (m *MockSync) Lock(ctx context.Context, path string, block bool) (chan struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, path, block)
	ret0, _ := ret[0].(chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock
func (mr *MockSyncMockRecorder) Lock(ctx, path, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockSync)(nil).Lock), ctx, path, block)
}

// Unlock mocks base method
func (m *MockSync) Unlock(ctx context.Context, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", ctx, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock
func (mr *MockSyncMockRecorder) Unlock(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockSync)(nil).Unlock), ctx, path)
}

// Fetch mocks base method
func (m *MockSync) Fetch(ctx context.Context, path string) (*sync.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, path)
	ret0, _ := ret[0].(*sync.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockSyncMockRecorder) Fetch(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockSync)(nil).Fetch), ctx, path)
}

// Update mocks base method
func (m *MockSync) Update(ctx context.Context, path, json string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, path, json)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSyncMockRecorder) Update(ctx, path, json interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSync)(nil).Update), ctx, path, json)
}

// Delete mocks base method
func (m *MockSync) Delete(ctx context.Context, path string, prefix bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, path, prefix)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSyncMockRecorder) Delete(ctx, path, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSync)(nil).Delete), ctx, path, prefix)
}

// Watch mocks base method
func (m *MockSync) Watch(ctx context.Context, path string, revision int64) <-chan *sync.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", ctx, path, revision)
	ret0, _ := ret[0].(<-chan *sync.Event)
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockSyncMockRecorder) Watch(ctx, path, revision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockSync)(nil).Watch), ctx, path, revision)
}

// Compact mocks base method
func (m *MockSync) Compact(ctx context.Context, revision int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compact", ctx, revision)
	ret0, _ := ret[0].(error)
	return ret0
}

// Compact indicates an expected call of Compact
func (mr *MockSyncMockRecorder) Compact(ctx, revision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compact", reflect.TypeOf((*MockSync)(nil).Compact), ctx, revision)
}

// CompareAndSwap mocks base method
func (m *MockSync) CompareAndSwap(ctx context.Context, path, data string, condition ...sync.CASCondition) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path, data}
	for _, a := range condition {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompareAndSwap", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompareAndSwap indicates an expected call of CompareAndSwap
func (mr *MockSyncMockRecorder) CompareAndSwap(ctx, path, data interface{}, condition ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path, data}, condition...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareAndSwap", reflect.TypeOf((*MockSync)(nil).CompareAndSwap), varargs...)
}

// ByValue mocks base method
func (m *MockSync) ByValue(value string) sync.CASCondition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByValue", value)
	ret0, _ := ret[0].(sync.CASCondition)
	return ret0
}

// ByValue indicates an expected call of ByValue
func (mr *MockSyncMockRecorder) ByValue(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByValue", reflect.TypeOf((*MockSync)(nil).ByValue), value)
}

// GetProcessID mocks base method
func (m *MockSync) GetProcessID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetProcessID indicates an expected call of GetProcessID
func (mr *MockSyncMockRecorder) GetProcessID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessID", reflect.TypeOf((*MockSync)(nil).GetProcessID))
}

// Close mocks base method
func (m *MockSync) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSyncMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSync)(nil).Close))
}

// MockCASCondition is a mock of CASCondition interface
type MockCASCondition struct {
	ctrl     *gomock.Controller
	recorder *MockCASConditionMockRecorder
}

// MockCASConditionMockRecorder is the mock recorder for MockCASCondition
type MockCASConditionMockRecorder struct {
	mock *MockCASCondition
}

// NewMockCASCondition creates a new mock instance
func NewMockCASCondition(ctrl *gomock.Controller) *MockCASCondition {
	mock := &MockCASCondition{ctrl: ctrl}
	mock.recorder = &MockCASConditionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCASCondition) EXPECT() *MockCASConditionMockRecorder {
	return m.recorder
}
