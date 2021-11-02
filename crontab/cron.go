/*
 * @Author: yxm
 * @Date: 2020-11-17 01:56:01
 * @LastEditTime: 2020-12-26 13:36:56
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/crontab/cron.go
 */
package crontab

import (
	"fmt"

	"github.com/robfig/cron"
)

//	 second minute hour day month week   command
//顺序：秒      分    时   日   月    周      命令
func Init() {
	fmt.Print("Starting...")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		fmt.Print("Run frist...")
	})

	c.AddFunc("* * * * * *", func() {
		fmt.Print("Run two...")
	})
	c.Start()

}

func batchCreateMd5Int() error {

	return nil
}
