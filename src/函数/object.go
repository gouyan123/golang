package 函数
import ("fmt")
/*接口*/
type Animal interface {
	fly();
	run();
	eat();
}
/*类，类不用实现接口，覆写接口方法时，指定类就行*/
type Bird struct{

}
/*覆写接口方法fly()，指定具体实现类Bird*/
func (bird Bird) fly(){
	fmt.Println("鸟会飞");
}
/*覆写接口方法run()，指定具体实现类Bird*/
func (bird Bird)run(){
	fmt.Println("鸟会跑");
}
func (bird Bird)eat()  {

}
func main()  {
	/* := 代表声明变量并赋值；= 代表赋值*/
	/*go语言中，声明的变量一定要使用，否则删掉，要不然会报错*/
	var animal Animal;
	bird := new(Bird);
	animal = bird;
	animal.fly();
	animal.run();
	/*interface{}类型处理，类型查询*/
	var v1 interface{};
	v1 = 6.78;
	if v,ok := v1.(float64); ok {/*ok是否为true*/
		fmt.Print(v,ok);
	}
}
