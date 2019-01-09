package main

import (
	"fmt"
	"time"
)

//定义一个打印机，参数为字符串，按每个字符打印
func Printer(str string)  {
	for _,data := range str {
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func person1()  {
	Printer("Hello")
}

func person2()  {
	Printer("World")
}

func main()  {
	//新建 2个协程，代表2个人，2个人同时使用打印机
	go person1()
	go person2()
	//死循环，避免 主协程goroutine退出
	for {}
}
