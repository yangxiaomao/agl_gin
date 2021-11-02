/*
 * @Author: your name
 * @Date: 2020-12-19 14:13:29
 * @LastEditTime: 2021-01-02 13:45:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/routes/api_router.go
 */
package routes

import (
	"gin/controllers"
	"gin/filters/auth"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/test/index", controllers.IndexApi)
	}

	api := router.Group("/web")
	api.GET("/index", controllers.IndexApi)

	api.GET("/cookie/set/:userid", controllers.CookieSetExample)

	// cookie身份验证中间件
	api.Use(auth.Middleware(auth.CookieAuthDriverKey))
	{
		api.GET("/orm", controllers.OrmExample)
		api.GET("/store", controllers.StoreExample)
		api.GET("/db", controllers.DBExample)
		api.GET("/cookie/get", controllers.CookieGetExample)
	}

	jwtApi := router.Group("/api")
	jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)
	// rsa解密
	jwtApi.POST("/rsa/set/decrypt", controllers.RsaDecrypt)
	// 用户基本信息入redis
	jwtApi.POST("/redis/set/user_info", controllers.UserInfoSetRedis)
	// redis基本操作
	jwtApi.POST("/redis/set/redis_test", controllers.RedisTest)

	// jwt认证中间件
	jwtApi.Use(auth.Middleware(auth.JwtAuthDriverKey))
	{
		jwtApi.POST("/jwt/get", controllers.JwtGetExample)
	}

	rsaApi := router.Group("/rsaapi")
	// rsa加密
	rsaApi.POST("/rsa/set/encrypt", controllers.RsaEncrypt)
	rsaApi.POST("/user/register", controllers.UserRegister)
	rsaApi.POST("/user/login", controllers.UserLogin)
	// 测试优雅关停功能
	rsaApi.POST("/test/graceShutdown", controllers.GraceShutdown)
	// 测试GO并发功能
	rsaApi.POST("/test/concurrent", controllers.Concurrent)
	// 模仿微信群发红包
	rsaApi.POST("/test/wxreward", controllers.Wxreward)

	// jwt认证中间件
	rsaApi.Use(auth.RsaMiddleware(auth.RsaAuthDriverKey))
	{
		rsaApi.POST("/rsa/getuser", controllers.GetUserInfo)
	}
}
