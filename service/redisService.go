/*
 * @Author: your name
 * @Date: 2020-12-21 14:34:14
 * @LastEditTime: 2020-12-22 11:00:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/service/redisService.go
 */
package service

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool //创建redis连接池

//获取Redis连接池
func newRedisPool(server, password string, db int8) (*redis.Pool, error) {
	var err error
	return &redis.Pool{
		MaxIdle:     32,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			c, err = redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err = c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, err
}

/*
 获取redis数据库连接
*/
func GetRedisConnection(redis_db int8) (redis.Conn, error) {
	if redisPool == nil {
		var err error
		redisPool, err = newRedisPool("127.0.0.1:6379", "", redis_db)
		if err != nil {
			return nil, err
		}
	}

	return redisPool.Get(), nil
}
