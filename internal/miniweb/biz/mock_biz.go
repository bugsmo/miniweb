// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bugsmo/miniweb/internal/miniweb/biz (interfaces: IBiz)

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	post "github.com/bugsmo/miniweb/internal/miniweb/biz/post"
	user "github.com/bugsmo/miniweb/internal/miniweb/biz/user"
)

// MockIBiz is a mock of IBiz interface
type MockIBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIBizMockRecorder
}

// MockIBizMockRecorder is the mock recorder for MockIBiz
type MockIBizMockRecorder struct {
	mock *MockIBiz
}

// NewMockIBiz creates a new mock instance
func NewMockIBiz(ctrl *gomock.Controller) *MockIBiz {
	mock := &MockIBiz{ctrl: ctrl}
	mock.recorder = &MockIBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIBiz) EXPECT() *MockIBizMockRecorder {
	return m.recorder
}

// Posts mocks base method
func (m *MockIBiz) Posts() post.PostBiz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Posts")
	ret0, _ := ret[0].(post.PostBiz)
	return ret0
}

// Posts indicates an expected call of Posts
func (mr *MockIBizMockRecorder) Posts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Posts", reflect.TypeOf((*MockIBiz)(nil).Posts))
}

// Users mocks base method
func (m *MockIBiz) Users() user.UserBiz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(user.UserBiz)
	return ret0
}

// Users indicates an expected call of Users
func (mr *MockIBizMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockIBiz)(nil).Users))
}
