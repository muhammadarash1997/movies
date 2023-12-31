// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package _interface is a generated GoMock package.
package _interface

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/muhammadarash1997/movies/model"
)

// MockReadWriter is a mock of ReadWriter interface.
type MockReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockReadWriterMockRecorder
}

// MockReadWriterMockRecorder is the mock recorder for MockReadWriter.
type MockReadWriterMockRecorder struct {
	mock *MockReadWriter
}

// NewMockReadWriter creates a new mock instance.
func NewMockReadWriter(ctrl *gomock.Controller) *MockReadWriter {
	mock := &MockReadWriter{ctrl: ctrl}
	mock.recorder = &MockReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadWriter) EXPECT() *MockReadWriterMockRecorder {
	return m.recorder
}

// CheckMovieExist mocks base method.
func (m *MockReadWriter) CheckMovieExist(arg0 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMovieExist", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckMovieExist indicates an expected call of CheckMovieExist.
func (mr *MockReadWriterMockRecorder) CheckMovieExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMovieExist", reflect.TypeOf((*MockReadWriter)(nil).CheckMovieExist), arg0)
}

// Close mocks base method.
func (m *MockReadWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReadWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadWriter)(nil).Close))
}

// DeleteMovie mocks base method.
func (m *MockReadWriter) DeleteMovie(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockReadWriterMockRecorder) DeleteMovie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockReadWriter)(nil).DeleteMovie), arg0)
}

// ReadAllMovies mocks base method.
func (m *MockReadWriter) ReadAllMovies() ([]model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllMovies")
	ret0, _ := ret[0].([]model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllMovies indicates an expected call of ReadAllMovies.
func (mr *MockReadWriterMockRecorder) ReadAllMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllMovies", reflect.TypeOf((*MockReadWriter)(nil).ReadAllMovies))
}

// ReadMovie mocks base method.
func (m *MockReadWriter) ReadMovie(arg0 int64) (model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMovie", arg0)
	ret0, _ := ret[0].(model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMovie indicates an expected call of ReadMovie.
func (mr *MockReadWriterMockRecorder) ReadMovie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMovie", reflect.TypeOf((*MockReadWriter)(nil).ReadMovie), arg0)
}

// UpdateMovie mocks base method.
func (m *MockReadWriter) UpdateMovie(arg0 model.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockReadWriterMockRecorder) UpdateMovie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockReadWriter)(nil).UpdateMovie), arg0)
}

// WriteMovie mocks base method.
func (m *MockReadWriter) WriteMovie(arg0 model.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMovie", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMovie indicates an expected call of WriteMovie.
func (mr *MockReadWriterMockRecorder) WriteMovie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMovie", reflect.TypeOf((*MockReadWriter)(nil).WriteMovie), arg0)
}
