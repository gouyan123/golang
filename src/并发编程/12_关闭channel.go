package main

import "fmt"

func main()  {
	//创建一个无缓存channel
	ch := make(chan int)

	//新建协程goroutine
	go func() {
		for i := 0;i < 5 ;i++  {
			//往管道 ch里面存数据
			ch <- i
			fmt.Println("子协程 i = ",i)
		}
		//不需要再写数据时，关闭 channel
		close(ch)
	}()

	for {
		//ok为true，说明管道没关闭；num为从管道 ch中取出的值
		if num,ok := <-ch; ok == true{
			fmt.Println(num)
		}else {
			//管道关闭，跳出for循环
			break
		}
	}
}
