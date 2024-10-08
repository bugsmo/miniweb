// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

package miniweb

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/bugsmo/miniweb/internal/miniweb/controller/v1/post"
	"github.com/bugsmo/miniweb/internal/miniweb/controller/v1/user"
	"github.com/bugsmo/miniweb/internal/miniweb/store"
	"github.com/bugsmo/miniweb/internal/pkg/core"
	"github.com/bugsmo/miniweb/internal/pkg/errno"
	"github.com/bugsmo/miniweb/internal/pkg/log"
	mw "github.com/bugsmo/miniweb/internal/pkg/middleware"
	"github.com/bugsmo/miniweb/pkg/auth"
)

// installRouters 安装 miniweb 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// 注册 pprof 路由
	pprof.Register(g)

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)
	pc := post.New(store.S)

	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)                             // 创建用户
			userv1.PUT(":name/change-password", uc.ChangePassword) // 修改用户密码
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name", uc.Get)       // 获取用户详情
			userv1.PUT(":name", uc.Update)    // 更新用户
			userv1.GET("", uc.List)           // 列出用户列表，只有 root 用户才能访问
			userv1.DELETE(":name", uc.Delete) // 删除用户
		}

		// 创建 posts 路由分组
		postv1 := v1.Group("/posts", mw.Authn())
		{
			postv1.POST("", pc.Create)             // 创建博客
			postv1.GET(":postID", pc.Get)          // 获取博客详情
			postv1.PUT(":postID", pc.Update)       // 更新用户
			postv1.DELETE("", pc.DeleteCollection) // 批量删除博客
			postv1.GET("", pc.List)                // 获取博客列表
			postv1.DELETE(":postID", pc.Delete)    // 删除博客
		}
	}

	return nil
}
