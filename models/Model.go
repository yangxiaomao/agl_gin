/*
 * @Author: your name
 * @Date: 2020-12-19 14:04:47
 * @LastEditTime: 2020-12-23 13:26:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/modules/Model.go
 */
package models

import (
	"gin/config"
	"gin/modules/log"

	"github.com/jinzhu/gorm"
)

var Model *gorm.DB

func init() {
	var err error
	log.Println(config.GetEnv().Database.FormatDSN())
	Model, err = gorm.Open("mysql", config.GetEnv().Database.FormatDSN())

	if err != nil {
		panic(err)
	}
}

//返回带前缀的表名
func PrefixTableName(str string) string {
	return config.GetEnv().DbPrefix + str
}
