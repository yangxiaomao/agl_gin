/*
 * @Author: your name
 * @Date: 2020-12-23
 * @LastEditTime: 2020-12-24 14:26:37
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/service/redisService.go
 */
package service

import (
	"errors"
	"gin/common"
	"gin/models"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

/**
 * @description: 用户注册服务
 * @param {*models.User} yxm --- 2020-10-29
 * @return {*} response map[string]interface{}, err error
 */
func UserRegisterService(c *gin.Context, json models.User) (err error, response models.User) {

	u1 := uuid.NewV4().String()
	useruuid := common.Md5(u1)
	json.Uuid = useruuid
	json.Password = common.Md5(json.Password)
	json.Email = "353125014@qq.com"
	json.LastIp = common.GetLocalIP()
	json.LastTime = time.Now().Unix()
	tx := models.Model.Begin()
	//延迟调用处理事务
	defer func() {
		recovered := recover()
		if recovered != nil {
			tx.Rollback()

		}
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	var userinfo models.User
	// 事务处理过程
	if err, userinfo = models.CreateUser(c, json, tx); err != nil {
		return err, json
	}
	err = errors.New("事务异常")
	//自定义日志写入
	// sign_map := make(map[string]interface{})
	// sign_map["sign"] = "123412423"
	// sign_map["timestep"] = 123412312
	// var logmsg log.E
	// logmsg.Function = "UserRegisterService"
	// logmsg.Error = err
	// logmsg.Title = "错误日志1"
	// logmsg.Level = "info"
	// logmsg.Info = sign_map
	// log.Info(logmsg)

	response = userinfo
	return nil, response
}
