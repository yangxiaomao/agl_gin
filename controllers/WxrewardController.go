/*
 * @Author: your name
 * @Date: 2020-12-19 14:14:33
 * @LastEditTime: 2021-01-07 18:33:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/controllers/MainController.go
 */
package controllers

import (
	"fmt"
	"gin/common"
	"gin/modules/log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//定义红包结构体
type Reward struct {
	Count          int   //个数
	Money          int   //总金额（分）
	RemainCount    int   //剩余个数
	RemainMoney    int   //剩余金额（分）
	BestMoney      int   //手气最佳金额
	BestMoneyIndex int   //手气最佳序号
	MoneyList      []int //拆分列表
}

/**
 * @description: 	协程测试微信发红包接口
 * @param {*}
 * @return {*}
 */

func Wxreward(c *gin.Context) {
	chanReward := make(chan Reward)
	rand.Seed(time.Now().UnixNano())
	//开启携程随机生成红包
	go func() {
		//获取群发红包数量
		count, err := strconv.Atoi(c.PostForm("rewardCount"))

		if err != nil {
			fmt.Println(err.Error())
			common.ResponJson(c, err.Error(), gin.H{}, 0)
			return
		}
		//获取群发红包金额
		money, _ := strconv.Atoi(c.PostForm("rewardMoney"))

		if err != nil {
			common.ResponJson(c, err.Error(), gin.H{}, 0)
			return
		}

		if money < count {
			fmt.Println("金额不够分！")
			common.ResponJson(c, "金额不够分！", gin.H{}, 0)
			return
		}

		avg := money / count

		for avg == 0 {
			//保证金额足够分配
			count = rand.Intn(50) + 1
			money = rand.Intn(10000) + 100
			avg = money / count

		}
		//初始化结构体并赋值
		reward := Reward{Count: count, Money: money,
			RemainCount: count, RemainMoney: money}
		//结构体入管道

		chanReward <- reward
		//关闭管道
		close(chanReward)
	}()

	student := [10]string{"张三", "李四", "王五", "小明", "小红", "小蓝", "小黑", "小马", "小何", "小孟"}

	// 打印拆包列表， 带手气最佳
	for reward := range chanReward {
		for i := 0; reward.RemainCount > 0; i++ {
			money := GrabReward(&reward)
			if money > reward.BestMoney {
				reward.BestMoneyIndex, reward.BestMoney = i, money
			}
			reward.MoneyList = append(reward.MoneyList, money)
		}

		fmt.Printf("红包总个数：%d, 红包总金额：%.2f\n", reward.Count, float32(reward.Money)/100)
		for i := range reward.MoneyList {
			money := reward.MoneyList[i]
			isBest := ""
			if reward.BestMoneyIndex == i {
				isBest = "** 这位同学手气最佳"
			}

			fmt.Printf("红包被【%s】抢了 : (%.2f)元%s\n", student[i], float32(money)/100, isBest)
			time.Sleep(time.Second)
		}

		fmt.Printf("-------------")
	}
	common.ResponJson(c, "协程测试微信发红包", gin.H{}, 0)
}

//创建抢红包方法
func GrabReward(reward *Reward) int {
	//如果剩余红包不存
	if reward.RemainCount <= 0 {
		panic("RemmainCount <= 0")
	}
	//如果还剩最后一个红包
	if reward.RemainCount == 1 {
		money := reward.RemainMoney
		reward.RemainCount = 0
		reward.RemainMoney = 0
		return money
	}
	//是否可以直接0.01
	if (reward.RemainMoney / reward.RemainCount) == 1 {
		money := 1
		reward.RemainMoney -= money
		reward.RemainCount--
		return money
	}

	//最大可领金额 = 剩余金额的平均值X2 = （剩余金额 / 剩余数量） * 2
	//领取金额范围 = 0.01 ~ 最大可领金额
	maxMoney := int(reward.RemainMoney/reward.RemainCount) * 2
	rand.Seed(time.Now().UnixNano())
	money := rand.Intn(maxMoney)
	for money == 0 {
		//防止零
		money = rand.Intn(maxMoney)
	}
	reward.RemainMoney -= money
	//防止剩余金额为负数
	if reward.RemainMoney < 0 {
		money += reward.RemainMoney
		reward.RemainMoney = 0
		reward.RemainCount = 0

	} else {
		reward.RemainCount--
	}
	return money
}

//测试批量发红包
func BatchWxreward(c *gin.Context) {
	chanReward := make(chan Reward)
	rand.Seed(time.Now().UnixNano())
	//开启携程随机生成红包
	go func() {
		//随机生成1000个红包
		for i := 0; i < 1000; i++ {
			//随机红包个数1~50
			count := rand.Intn(50) + 1
			//随机红包总金额 1~100元
			money := rand.Intn(10000) + 100

			avg := money / count

			for avg == 0 {
				//保证金额足够分配
				count = rand.Intn(50) + 1
				money = rand.Intn(10000) + 100
				avg = money / count

			}
			//初始化结构体并赋值
			reward := Reward{Count: count, Money: money,
				RemainCount: count, RemainMoney: money}
			//结构体入管道
			chanReward <- reward
		}
		//关闭管道
		close(chanReward)
	}()

	// 打印拆包列表， 带手气最佳
	for reward := range chanReward {
		for i := 0; reward.RemainCount > 0; i++ {
			money := GrabReward(&reward)
			if money > reward.BestMoney {
				reward.BestMoneyIndex, reward.BestMoney = i, money
			}
			reward.MoneyList = append(reward.MoneyList, money)
		}
		var logmsg log.E
		totalMap := make(map[string]interface{})
		totalMap["total_num"] = reward.Count
		totalMap["total_money"] = float32(reward.Money) / 100
		logmsg.Function = "Wxreward"
		logmsg.Title = "协程测试微信发红包"
		logmsg.Level = "info"
		logmsg.Info = totalMap
		log.Info(logmsg)
		fmt.Println("总个数：%d, 总金额：%.2f", reward.Count, float32(reward.Money)/100)
		for i := range reward.MoneyList {
			money := reward.MoneyList[i]
			isBest := ""
			if reward.BestMoneyIndex == i {
				isBest = "** 手气最佳"
			}
			singleMap := make(map[string]interface{})
			singleMap["money_id"] = i + 1
			singleMap["money"] = float32(money) / 100
			singleMap["is_best"] = isBest
			logmsg.Info = singleMap
			log.Info(logmsg)
			fmt.Println("money_%d : (%.2f)%s\n", i+1, float32(money)/100, isBest)
		}
		infoMap := make(map[string]interface{})
		infoMap["info"] = "-----------------------"
		logmsg.Info = infoMap
		log.Info(logmsg)
		fmt.Println("-------------")
	}
	common.ResponJson(c, "协程测试模拟微信批量发群红包", gin.H{}, 0)
}
