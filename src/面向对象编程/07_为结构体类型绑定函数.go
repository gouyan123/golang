package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

//带有接收者的函数 → 方法；函数绑定 类型对象，即为方法
//打印Person信息
func (person Person) PrintInfo(){
	fmt.Println("person = ",person)
}
//修改Person信息：要修改，需要传指针
func (person *Person) SetInfo(name string,sex byte,age int){
	person.name = name
	person.age = age
	person.sex = sex
	fmt.Println("person = ",person)
}

func main()  {
	//定义结构体Person，并初始化
	person := Person{name:"James",age:34}
	person.PrintInfo()
	(&person).SetInfo("gy",1,30)
}
