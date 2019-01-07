package main

import "fmt"

//实现 2个数相加：①面向过程；②面向对象
//①面向过程，使用函数
func Add(a int,b int) int{
	return a + b
}
//②面向对象，使用方法：给某个类型绑定一个函数
type myint int

func (tmp myint) Add(other myint) myint{
	return tmp + other
}
func main()  {
	//--------------------------面向过程 调用函数
	r1 := Add(1,2)
	fmt.Println(r1)
	//--------------------------面向对象 调用方法：调用格式 变量名.函数()
	//定义一个 myint类型，再调用该类型绑定的函数
	var a myint = 1
	r2 := a.Add(2)
	fmt.Println(r2)
}
