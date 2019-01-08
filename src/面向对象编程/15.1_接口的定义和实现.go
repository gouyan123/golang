package main

import "fmt"

//定义接口类型
type Humaner interface {
	//定义方法，只有方法声明，没有方法实现；重点：接口方法实现 接口方法绑定其他类型，由其他类型实现；
	sayhi()
}
//定义 Student 结构体类型，实现 Humaner接口类型的方法
type Student struct {
	name string
	id int
}
type Teacher struct {
	name string
	id int
}

//Humaner接口方法 sayhi() 绑定 结构体类型 Student；
func (student *Student) sayhi()  {
	fmt.Printf("Student[%s,%d] sayhi\n",student.name,student.id)
}
func (teacher *Teacher) sayhi()  {
	fmt.Printf("Student[%s,%d] sayhi\n",teacher.name,teacher.id)
}


func main()  {
	//定义接口类型变量，只要实现了该接口方法的类型，该实现类型就可以该接口类型变量赋值
	var h Humaner
	//s := &Student{"James",23}

	s := new(Student)
	s.name = "James"
	s.id = 23
	h = s
	h.sayhi()
	//----------------------------------------------------------------------
	t := new(Teacher)
	t.name = "Java"
	t.id = 1
	h = t
	h.sayhi()
}