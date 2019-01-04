package main

import (
	_ "fmt"
	"fmt"
)

func main()  {
	//a := 10
	//str := "James"
	///*匿名函数定义在函数里面，构成闭包*/
	//func(){
	//	/*匿名函数通过变量地址引用修改的变量，而不是把变量赋值到该函数内部，因此，外部函数的变量也会改变*/
	//	a = 99
	//	str = "James No.1"
	//	fmt.Println("匿名函数内部",a,str)
	//}()
	//fmt.Println("匿名函数外部",a,str)

	a := 99
	str := "Jordan"

	func(){
		a = 100
		str = "James"
		fmt.Println(a,str)
	}()
	fmt.Println(a,str)
}
