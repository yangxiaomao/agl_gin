/*
 * @Author: your name
 * @Date: 2020-12-19 14:14:33
 * @LastEditTime: 2021-01-01 22:23:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/controllers/MainController.go
 */
package controllers

import (
	"fmt"
	"gin/common"
	"time"

	"github.com/gin-gonic/gin"
)

// var (
// 	m    = make(map[int]uint64)
// 	lock sync.Mutex
// )

// type task struct {
// 	n int
// }

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

type Addr struct {
	city     string
	district string
}

/**
 * @description: 	测试并发接口
 * @param {*}
 * @return {*}
 */

func Concurrent(c *gin.Context) {
	// n := runtime.GOMAXPROCS(4)
	// fmt.Println("n=", n)
	// fmt.Println("run in main goroutine")
	// i := 1
	// for k := 1; k < 1000000; k++ {
	// 	go func() {
	// 		j := 1
	// 		for {
	// 			j = j + 1
	// 		}
	// 	}()
	// 	if i%10000 == 0 {
	// 		fmt.Printf("%d goroutine started\n", i)
	// 	}
	// 	i++
	// }
	// time.Sleep(3600 * time.Second)
	// go test() //起一个协程执行test()
	// for {
	// 	fmt.Println("i : runnging in main")
	// 	time.Sleep(time.Second)
	// }
	// num := runtime.NumCPU()
	// runtime.GOMAXPROCS(num)
	// fmt.Println(num)
	// for i := 0; i < 16; i++ {
	// 	t := &task{n: i}
	// 	go calc(t) //并发执行，
	// }
	// time.Sleep(10 * time.Second)
	// intChan := make(chan int, 3)
	// var i int
	// go func() {

	// 	for i = 1; i < 10; i++ {
	// 		intChan <- i
	// 		fmt.Println(i)
	// 	}

	// }()

	// value := <-intChan
	// fmt.Println("value : ", value)
	//testTranslateStruct()
	//testClose()
	//testMergeInput()
	//testQuit()
	// testPCB()
	go testWorker("xiaoming")
	go testWorker("hanmeimei")
	common.ResponJson(c, "并发测试成功", gin.H{}, 0)
}

//协程调用方法
func testWorker(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("name : ", name, time.Now())
		time.Sleep(1 * time.Second)
	}
}

// // 生产者消费者问题
// func testPCB() {
// 	fmt.Println("test PCB")

// 	intchan := make(chan int)
// 	quitChan := make(chan bool)
// 	quitChan2 := make(chan bool)

// 	value := 0

// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			value = value + 1
// 			intchan <- value

// 			fmt.Println("写入完成，值：", value)

// 			time.Sleep(time.Second)
// 		}
// 		quitChan <- true
// 	}()

// 	go func() {
// 		for {
// 			select {
// 			case v := <-intchan:
// 				fmt.Println("读取完成，值：", v)
// 			case <-quitChan:
// 				quitChan2 <- true
// 				return
// 			}
// 		}
// 	}()

// 	<-quitChan2
// 	fmt.Println("task is done ")
// }

// // 测试channel 用于通知中断退出的问题
// func testQuit() {
// 	g := make(chan int)
// 	quit := make(chan bool)

// 	go func() {
// 		for {
// 			select {
// 			case v := <-g:
// 				fmt.Println(v)
// 			case <-quit:
// 				fmt.Println("B退出")
// 				return
// 			}
// 		}
// 	}()

// 	for i := 0; i < 3; i++ {
// 		g <- i
// 	}

// 	quit <- true
// 	fmt.Println("testAB退出")
// }

// //将多个输入的channel进行合并成一个channel
// func testMergeInput() {
// 	input1 := make(chan int)
// 	input2 := make(chan int)
// 	output := make(chan int)

// 	go func(in1, in2 <-chan int, out chan<- int) {
// 		for {
// 			select {
// 			case v := <-in1:
// 				out <- v
// 			case v := <-in2:
// 				out <- v
// 			}
// 		}
// 	}(input1, input2, output)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			input1 <- i
// 			time.Sleep(time.Millisecond * 100)
// 		}
// 	}()

// 	go func() {
// 		for i := 20; i < 30; i++ {
// 			input2 <- i
// 			time.Sleep(time.Millisecond * 100)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			select {
// 			case value := <-output:
// 				fmt.Println("输出：", value)
// 			}
// 		}
// 	}()

// 	time.Sleep(time.Second * 5)
// 	fmt.Println("主线程退出")
// }

// //测试关闭channel
// func testClose() {
// 	ch := make(chan int, 5)
// 	sign := make(chan int, 2)

// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			ch <- i
// 			time.Sleep(time.Second)
// 		}

// 		close(ch)
// 		fmt.Println("the channel is closed")

// 		sign <- 0
// 	}()

// 	go func() {
// 		for {
// 			i, ok := <-ch
// 			fmt.Printf("%d, %v \n", i, ok)

// 			if !ok {
// 				break
// 			}

// 			time.Sleep(time.Second * 2)
// 		}

// 		sign <- 1
// 	}()

// 	<-sign
// 	<-sign
// }

// //测试channel传输复杂的Struct数据
// func testTranslateStruct() {
// 	personChan := make(chan Person, 1)

// 	person := Person{"xiaoming", 10, Addr{"shenzhen", "longgang"}}
// 	personChan <- person

// 	person.Address = Addr{"guangzhou", "huadu"}
// 	fmt.Printf("src person : %+v \n", person)

// 	newPerson := <-personChan
// 	fmt.Printf("new person : %+v \n", newPerson)
// }

// func calc(t *task) {
// 	var sum uint64
// 	sum = 1
// 	for i := 1; i < t.n; i++ {
// 		sum *= uint64(i)
// 	}
// 	fmt.Println(t.n, sum)
// 	lock.Lock()
// 	m[t.n] = sum
// 	lock.Unlock()
// }
