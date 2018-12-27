package 函数
import (
	_ "fmt"
	"fmt"
)
func test(i int) int {
	return 9/i
}

func main()  {
	defer fmt.Println("111111111111")
	defer fmt.Println("222222222222")
	/*当test(0)前不加defer时*/
	test(0)
	defer fmt.Println("3333333333333")
}
