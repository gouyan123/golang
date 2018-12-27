package 函数
import (
	_ "fmt"
	"fmt"
)
/*回调函数：a函数的参数含有函数类型名，这个函数类型名就是回调函数，在a函数中调用这个类型名*/

/*定义函数类型名*/
type myFuncType func(int,int) int
/*定义函数类型名的具体实现函数*/
func mul(a int,b int) int {
	return a * b
}
/*定义回调函数*/
func Calc(a int,b int,fTest myFuncType) int {
	result := fTest(a,b)
	return result
}

func main()  {
	result := Calc(3,7,mul)
	fmt.Print(result)
}
