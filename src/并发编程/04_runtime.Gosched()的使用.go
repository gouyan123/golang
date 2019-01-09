package main

import (
	"fmt"
	"runtime"
)

func main(){
	//go表示开启一个新协程goroutine，func(){}()为匿名函数
	go func() {
		for i := 0;i < 5;i++ {
			fmt.Println("son goroutine")
		}
	}()	//	()表示直接调用匿名函数

	for i := 0;i < 2;i++ {
		//让出 CPU时间片，让其他协程先执行，其他线程执行完，再争取CPU时间片
		runtime.Gosched()
		fmt.Println("main goroutine")
	}	
}
