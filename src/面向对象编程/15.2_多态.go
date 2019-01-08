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
//*****多态：定义一个普通函数，函数的参数为接口类型；多态：函数只有一个，传入不同的接口实现，会有不同的表现
func WhoSayHi(h Humaner)  {
	h.sayhi()
}

func main()  {
	s := &Student{"James",23}
	t := &Teacher{"Go",17}
	WhoSayHi(s)
	WhoSayHi(t)
	//--------------------------------------------------------------
	//创建切片
	x := make([]Humaner,2)
	x[0] = s
	x[1] = t
	// _ 代表下标，h 代表下标对应的元素；range 表示遍历 集合
	for _,h := range x {
		h.sayhi()
	}
}
