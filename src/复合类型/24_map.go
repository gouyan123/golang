package main

import "fmt"

func main()  {
	/*----------------------------------map创建和初始化------------------------------------*/
	/*map类型声明 格式为：map[keyType]valueType*/
	var m1 map[int]string
	/*创建map类型变量 m := make(map[keyType]valueType)*/
	m1 = make(map[int]string)
	m1[1] = "James"
	m1[2] = "Jordan"
	fmt.Println(m1)
	/*m[key] 判断key值是否存在，返回对应value值和ok*/
	value,ok := m1[1]
	if ok{
		fmt.Println("value = ",value)
	}
	/*删除map变量中key的key-value对 delet(map变量,key) */
	delete(m1,2)
	fmt.Println(m1)
	/*map做函数参数实际是引用传递*/

}
