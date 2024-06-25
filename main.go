package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	fmt.Println("Привет от Go!")

	// var a string
	// fmt.Scan(&a)

	var age uint8 = 23
	fmt.Println(age)

	var number float64 = 256.672
	fmt.Println(number)

	ageV := 16.45
	fmt.Println(ageV)

	name := "Dima"
	fmt.Println(name)
	fmt.Println(len(name))

	var nameScan string
	var ageScan uint
	fmt.Println("What is your name?")
	fmt.Scan(&nameScan)
	fmt.Println("Hello " + nameScan + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&ageScan)
	fmt.Println("You are " + fmt.Sprint(ageScan) + " years!")

	var tempH int8 = 2
	var tempG float64 = float64(tempH)
	fmt.Println(tempG)

	num := 3
	if num > 0 {
		fmt.Println("Number " + fmt.Sprint(num) + " is greater then 0")
	} else if num < 0 {
		fmt.Println("Number " + fmt.Sprint(num) + " is less then 0")
	} else {
		fmt.Println("Number " + fmt.Sprint(num) + " equals 3")
	}
}
