package main

import (
	"fmt"
	"runtime"
)

func main()  {
	//n为服务器总核数，2为设置使用核数
	n := runtime.GOMAXPROCS(8)
	fmt.Println("n = ",n)
	for{
		go fmt.Print(1)
		fmt.Print(0)
	}
}
