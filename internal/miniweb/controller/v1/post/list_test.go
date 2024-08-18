// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

package post

import (
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/bugsmo/miniweb/internal/miniweb/biz"
)

func TestPostController_List(t *testing.T) {
	type fields struct {
		b biz.IBiz
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &PostController{
				b: tt.fields.b,
			}
			ctrl.List(tt.args.c)
		})
	}
}
