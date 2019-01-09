package main

import (
	"fmt"
	"time"
)

func main()  {
	//创建管道 ch，一个协程goroutine往里面存数据，另一个协程goroutine从里面取数据，当管道ch里面没有数据时，协程会阻塞在 <- ch处 直到管道 ch里面有数据
	ch := make(chan string)

	defer fmt.Println("主协程结束")

	//创建并开启一个协程goroutine
	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0 ;i < 2 ;i++  {
			fmt.Println("子协程 i = ",i)
			time.Sleep(time.Second)
		}
		ch <- "abc"
	}()
	//如果管道 ch里面没有数据，主协程则阻塞在这里直到管道 ch里面有数据为止
	str := <-ch
	fmt.Println("str = ",str)
}
