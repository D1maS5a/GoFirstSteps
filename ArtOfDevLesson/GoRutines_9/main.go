package main

import (
	"fmt"
	"time"
)

// Сделать калькулятор на горутинах

func main() {
	// ch := make(chan int)
	// // fmt.Println("Hello world!")
	// // word := "Hello"
	// // go say("Hello"/*, ch*/)
	// go sayHello(ch)
	// s := <-ch
	// fmt.Println(s)
	// for i := range ch {
	// 	fmt.Println(i)
	// }
	// fmt.Println("1")
	// fmt.Println("2")
	// fmt.Println("3")
	// fmt.Println("4")
	// fmt.Println("5")
	// // time.Sleep(1 * time.Second)
	// fmt.Println(<-ch)
	data := make(chan int)
	exit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)
}

// асинхронно
func say(word string /*,ch chan string*/) {
	// time.Sleep(2 * time.Second)
	fmt.Println(word)
	// ch <- "exit"
}

func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		// say("Hello")
		exit <- i
	}
	close(exit)
	// exit <- "s"
}

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("waiting")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
