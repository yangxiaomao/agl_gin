/*
 * @Author: your name
 * @Date: 2020-12-19 14:23:40
 * @LastEditTime: 2020-12-19 14:23:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/filters/auth/drivers/cookie.go
 */
package drivers

import (
	"gin/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(config.GetEnv().AppSecret))

type cookieAuthManager struct {
	name string
}

func NewCookieAuthDriver() *cookieAuthManager {
	return &cookieAuthManager{
		name: config.GetCookieConfig().NAME,
	}
}

func (cookie *cookieAuthManager) Check(c *gin.Context) bool {
	// read cookie
	session, err := store.Get(c.Request, cookie.name)
	if err != nil {
		return false
	}
	if session == nil {
		return false
	}
	if session.Values == nil {
		return false
	}
	if session.Values["id"] == nil {
		return false
	}
	return true
}

func (cookie *cookieAuthManager) User(c *gin.Context) interface{} {
	// get model user
	session, err := store.Get(c.Request, cookie.name)
	if session == nil {
		panic("wrong session")
	}
	if err != nil {
		return session.Values
	}
	return session.Values
}

func (cookie *cookieAuthManager) Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {
	// write cookie
	session, err := store.Get(http, cookie.name)
	if err != nil {
		return false
	}
	session.Values["id"] = user["id"]
	_ = session.Save(http, w)
	return true
}

func (cookie *cookieAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool {
	// del cookie
	session, err := store.Get(http, cookie.name)
	if err != nil {
		return false
	}
	session.Values["id"] = nil
	_ = session.Save(http, w)
	return true
}
