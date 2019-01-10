package main

import (
	"fmt"
	"time"
)

func main()  {
	//创建定时器timer，设置时间为2s，2s后往timer通道里面写内容（内容为当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：",time.Now())
	//从time通道中取出内容，如果time通道没有内容，则阻塞于此，直到time通道被写入内容
	t := <- timer.C
	fmt.Println("t = ",t)
}
