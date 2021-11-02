/*
 * @Author: your name
 * @Date: 2020-12-19 13:45:03
 * @LastEditTime: 2020-12-19 13:45:56
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/config/jwt.go
 */
package config

import "time"

type JwtConfig struct {
	SECRET string
	EXP    time.Duration // 过期时间
	ALG    string        // 算法
}

func GetJwtConfig() *JwtConfig {
	return &JwtConfig{
		SECRET: GetEnv().AppSecret,
		EXP:    time.Hour,
		ALG:    "HS256",
	}
}
