/*
 * @Author: your name
 * @Date: 2020-12-19 14:14:33
 * @LastEditTime: 2020-12-23 11:49:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/controllers/MainController.go
 */
package controllers

import (
	"gin/filters/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserPersonalCenter(c *gin.Context) {

}

func GetUserInfo(c *gin.Context) {
	authDr, _ := c.MustGet("rsa_auth").(auth.Auth)

	info := authDr.User(c).(map[string]interface{})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"id": info["id"],
		},
	})
}
