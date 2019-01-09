package main

import (
	"fmt"
	"time"
)

func newTask()  {
	for  {
		fmt.Println("This is a newTask goroutine协程")
		//延时 1s
		time.Sleep(time.Second)
	}
}

func main()  {
	//新建一个协程 goroutine，两个任务，变串联（同步）为并联（异步）；
	go newTask()
	for  {
		fmt.Println("This is a main goroutine协程")
		//延时 1s
		time.Sleep(time.Second)
	}

}