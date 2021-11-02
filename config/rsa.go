/*
 * @Author: your name
 * @Date: 2020-12-19 13:45:03
 * @LastEditTime: 2020-12-22 14:54:40
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/config/jwt.go
 */
package config

type RsaConfig struct {
	SIGN string
}

func GetRsaConfig() *RsaConfig {
	return &RsaConfig{
		SIGN: "sign",
	}
}
