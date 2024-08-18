// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

package main

import (
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/bugsmo/miniweb/internal/miniweb"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := miniweb.NewMiniWebCommand()
	if err := command.Execute(); err != nil {
		os.Exit(2)
	}
}
