package main

import (
	_ "fmt"
	"fmt"
)

func main() {
	a, b := 10, 20
	//匿名函数先加载，在外层函数执行完毕之前再调用
	defer func(x int) { // a以值传递方式传给x
		fmt.Println("defer:", x, b) // b 闭包引用
	}(a)								//(a)代表调用此匿名函数，先把参数传递过去，只是defer方法最后被调用，此时 (a)捕获的是 a = 10；

	a += 10
	b += 100

	fmt.Printf("a = %d, b = %d\n", a, b) //20 120

}
