package main

import "fmt"
/*执行该函数时，新建一块栈内存，将变量a,b传到函数swap时，再在函数内存空间开辟空间保存a,b，函数执行完，内存
空间自动释放*/
func swap01(a int,b int)  {
	temp := a
	a = b
	b = temp
	fmt.Println("swap内 a = ",a," b = ",b)
}
/*该函数修改的是指针变量p1，p2指向的a，b的内存内容*/
func swap02(p1 *int,p2 *int)  {
	temp := *p1
	*p1 = *p2
	*p2 = temp
}
func main()  {
	a,b := 10,20
	/*通过函数swap()交换a,b的内容，站在变量角度，变量本身值传递*/
	swap01(a,b)
	fmt.Println("调用swap01(a,b)后 a = ",a," b = ",b)
	swap02(&a,&b)
	fmt.Println("调用swap02(&a,&b)后 a = ",a," b = ",b)
}
