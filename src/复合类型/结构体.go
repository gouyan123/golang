package main

import "fmt"

/*定义一个结构体*/
type Student struct {
	name string
	age int
}
func main()  {
	/*--------------------------------结构体声明与初始化--------------------------*/
	/*顺序初始化，全部成员变量必须初始化*/
	s1 := Student{"James",33}
	fmt.Println("s1 = ",s1)
	/*指定成员变量初始化，可以部分初始化，其他自动赋值为0*/
	s2 := Student{name:"James"}
	fmt.Println("s2 = ",s2)
	/*--------------------------------结构体指针变量初始化--------------------------*/
	var p1 *Student = &Student{"Jordan",38}
	fmt.Println("*p1 = ",*p1)
	/*--------------------------------结构体成员变量的使用--------------------------*/
	var s3 Student
	s3.name = "Kobe"
	s3.age = 34
	fmt.Println("s3 = ",s3)
	/*--------------------------------指针有合法指向后才操作成员--------------------------*/
	/*先定义一个结构体变量*/
	var s4 Student
	/*再定义一个指针，保存s4地址*/
	var p2 *Student = &s4
	/*通过指针操作成员变量：①(*p2).id = 99;②var p3 *Student = new(Student) 新开辟一个结构体空间*/
	(*p2).name = "Onel"
	(*p2).age = 38
	fmt.Println("s4 = ",s4)
	var p3 *Student = new(Student)
	(*p3).name = "Curry"
	(*p3).age = 30
	fmt.Println("*p3 = ",*p3)
	/*同类型的两个结构体可以相互赋值*/
	/*结构体变量比较：s1 == s2 比较的是内存内容，不是内存地址*/
	/*--------------------------------结构体作为函数参数--------------------------*/
	/*①值传递，函数接收后，自己开辟空间保存，不会改变原来的值；②引用传递，会改变原来的值*/
	test01(s1)/*值传递，形成无法改变实参*/
	fmt.Println("test01执行后 s1 = ",s1)
	test02(&s1)
	fmt.Println("test02执行后 s1 = ",s1)
	/*--------------------------------可见性-------------------------------------*/
	/*要想某个符号对其他包可见，这个符号需要首字母大写，在其他文件中调用格式：包名.符号*/
}

func test01(s Student)  {
	s.age = 18
	fmt.Println("test01值传递 s = ",s)
}
func test02(p *Student)  {
	(*p).age = 18
	fmt.Println("test02引用传递 *p = ",*p)
}