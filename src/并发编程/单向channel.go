package main

func main()  {
	///*内置函数make创建一个双向 channel，里面存储int类型*/
	//ch := make(chan int)
	///*创建单向写入channel*/
	//var writeCh chan<- int =ch
	///*创建单向读取channel*/
	////var readCh <-chan int = ch
	//writeCh <- 6
	////<-readCh
	///*var x int = <-readCh
	//fmt.Println("x = ",x)*/
	c := make(chan int,3)
	var send chan<- int = c // send-only
	var recv <-chan int = c // receive-only
	send <- 1
	//<-send //invalid operation: <-send (receive from send-only type chan<- int)
	<-recv
	//recv <- 2 //invalid operation: recv <- 2 (send to receive-only type <-chan int)
}
