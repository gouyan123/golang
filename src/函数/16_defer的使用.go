package 函数

import (
	_ "fmt"
	"fmt"
)
/*defer语句最后执行*/
func main()  {
	defer fmt.Println("defer语句无论在函数中什么位置，都最后执行")
	fmt.Println("普通语句都执行完，最后方法结束前再执行 defer 语句")
}
