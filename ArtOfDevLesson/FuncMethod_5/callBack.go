package main

import "fmt"

func doSomething(callBack func(int, int) int, s string) int {
	result := callBack(5, 1)
	fmt.Println(s)
	return result
}

//							возвращает функцию которая возвращает int
func totalPrice(initPrice int) func(int) int {
	sum := initPrice         // выполнится 1 раз при вызове функции
	return func(i int) int { // при вывове меняется а sum остается неизменным
		sum += i
		return sum
	}
}

func main() {
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := doSomething(sumCallback, "plus")
	fmt.Println(result)

	multipleCallback := func(n1, n2 int) int {
		return n1 * n2
	}

	result = doSomething(multipleCallback, "multiple")
	fmt.Println(result)

	// Замыкание
	orderPrice := totalPrice(1) // sum = 1
	fmt.Println(orderPrice(1))  // sum = 1 + 1
	fmt.Println(orderPrice(1))  // sum = 2 + 1

	// остановился на 11 минуте
}
