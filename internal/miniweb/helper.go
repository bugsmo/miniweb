// Copyright 2024 Innkeeper BugsMo(莫维龙) &lt;18550039021@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/bugsmo/miniweb.

package miniweb

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bugsmo/miniweb/internal/miniweb/store"
	"github.com/bugsmo/miniweb/internal/pkg/log"
	"github.com/bugsmo/miniweb/pkg/db"
	generateLdapOpts "github.com/bugsmo/miniweb/pkg/ldap"
)

const (
	// recommendedHomeDir 定义放置 miniweb 服务配置的默认目录.
	recommendedHomeDir = ".miniweb"

	// defaultConfigName 指定了 miniweb 服务的默认配置文件名.
	defaultConfigName = "miniweb.yaml"
)

// initConfig 设置需要读取的配置文件名、环境变量，并读取配置文件内容到 viper 中.
func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)

		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath(".")

		// 设置配置文件格式为 YAML (YAML 格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为 miniweb，如果是 miniweb，将自动转变为大写。
	viper.SetEnvPrefix("miniweb")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}

	// 打印 viper 当前使用的配置文件，方便 Debug.
	log.Debugw("Using config file", "file", viper.ConfigFileUsed())
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

// initStore 读取 db 配置，创建 gorm.DB 实例，并初始化 miniweb store 层.
func initStore() error {
	dbOptions := &db.MySQLOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	dbIns, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}

	ldapOption := &generateLdapOpts.LdapOptions{
		Host:      viper.GetString("ldap.host"),
		AdminUser: viper.GetString("ldap.admin-user"),
		BaseDN:    viper.GetString("ldap.base-dn"),
		Password:  viper.GetString("ldap.password"),
	}

	ldapIns, err := generateLdapOpts.NewLdap(ldapOption)
	if err != nil {
		return err
	}

	_ = store.NewStore(dbIns, ldapIns)

	return nil
}
