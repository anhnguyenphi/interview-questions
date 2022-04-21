package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int

	go send(&ch)
	go rc(&ch)
	time.Sleep(time.Second * 10)
}

func send(ch *chan int) {
	fmt.Println("Sending value to channnel start")
	time.Sleep(time.Second * 2)
	fmt.Println(ch)
	*ch <- 2
	fmt.Println("Sending value to channnel finish")
}

func rc(ch *chan int) {
	time.Sleep(time.Second)
	*ch = make(chan int)
	fmt.Println(<-*ch)
}
