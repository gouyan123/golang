package 函数

import (
	_ "fmt"
	"os"
	"fmt"
)

func main()  {
	/*接收用户传递的参数，都是以字符串形式传递*/
	list := os.Args
	n := len(list)
	fmt.Println("n = ",n)
	for i,data := range list{
		fmt.Println(i,data)
	}
}
