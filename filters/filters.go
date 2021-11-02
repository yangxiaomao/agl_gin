/*
 * @Author: your name
 * @Date: 2020-12-19 14:01:27
 * @LastEditTime: 2020-12-19 14:01:38
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/filters/filters.go
 */
package filters

import (
	"gin/config"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterSession() gin.HandlerFunc {
	store, _ := sessions.NewRedisStore(
		10,
		"tcp",
		config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort,
		config.GetEnv().RedisPassword,
		[]byte(config.GetEnv().SessionSecret))
	return sessions.Sessions(config.GetEnv().SessionKey, store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort, config.GetEnv().RedisPassword, time.Minute)
	return cache.Cache(&cacheStore)
}
