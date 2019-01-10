package main

func main()  {
	/*内置函数make创建一个双向 channel，里面存储int类型*/
	ch := make(chan int)
	/*双向channel能隐式转换为单向channel，写*/
	var writeCh chan<- int = ch
	/*双向channel能隐式转换为单向channel，读*/
	var readCh <-chan int = ch
	writeCh <- 6
	<-readCh
	//单向无法转换为双向
	c := make(chan int,3)
	var send chan<- int = c // send-only
	var recv <-chan int = c // receive-only
	send <- 1
	//<-send //invalid operation: <-send (receive from send-only type chan<- int)
	<-recv
	//recv <- 2 //invalid operation: recv <- 2 (send to receive-only type <-chan int)
}
