// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

package user

import (
	"github.com/bugsmo/miniweb/internal/miniweb/biz"
	"github.com/bugsmo/miniweb/internal/miniweb/store"
	"github.com/bugsmo/miniweb/pkg/auth"
	pb "github.com/bugsmo/miniweb/pkg/proto/miniweb/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	a *auth.Authz
	b biz.IBiz
	pb.UnimplementedMiniWebServer
}

// New 创建一个 user controller.
func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.NewBiz(ds)}
}
