package 函数
import (
	_ "fmt"
	"fmt"
)
//func main(){
//	a := 10
//	b := 20
//	defer func(){
//		fmt.Printf("内部 a = %d,b = %d\n",a,b)
//	}()
//	a = 111
//	b = 222
//	fmt.Printf("外部 a = %d,b = %d\n",a,b)
//}

func main(){
	a := 10
	b := 20
	/*该匿名函数按执行顺序先加载完毕，只是暂时不调用，等外部函数结束前再调用defer定义的内容，但是加载函数
	* 是按顺序加载
	*/
	defer func(a int,b int){
		fmt.Printf("内部 a = %d,b = %d\n",a,b)
	}(a,b)
	a = 111
	b = 222
	fmt.Printf("外部 a = %d,b = %d\n",a,b)
}
