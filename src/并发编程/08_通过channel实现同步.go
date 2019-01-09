package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个 channel：一边存数据，另一边取数据，没有数据时阻塞；
var ch = make(chan int)

//定义一个打印机，参数为字符串，按每个字符打印
func Printer(str string)  {
	for _,data := range str {
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}
//person1()执行完，向管道 ch保存一个int型 数据，person2() 从管道 ch获取一个int型数据，如果与person1()存的数据匹配，person2()才能执行；
func person1()  {
	Printer("Hello")
	ch <- 666				//person1()任务执行完毕，向管道 ch存入数据；
}

func person2()  {
	<- ch					//从管道 ch取出数据，如果管道 ch没有数据，进程就阻塞在这里等待，不停轮休判断，直到管道 ch有数据为止；
	Printer("World")
}

func main()  {
	//新建 2个协程，代表2个人，2个人同时使用打印机
	go person1()
	go person2()
	//死循环，避免 主协程goroutine退出
	for {}
}
