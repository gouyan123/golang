package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}


func (person Person) SetInfoValue()  {
	fmt.Println("SetInfoValue()")
}

func (person *Person) SetInfoPointer()  {
	fmt.Println("SetInfoPinter()")
}

func main()  {
	p := new(Person)
	p.name = "James"
	p.sex = 1
	p.age = 34
	fmt.Println(p)
	p.SetInfoValue()
	//会取p的地址 &p，再调用SetInfoPointer()
	p.SetInfoPointer()
	//两者等价；
	(*p).SetInfoPointer()
}

