/*
 * @Author: your name
 * @Date: 2020-12-22 14:46:20
 * @LastEditTime: 2020-12-26 13:39:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/main.go
 */
package main

import (
	"fmt"
	"gin/config"
	"gin/modules/server"
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	// crontab.Init()

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Print(config.GetEnv().Debug)

	if config.GetEnv().Debug {
		fmt.Print(gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := initRouter()
	fmt.Println(reflect.TypeOf(router))

	server.Run(router)
}
