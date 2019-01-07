package main

import "fmt"

//定义结构体类型，结构体类型就是一种数据类型
type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main()  {
	//定义一个Student结构体类型的 普通变量
	var s1 Student
	//操作结构体Student中的成员变量，需要使用 . 运算符
	s1.id = 23
	s1.name = "James"
	s1.addr = "USA"
	fmt.Println(s1)
}
