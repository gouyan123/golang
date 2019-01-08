package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

// 参数为普通变量，非指针，函数会新开辟一个空间，保存Person，函数运行结束后，释放空间,不会改变原Person
func (person Person) SetInfoValue(name string,sex byte,age int)  {
	person.name = name
	person.sex = sex
	person.age = age
	fmt.Println("SetInfoValue: p = ",person)
}
//参数为指针变量，表示使用Person的地址而不是内存内容，会改变原参数
func (person *Person) SetInfoPointer(name string,sex byte,age int)  {
	person.name = name
	person.sex = sex
	person.age = age
	fmt.Println("SetInfoPinter: p = ",person)
}

func main()  {
	//开辟 类型Person的内存空间
	p := new(Person)
	p.SetInfoValue("James",1,34)
	fmt.Println(p) 										// 参数值保存到新创建的Person中了，没有改变实参 p，因此p为空
	p.SetInfoPointer("Jordan",1,40)
	fmt.Println(p)
}
