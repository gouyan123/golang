/*匿名函数就是一个函数名/函数引用地址*/
package 函数

import (
	_ "fmt"
	"fmt"
)

func main()  {
	/*定义匿名函数 将函数引用地址赋给 f*/
	f := func(){
		fmt.Println("执行匿名函数f")
	}
	/*调用匿名函数，函数引用后加()*/
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
