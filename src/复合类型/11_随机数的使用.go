package main

import (
	"math/rand"			/*rand不是全局的，不能直接使用，需要引入当前文件，才能使用rand中文件中的方法*/
	"fmt"
	"time"
)

func main()  {
	/*设置种子，只需要一次*/
	/*如果种子参数一样，每次运行产生的随机数都一样*/
	//rand.Seed(666)
	/*以当前时间为种子参数*/
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<5;i++{
		/*产生随机数*/
		fmt.Println(rand.Intn(100))
	}
}
