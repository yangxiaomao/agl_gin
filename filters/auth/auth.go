/*
 * @Author: your name
 * @Date: 2020-12-19 14:02:27
 * @LastEditTime: 2020-12-22 15:52:14
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/filters/auth/auth.go
 */
package auth

import (
	"gin/filters/auth/drivers"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	JwtAuthDriverKey    = "jwt"
	CookieAuthDriverKey = "cookie"
	RsaAuthDriverKey    = "rsa"
)

var driverList = map[string]Auth{
	CookieAuthDriverKey: drivers.NewCookieAuthDriver(),
	JwtAuthDriverKey:    drivers.NewJwtAuthDriver(),
	RsaAuthDriverKey:    drivers.NewRsaAuthDriver(),
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(http *http.Request, w http.ResponseWriter) bool
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, GenerateAuthDriver(authKey))
		c.Next()
	}
}

func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !GenerateAuthDriver(authKey).Check(c) {
			c.HTML(http.StatusOK, "index.tpl", gin.H{
				"title": "先登录",
			})
			c.Abort()
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) Auth {
	return driverList[string]
}

func GetCurUser(c *gin.Context, key string) map[string]interface{} {
	authDriver, _ := c.MustGet(key).(Auth)
	return authDriver.User(c).(map[string]interface{})
}

func RsaMiddleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !GenerateAuthDriver(authKey).Check(c) {
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "Authentication failed",
				"data": gin.H{},
			})
			c.Abort()
		}
		c.Next()
	}
}
