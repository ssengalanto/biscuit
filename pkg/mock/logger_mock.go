package mock

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

// LoggerRecorder is the mock recorder for MockLogger.
type LoggerRecorder struct {
	mock *Logger
}

// Logger is mock of interfaces.Logger interface.
type Logger struct {
	ctrl     *gomock.Controller
	recorder *LoggerRecorder
}

// NewLogger creates a new MockLogger instance.
func NewLogger(ctrl *gomock.Controller) *Logger {
	mock := &Logger{ctrl: ctrl}
	mock.recorder = &LoggerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (l *Logger) EXPECT() *LoggerRecorder {
	return l.recorder
}

// Info mocks base method.
func (l *Logger) Info(msg string, fields interfaces.Fields) {
	l.ctrl.T.Helper()
	l.ctrl.Call(l, "Info", msg, fields)
}

// Info indicates an expected call of Info.
func (lr *LoggerRecorder) Info(msg, fields interface{}) *gomock.Call {
	lr.mock.ctrl.T.Helper()
	return lr.mock.ctrl.RecordCallWithMethodType(lr.mock, "Info", reflect.TypeOf((*Logger)(nil).Info), msg, fields)
}

// Error mocks base method.
func (l *Logger) Error(msg string, fields interfaces.Fields) {
	l.ctrl.T.Helper()
	l.ctrl.Call(l, "Error", msg, fields)
}

// Error indicates an expected call of Error.
func (lr *LoggerRecorder) Error(msg, fields interface{}) *gomock.Call {
	lr.mock.ctrl.T.Helper()
	return lr.mock.ctrl.RecordCallWithMethodType(lr.mock, "Error", reflect.TypeOf((*Logger)(nil).Error), msg, fields)
}

// Debug mocks base method.
func (l *Logger) Debug(msg string, fields interfaces.Fields) {
	l.ctrl.T.Helper()
	l.ctrl.Call(l, "Debug", msg, fields)
}

// Debug indicates an expected call of Debug.
func (lr *LoggerRecorder) Debug(msg, fields interface{}) *gomock.Call {
	lr.mock.ctrl.T.Helper()
	return lr.mock.ctrl.RecordCallWithMethodType(lr.mock, "Debug", reflect.TypeOf((*Logger)(nil).Debug), msg, fields)
}

// Warn mocks base method.
func (l *Logger) Warn(msg string, fields interfaces.Fields) {
	l.ctrl.T.Helper()
	l.ctrl.Call(l, "Warn", msg, fields)
}

// Warn indicates an expected call of Warn.
func (lr *LoggerRecorder) Warn(msg, fields interface{}) *gomock.Call {
	lr.mock.ctrl.T.Helper()
	return lr.mock.ctrl.RecordCallWithMethodType(lr.mock, "Warn", reflect.TypeOf((*Logger)(nil).Warn), msg, fields)
}

// Fatal mocks base method.
func (l *Logger) Fatal(msg string, fields interfaces.Fields) {
	l.ctrl.T.Helper()
	l.ctrl.Call(l, "Warn", msg, fields)
}

// Fatal indicates an expected call of Fatal.
func (lr *LoggerRecorder) Fatal(msg, fields interface{}) *gomock.Call {
	lr.mock.ctrl.T.Helper()
	return lr.mock.ctrl.RecordCallWithMethodType(lr.mock, "Fatal", reflect.TypeOf((*Logger)(nil).Fatal), msg, fields)
}

// Panic mocks base method.
func (l *Logger) Panic(msg string, fields interfaces.Fields) {
	l.ctrl.T.Helper()
	l.ctrl.Call(l, "Panic", msg, fields)
}

// Panic indicates an expected call of Panic.
func (lr *LoggerRecorder) Panic(msg, fields interface{}) *gomock.Call {
	lr.mock.ctrl.T.Helper()
	return lr.mock.ctrl.RecordCallWithMethodType(lr.mock, "Panic", reflect.TypeOf((*Logger)(nil).Panic), msg, fields)
}
