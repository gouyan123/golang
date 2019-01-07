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
	//顺序初始化结构体，每个成员变量必须初始化；
	var s1 *Student = &Student{23,"James",1,34,"USA"}
	fmt.Println(*s1)
	//指定成员初始化，没有指定的成员，自动赋值；
	s2 := &Student{id:23,name:"Jordan"}
	fmt.Println(*s2)
}

