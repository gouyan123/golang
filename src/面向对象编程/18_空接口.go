package main

import "fmt"

//interface{} 等价于 Object；...interface{} 等价于 Object[]
func doSometing(args ...interface{})  {
	//_ 表示集合元素下标，i 表示集合下标对应元素；range表示遍历 args；
	for _,i := range args {
		fmt.Println(i)
	}
}

func main()  {
	//i 为空接口类型，可以保存任何数据类型，类似 java中的 Object；
	var i interface{}
	i = 23
	fmt.Printf("i = %d\n",i)
	i = "James"
	fmt.Printf("i = %s\n",i)
}
