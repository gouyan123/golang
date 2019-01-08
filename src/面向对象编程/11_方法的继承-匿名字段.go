package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

//Person类型绑定一个函数PrintInfo()；*Person表示传递的是变量地址，而不是变量值，传变量值需要在函数中新开辟空间保存（不推荐）
func (person *Person) PrintInfo() {
	fmt.Println(person)
}
//Student继承Person，即持有Person类型的匿名变量，成员变量和方法都继承了
type Student struct {
	Person
	id int
	addr string
}
// 子类Student也绑定了一个同名函数 PrintInfo()
func (student *Student) PrintInfo()  {
	fmt.Println(student)
}
func main()  {
	s := new(Student)
	s.Person.name = "James"
	s.id = 23
	fmt.Println(s)
	s.Person.PrintInfo()
	s.PrintInfo()
}
