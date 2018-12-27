package main

import "fmt"

func main()  {
	//var a int = 10			/*声明并初始化变量a*/
	var p *int				/*声明一个指针变量p，指向int型，此时只有变量地址，并未开辟变量内存*/
	p = new(int)			/*给变量p开辟内存空间*/
	*p = 666				/*p的内存中的地址内容指向的内存值设为666*/
	fmt.Println("p = ",*p,"&p = ",&p,"*(&p) = ",*(&p))

	q := new(int)			/*自动推导类型*/
	*q = 999
	fmt.Println("*q = ",*q)
}