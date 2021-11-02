/*
 * @Author: your name
 * @Date: 2020-12-19 13:44:35
 * @LastEditTime: 2020-12-19 15:23:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/config/cookie.go
 */
package config

type CookieConfig struct {
	NAME string
}

func GetCookieConfig() *CookieConfig {
	return &CookieConfig{
		NAME: "aglgin_session",
	}
}
