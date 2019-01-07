package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person				// 匿名字段：只有类型，没有名字；匿名字段作用：Student结构体继承了Person结构体的所有字段
	id int
	addr string
}

type Teacher struct {
	int					//内置类型 匿名变量，实际没有意义
	//Person				//结构体 匿名变量
	*Person				//指针 匿名变量
	mystring			//自定义类型 匿名变量
}

type mystring string	//自定义类型 mystring

func main()  {
	//----------------------结构体 成员变量 初始化
	//顺序初始化 结构体Student；结构体Student继承结构体Person，即Student拥有 Person类型匿名字段
	s1 := Student{Person{"James",1,34},23,"USA"}
	// %+v 表示输出详细信息
	fmt.Printf("s1 = %+v\n",s1)

	//指定名称初始化 结构体Student；
	s2 := Student{Person:Person{name : "Jordan"},id : 23}
	fmt.Printf("s2 = %+v\n",s2)
	//----------------------结构体 成员变量 操作
	var s3 Student
	s3.Person = Person{name:"Kobe",age:40}
	s3.id = 24
	fmt.Println("")
	fmt.Printf("s3 = %+v\n",s3)

	var t1 Teacher
	t1 = Teacher{1,&Person{name:"Tom"},"haha"}
	fmt.Printf("t1 = %+v\n",t1)

	var t2 Teacher
	t2.Person = new(Person)			// 先分配空间，再赋值
	t2.Person.name = "www"
	fmt.Println(t2.Person.name)
}
