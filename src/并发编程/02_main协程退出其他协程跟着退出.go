package main

import (
	"fmt"
	"time"
)

func main()  {
	/*函数调用前加关键字go，开启一个新协程，一个协程对应一个任务，函数就是任务*/
	go func() {
		var i int = 0
		for{
			i++
			fmt.Println("子协程 i = ",i)
			time.Sleep(time.Second)
		}
	}()
	/*让main协程阻塞在这里，否则main协程退出，子协程也跟着退出了*/
	for{
		fmt.Println("main协程")
		time.Sleep(time.Second)
	}

}
