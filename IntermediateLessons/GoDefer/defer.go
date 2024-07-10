package main

import "fmt"

func main() {
	// defer fmt.Println("world") // Когда закончится выполнится
	// stack Last In First Out
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("Hello")
}
