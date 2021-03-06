// Code generated by MockGen. DO NOT EDIT.
// Source: database/user_query.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/zakiry/golang_api_sample/entity"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// GetById mocks base method
func (m *MockUser) GetById(id int64) entity.User {
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(entity.User)
	return ret0
}

// GetById indicates an expected call of GetById
func (mr *MockUserMockRecorder) GetById(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUser)(nil).GetById), id)
}
