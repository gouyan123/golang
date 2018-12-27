package main/*包 复合类型的入口*/
import "fmt"
func main()  {
	var a int = 10			/*每个变量有两层含义：①变量内存；②变量地址*/
	fmt.Println("a表示内存内容",a)
	fmt.Println("&a表示内存内容",&a)
	/*需要指针类型保存变量地址 *int保存int的地址*/
	var p *int				/*定义指针类型变量，包括两部分：①变量p内存内容，保存其他变量的地址，*p指向
							*p的内存保存的地址指向的内存；②变量p内存地址，&p获取变量p的内存地址*/
	p = &a					/*将a地址&a的值赋给p，p的类型是 *int */
	fmt.Println("p = ",p)
	*p = 666				/* *p操作的不是p的内存，而是p所指向的内存（就是a）*/
	fmt.Println("*p表示地址p指向的内容",*p)
	/* &a表示取得变量a的内存地址，*p表示以p的内存内容为地址 找到指向内容*/
}


