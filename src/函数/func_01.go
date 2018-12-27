package 函数

import (
	_ "fmt"
"fmt")

func swap(a int,b int)(int,int)  {
	return b,a;
}

func add(a *int)(int)  {
	*a = *a + 1;
	return *a;
}

func main()  {
	a := 1;
	var b int;
	b = add(&a);
	fmt.Print(a,b);
	//b := 2;
	//var c int;
	//var d int;
	//c,d = swap(a,b);
	//fmt.Print(c,d);
}
