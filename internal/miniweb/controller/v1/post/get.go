// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/bugsmo/miniweb/internal/pkg/core"
	"github.com/bugsmo/miniweb/internal/pkg/known"
	"github.com/bugsmo/miniweb/internal/pkg/log"
)

// Get 获取指定的博客.
func (ctrl *PostController) Get(c *gin.Context) {
	log.C(c).Infow("Get post function called")

	post, err := ctrl.b.Posts().Get(c, c.GetString(known.XUsernameKey), c.Param("postID"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
