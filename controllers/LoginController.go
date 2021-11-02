/*
 * @Author: your name
 * @Date: 2020-12-19 14:14:33
 * @LastEditTime: 2020-12-28 17:29:59
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/controllers/MainController.go
 */
package controllers

import (
	"bufio"
	"fmt"
	"gin/common"
	"gin/models"
	"gin/modules/log"
	"gin/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 	用户注册接口
 * @param {*}	{"user":"13126997216","password":"123456"}
 * @return {*}
 */

func UserRegister(c *gin.Context) {

	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, userinfo := service.UserRegisterService(c, json)
	if err != nil {
		common.ResponJson(c, err.Error(), "", -1)
		return
	}
	common.ResponJson(c, "用户创建成功", userinfo, 0)
}

/**
 * @description: 	用户登录接口
 * @param {*}	{"user":"13126997216","password":"123456"}
 * @return {*}
 */

func UserLogin(c *gin.Context) {
	optType := c.PostForm("optType")
	switch optType {
	case "1":
		test1()
		break
	}

	common.ResponJson(c, "用户登录成功", gin.H{}, 0)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func test1() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/**
 * @description: 	优雅关停测试接口
 * @param {*}
 * @return {*}
 */

func GraceShutdown(c *gin.Context) {
	time.Sleep(10 * time.Second)
	var logmsg log.E
	sign_map := make(map[string]interface{})
	sign_map["info_log"] = "优雅关停成功"
	logmsg.Function = "UserRegisterService"
	logmsg.Title = "错误日志1"
	logmsg.Level = "info"
	logmsg.Info = sign_map
	log.Info(logmsg)
	common.ResponJson(c, "优雅关停成功", gin.H{}, 0)
}
