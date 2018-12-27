package main

import "fmt"

func modify(a *[5]int)  {
	(*a)[0] = 666
	fmt.Println("modify里 a = ",a)
}

func main()  {
	/*数组长度必须是常量，不能写变量*/
	/*全部初始化*/
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println("a = ",a)
	/*部分初始化，没有初始化的元素，自动赋值为0*/
	var c [5]int = [5]int{1,2,3}
	fmt.Println("c = ",c)
	/*指定某个元素初始化，下标2赋值10，下标4赋值99，其他下标赋值0*/
	d := [5]int{2:10,4:99}
	fmt.Println("d = ",d)
	/*数组做参数是值拷贝，即新开辟一块内存空间保存数组参数，操作的是新开辟的数组空间，不是原来的空间*/
	/*数组指针做参数，在函数中操作的是原数组内存空间*/
	modify(&a)
	fmt.Println("main里 a = ",a)
}
