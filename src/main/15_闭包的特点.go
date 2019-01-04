package main

import (
	"fmt"
	_ "fmt"
)

func test01() int {
	var x int		/*函数test01被调用的时候，才给x分配内存，才初始化为0*/
	x++
	return x		/*函数调用完毕，自动释放 x 内存空间*/
}
/*函数的返回值是一个函数引用/类型 匿名函数就是一个函数引用*/
func test02() func() int {
	var x int
	return func() int {	/*此处返回的是一个函数引用，test02()() test02()返回函数引用 引用()调用内部函数*/
		x++
		return x + x
	}
}

func main()  {
	fmt.Println(test01())
	fmt.Println(test02()())
}
