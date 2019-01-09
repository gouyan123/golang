package main

import "fmt"

func main()  {
	//创建一个无缓存channel
	ch := make(chan int)

	//新建协程goroutine
	go func() {
		for i := 0;i < 5 ;i++  {
			fmt.Println("子协程 i = ",i)
			//往管道 ch里面存数据，可以存入3个数据，如果存入3个数据并没有取出，那么协程将阻塞并不停轮休判断，否则协程将阻塞于此
			ch <- i
		}
	}()
}
