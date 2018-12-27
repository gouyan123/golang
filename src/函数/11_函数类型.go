package 函数

import (
	_ "fmt"
	"fmt"
)

func Add(a int,b int) int {
	return a + b
}
func Minus(a int,b int) int {
	return a - b
}
/*函数也是一种数据类型，通过type给函数类型起名，FuncType是一个函数类型，func(int,int) int是函数类型的定义*/
type FuncType func(int,int) int		/*函数类型为FuncType，现在是空的，可以接收其他符合格式的函数类型地址
									*/
func main()  {
	/*定义FuncType类型的函数变量*/
	var fTest FuncType
	/*将函数Add地址给fTest*/
	fTest = Add
	result01 := fTest(9,8)
	fmt.Print(result01)

	fTest = Minus
	result02 := fTest(9,8)
	fmt.Print(result02)
}