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

// Delete 删除指定的博客.
func (ctrl *PostController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete post function called")

	if err := ctrl.b.Posts().Delete(c, c.GetString(known.XUsernameKey), c.Param("postID")); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
