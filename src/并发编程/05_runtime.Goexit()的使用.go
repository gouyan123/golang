package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("test()")
	//return			//终止此函数
	runtime.Goexit()
	fmt.Println("dddddddddddddddddddddddddd")
}


func main()  {
	//go 创建新协程goroutine；func(){}()为匿名函数
	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		//调用别的函数 test()
		test()
		fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	}()
	//死循环，避免主协程 结束
	for{}
}
