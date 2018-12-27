package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int)  {
	for i:=0;i<10;i++{
		out <- i*i
	}
	close(out)
}
func consumer(in <-chan int)  {
	//num := <-in
	for num := range in {
		fmt.Println("num = ",num)
	}

}
func main()  {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Minute)
}
