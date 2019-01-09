package main

import "fmt"

type Student struct {
	id int
	name string
}

func main(){
	//创建切片，interface{}类似Object
	x := make([]interface{},3)
	x[0] = 1
	x[1] = "abcd"
	x[2] = Student{id:23,name:"James"}
	//range x表示遍历集合x，index表示集合下标，data表示集合下标对应元素
	for index,data := range x{
		//data.(int)表示数据是否为int类型，value为返回的值，ok为断言结果 true或者false；
		if value,ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int，内容为 %d\n",index,value)
		}
		if value,ok := data.(string); ok == true{
			fmt.Printf("x[%d] 类型为string，内容为 %s\n",index,value)
		}
		if value,ok := data.(Student); ok == true{
			fmt.Printf("x[%d] 类型为Student，内容为 name = %s id = %d\n",index,value.name,value.id)
		}
	}
}
