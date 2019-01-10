package main

import (
	"fmt"
	_ "fmt"
)
//声明 全局变 a
var a int = 1

func main()  {
	//声明 局部变量 a
	//不同作用域，允许使用同名变量，使用变量的原则：就近原则；
	//var a int = 9
	fmt.Println(a)
}

