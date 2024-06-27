package main

import (
	"fmt"
)

func say(greet string, ch chan string, ch2 chan int) {
	fmt.Println(greet)
	ch <- "Hello from Gorutine"
	ch2 <- 1234
}

func newSay(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	// ch := make(chan string)
	// ch2 := make(chan int)

	// go say("Hello GO", ch, ch2)
	// time.Sleep(2 * time.Second)
	// data := <-ch2
	// data2 := <-ch
	// fmt.Println(<-ch, <-ch2)

	ch := make(chan int)
	go newSay(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
