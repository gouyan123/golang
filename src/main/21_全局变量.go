package main

import (
	_ "fmt"
	"fmt"
)

//定义在函数外面的变量，称为全局变量，全局变量在任何地方都可以用
//全局变量的声明
var a int

func t() {
	fmt.Printf("test a = %d\n", a)
}

func main() {
	a = 10
	fmt.Printf("main a = %d\n", a) //main a = 10
	t() //test a = 10
}
