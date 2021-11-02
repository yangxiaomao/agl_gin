/*
 * @Author: your name
 * @Date: 2020-12-19 13:44:08
 * @LastEditTime: 2020-12-23 13:27:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/config/env.go
 */
package config

import "github.com/go-sql-driver/mysql"

// 本文件建议在代码协同工具(git/svn等)中忽略

var env = Env{
	Debug: true,

	ServerPort: "8000",

	Database: mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Addr:                 "127.0.0.1:3306",
		DBName:               "db_gin",
		Collation:            "utf8mb4_unicode_ci",
		Net:                  "tcp",
		AllowNativePasswords: true,
	},
	DbPrefix:     "tb_",
	MaxIdleConns: 50,
	MaxOpenConns: 100,

	RedisIp:       "127.0.0.1",
	RedisPort:     "6379",
	RedisPassword: "",
	RedisDb:       0,

	RedisSessionDb: 1,
	RedisCacheDb:   2,

	AccessLog:     true,
	AccessLogPath: "storage/logs/access.log",

	ErrorLog:     true,
	ErrorLogPath: "storage/logs/error.log",

	InfoLog:     true,
	InfoLogPath: "storage/logs/info.log",

	TemplatePath: "frontend/templates",

	//APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
	AppSecret: "FFC150A160D37E92012C196B6AF4160D",
}
