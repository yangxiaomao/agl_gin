/*
 * @Author: your name
 * @Date: 2020-12-19 14:05:24
 * @LastEditTime: 2020-12-25 18:29:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/modules/User.go
 */
package models

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id         int64 ``
	Uuid       string
	Username   string
	Password   string
	Email      string
	LoginCount int64
	LastTime   int64
	LastIp     string
	State      int64
	Created    int64
	Updated    int64
}

type UserJoinInfo struct {
	Id         int64
	Uuid       string
	Username   string
	Password   string
	Email      string
	LoginCount int64
	LastTime   int64
	LastIp     string
	State      int64
	Sex        int64
	Soybean    float64
	Created    int64
	Updated    int64
}

func (User) TableName() string {

	return PrefixTableName("user")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func CreateUser(c *gin.Context, json User, tx *gorm.DB) (err error, userinfo User) {

	//查询用户是否存在
	var hasUser User
	tx.Where("username = ?", json.Username).First(&hasUser)
	if hasUser.Id != 0 {
		return errors.New("用户已存在"), hasUser
	}
	timeStr := time.Now().Unix()
	userinfo = User{Uuid: json.Uuid, Username: json.Username, Password: json.Password, Email: json.Email, LastIp: json.LastIp,
		Created: timeStr}
	result := tx.Omit("Id", "LoginCount", "LastTime", "State", "Sex", "Soybean", "Updated").Create(&userinfo)
	err = result.Error
	return
}
