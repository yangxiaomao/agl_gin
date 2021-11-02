/*
 * @Author: your name
 * @Date: 2020-12-19 14:14:33
 * @LastEditTime: 2020-12-22 15:44:32
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/controllers/MainController.go
 */
package controllers

import (
	"encoding/base64"
	"encoding/json"
	"gin/filters/auth"
	"gin/filters/auth/drivers"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
)

type RsaSignJson struct {
	Sign string
}

/**
 * @Author: yxm
 * @Date: 2020-12-21
 * @param {*}	 json串
 * @Description  Rsa加密
 * @return {*} json
 */

func RsaEncrypt(c *gin.Context) {
	authDr, _ := c.MustGet("rsa_auth").(auth.Auth)
	userId := c.Param("userid")
	sign, _ := authDr.Login(c.Request, c.Writer, map[string]interface{}{"id": userId}).(string)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"sign": sign,
		},
	})

}

/**
 * @Author: yxm
 * @Date: 2020-12-21
 * @param {*}	 json串
 * @Description  Rsa解密
 * @return {*} json
 */

func RsaDecrypt(c *gin.Context) {
	var signData RsaSignJson
	beego.Info(signData)
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	jsonDatabytes := c.Request.Header.Get("sign")

	signBytes, err := base64.StdEncoding.DecodeString(jsonDatabytes)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "Rsa解密失败!,失败原因：" + err.Error(),
			"data": gin.H{},
		})
	}
	origData, _ := drivers.AuthRsaDecrypt(signBytes)
	sign_map := make(map[string]string)
	err = json.Unmarshal(origData, &sign_map)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "Rsa解密失败!,失败原因：" + err.Error(),
			"data": gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"user_info": sign_map,
		},
	})

}
