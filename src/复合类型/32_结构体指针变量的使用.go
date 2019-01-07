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
	//指针有合法指向后，才可以操作成员变量
	//先定义一个普通结构体变量
	var s Student
	//再定义一个指针变量，保存 s 的地址
	var p *Student
	p = &s
	//通过指针操作成员，有 2种方法：p1.id 与 (*p).id完全等价
	(*p).id = 23
	p.name = "Jordan"
	fmt.Println(s,&s,p,*p);

	//通过 new 申请一个结构体
	p2 := new(Student)
	p2.id = 24
	p2.name = "Kobe"
	fmt.Println(p2,*p2);
}
