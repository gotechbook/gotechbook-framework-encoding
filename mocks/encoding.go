package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEncoding is a mock of Encoding interface.
type MockEncoding struct {
	ctrl     *gomock.Controller
	recorder *MockEncodingMockRecorder
}

// MockEncodingMockRecorder is the mock recorder for MockEncoding.
type MockEncodingMockRecorder struct {
	mock *MockEncoding
}

// NewMockEncoding creates a new mock instance.
func NewMockEncoding(ctrl *gomock.Controller) *MockEncoding {
	mock := &MockEncoding{ctrl: ctrl}
	mock.recorder = &MockEncodingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoding) EXPECT() *MockEncodingMockRecorder {
	return m.recorder
}

// Marshal mocks base method.
func (m *MockEncoding) Marshal(v interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal", v)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal.
func (mr *MockEncodingMockRecorder) Marshal(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockEncoding)(nil).Marshal), v)
}

// Name mocks base method.
func (m *MockEncoding) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockEncodingMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockEncoding)(nil).Name))
}

// Unmarshal mocks base method.
func (m *MockEncoding) Unmarshal(data []byte, v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", data, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal.
func (mr *MockEncodingMockRecorder) Unmarshal(data, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockEncoding)(nil).Unmarshal), data, v)
}
