package main

import "fmt"

/*回调函数：a函数的参数含有函数类型名，这个函数类型名就是回调函数，在a函数中调用这个类型名*/

/*定义函数类型*/
type myFuncType func(int,int) int


/*定义回调函数，方法参数中含有函数变量*/
/*多态：多种形态，调用同一个接口，不同的实现有不同的表现，例如 加，减，乘，除*/
func Calc(a int,b int,fTest myFuncType) (result int) {
	/*先有想法，后面再实现功能，fTest函数现在还不存在实现，就可以定义使用*/
	/*重点：类似接口，可以实现解耦，增加扩展性*/
	result = fTest(a,b)
	return result
}

/*定义函数类型 的具体实现函数*/
func mul(a int,b int) int {
	return a * b
}

func main()  {
	result := Calc(3,7,mul)
	fmt.Print(result)
}
