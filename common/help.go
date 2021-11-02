/*
 * @Author: your name
 * @Date: 2020-12-23 11:16:48
 * @LastEditTime: 2020-12-24 13:17:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/common/help.go
 */
package common

import (
	"crypto/md5"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
)

func ResponJson(c *gin.Context, msg string, data interface{}, errorCode interface{}) {
	c.JSON(200, gin.H{"msg": msg, "data": data, "errorCode": errorCode})
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
