/*匿名函数就是一个函数名/函数引用地址*/
package main

import (
	_ "fmt"
	"fmt"
)

func main()  {
	var x int = 8888
	/*定义匿名函数，没有函数名字，将函数引用地址赋给 f，还没调用*/
	f := func(){
		fmt.Println(x)
	}
	/*f只是一个函数类型变量，加()进行调用，即调用匿名函数*/
	f()

	/*定义匿名函数并直接调用func(){}()*/
	func(){
		fmt.Println("执行匿名函数")
	}()
	/*定义带参数的匿名函数并直接调用*/
	func(i int,j int){
		fmt.Println(i,j)
	}(10,20)
	/*定义有参有返回值的匿名函数 将函数地址赋给 fn*/
	fn := func(i int,j int) int {
		return i + j
	}
	fnResult := fn(1000,29872)
	fmt.Println(fnResult)
}
