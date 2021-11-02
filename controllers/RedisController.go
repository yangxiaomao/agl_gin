/*
 * @Author: your name
 * @Date: 2020-12-19 14:14:33
 * @LastEditTime: 2020-12-22 11:31:17
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/controllers/MainController.go
 */
package controllers

import (
	"encoding/json"
	"fmt"
	db "gin/connections/database/mysql"
	"gin/service"
	"log"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户基本信息入redis
 * @param {*gin.Context} c
 * @return {*} json
 * @Date		2020-12-21
 * @Author		yxm
 */
func UserInfoSetRedis(c *gin.Context) {
	id := c.PostForm("userid")
	log.Println("id: " + id)

	rs := db.Query("select name,avatar,id from users where id = ?", id)

	redisPool, err := service.GetRedisConnection(2)

	defer redisPool.Close() //函数运行结束 ，把连接放回连接池
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	jsonDataSeq, err := json.Marshal(rs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}
	redisKey := "gin:userinfo:uuid:" + id

	_, error := redisPool.Do("Set", redisKey, jsonDataSeq)
	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  error.Error(),
			"data": gin.H{},
		})
		return
	}
	redisPool.Close() //关闭连接池

	// 返回html
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"user": rs,
		},
	})
}

/**
 * @description: redis的基本操作
 * @param {*gin.Context} c
 * @return {*} json
 * @Date		2020-12-21
 * @Author		yxm
 */
func RedisTest(c *gin.Context) {
	redisPool, err := service.GetRedisConnection(3)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	defer redisPool.Close() //函数运行结束 ，把连接放回连接池

	id := c.PostForm("userid")
	redisKey := "gin:userinfo:uuid:" + id

	optType := c.PostForm("type")
	var conerr error
	switch optType {
	case "1":
		//正常字符串添加
		_, conerr = redisPool.Do("Set", redisKey, "hahahah")
	case "2":
		//设置过期时间
		_, conerr = redisPool.Do("Set", redisKey, "hahahah", "EX", 360)
	case "3":
		//设置list类型数据
		_, conerr = redisPool.Do("LPUSH", redisKey, "hahah", "hahahhahahah")
	case "4":
		//设置HSET类型数据
		_, conerr = redisPool.Do("HSET", redisKey, "name", "wd", "age", 22)
	case "5":
		go Subs()
		go Push("this is wd")
		time.Sleep(time.Second * 3)
	default:
		_, conerr = redisPool.Do("Set", redisKey, "hahahah")
	}
	redisPool.Close() //关闭连接池
	if conerr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
			"data": gin.H{},
		})
		return
	}

	// 返回html
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"userid": id,
		},
	})
}

func Subs() { //订阅者
	redisPool, err := service.GetRedisConnection(3)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer redisPool.Close()
	psc := redis.PubSubConn{redisPool}
	psc.Subscribe("channel1") //订阅channel1频道
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func Push(message string) { //发布者
	redisPool, _ := service.GetRedisConnection(3)
	_, err1 := redisPool.Do("PUBLISH", "channel1", message)
	if err1 != nil {
		fmt.Println("pub err: ", err1)
		return
	}

}
