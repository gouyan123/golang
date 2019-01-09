package main

import "fmt"

//定义接口类型
type Humaner interface {		//子集合
	//定义方法，只有方法声明，没有方法实现；重点：接口方法实现 接口方法绑定其他类型，由其他类型实现；
	sayhi()
}

type Personer interface {		//超集合；超集继承子集，因此 超集 可以转换为 子集
	//组合匿名字段，就是继承该 字段类型了，拥有了该字段类型的所有方法；
	Humaner
	//Personer接口类型自定义的方法
	sing(lrc string)
}

//定义Student结构体类型，实现 sayhi()和sing(lrc string)方法
type Student struct {
	id int
	name string
}
//Student结构体类型实现 接口方法sayhi()
func (s *Student) sayhi()  {
	fmt.Printf("Student[%s,%d] sayhi\n",s.name,s.id)
}
//Student结构体类型实现 接口方法sing(lrc string)
func (s *Student) sing(lrc string)  {
	fmt.Printf("Student[%s,%d] sing %s\n",s.name,s.id,lrc)
}
func main()  {
	//不用自己处理地址和指针，编译时会自动处理；使用值传递还是指针地址传递，定义方法时，指定即可；
	s := new(Student)
	s.id = 23
	s.name = "James"
	var p Personer = s
	//var p Personer = &Student{23,"James"}
	p.sayhi()
	p.sing("abcdefg")
}
