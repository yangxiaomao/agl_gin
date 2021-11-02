/*
 * @Author: your name
 * @Date: 2020-12-19 13:42:31
 * @LastEditTime: 2020-12-23 13:26:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/config/config.go
 */
package config

import "github.com/go-sql-driver/mysql"

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	Debug bool

	Database     mysql.Config
	DbPrefix     string
	MaxIdleConns int
	MaxOpenConns int
	ServerPort   string

	RedisIp        string
	RedisPort      string
	RedisPassword  string
	RedisDb        int
	RedisSessionDb int
	RedisCacheDb   int

	SessionKey    string
	SessionSecret string

	AppSecret string

	AccessLog     bool
	AccessLogPath string
	ErrorLog      bool
	ErrorLogPath  string
	InfoLog       bool
	InfoLogPath   string

	SqlLog bool

	TemplatePath string // 静态文件相对路径
}

func GetEnv() *Env {
	return &env
}
