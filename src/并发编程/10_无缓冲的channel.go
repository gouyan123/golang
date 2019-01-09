package main

import (
	"fmt"
	"time"
)

func main()  {
	//创建一个无缓存的 channel；0 表示channel中能缓存 0个值
	ch := make(chan int,0)
	fmt.Printf("len(ch) = %d,cap(ch) = %d\n",len(ch),cap(ch))
	//新建协程goroutine
	go func() {
		for i := 0;i < 3 ;i++  {
			fmt.Println("子协程 i = ",i)
			//往管道 ch里面存数据，只有存入的数据被取走后，才能继续存，否则会阻塞，直到管道取出数据
			ch <- i
		}
	}()

	//延时2s
	time.Sleep(time.Second*2)

	for i := 0;i < 3 ;i++  {
		//从管道 ch读取数据，当管道里面没有数据时会阻塞，直到管道存入数据
		num := <- ch
		fmt.Println("主协程 num = ",num)
	}
}

