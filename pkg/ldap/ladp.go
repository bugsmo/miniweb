// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

package db

import (
	"github.com/go-ldap/ldap/v3"
)

// MySQLOptions 定义 MySQL 数据库的选项.
type LdapOptions struct {
	Host      string // ldap://ldap.example.com:389
	AdminUser string
	BaseDN    string
	Password  string
}

// NewMySQL 使用给定的选项创建一个新的 gorm 数据库实例.
func NewLdap(opts *LdapOptions) (*ldap.Conn, error) {
	l, err := ldap.DialURL(opts.Host)

	return l, err
}
