package main

import (
	_ "fmt"
	"fmt"
)

func Add(a int,b int) int {
	return a + b
}
func Minus(a int,b int) int {
	return a - b
}
/*函数类型也是一种数据类型，FuncType是一个函数类型，func(int,int) int是函数类型的定义*/
type FuncType func(int,int) int

func main()  {
	/*函数名 fTest就是一个变量，其类型为FuncType，是一种函数类型*/
	var fTest FuncType
	/*将函数Add地址给fTest，函数名是一个变量*/
	fTest = Add
	/*函数名是一个变量，函数名 + () 表示调用该函数*/
	result01 := fTest(3,4)
	fmt.Println(result01)

	fTest = Minus
	result02 := fTest(9,8)
	fmt.Println(result02)
}