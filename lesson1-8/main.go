package main

import (
	"fmt"
	"math"
)

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
	if num > 0 { // >= <=
		fmt.Println("Number " + fmt.Sprint(num) + " is greater then 0")
	} else if num < 0 {
		fmt.Println("Number " + fmt.Sprint(num) + " is less then 0")
	} else {
		fmt.Println("Number " + fmt.Sprint(num) + " equals 3")
	}

	fmt.Print("Квадратное уравнение. A = ")
	var a float64
	var b float64
	var c float64
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)
	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("\nИмеет два корня\nD = " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("\nИмеет один корень\nD = " + fmt.Sprint(D))
		fmt.Println("X: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Уравнение не имеет корней \nD = " + fmt.Sprint(D))
	}

	// fmt.Scan(&a)
	// go build -ldflags "-s -w" main.go  --- Уменьшение веса exe-файла

	for i := 0; i < 5; i++ {
		fmt.Printf("Hello %d\n", i)
	}
	fmt.Println("End")

	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		if i == 50 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("End")

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("End")

	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println("End")

	for index, element := range nums {
		fmt.Printf("Index: %d Element: %d\n", index, element)
	}
	fmt.Println("End")

	for _, element := range nums {
		fmt.Printf("Index: %d\n", element)
	}
	fmt.Println("End")

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

	namePerson := "John"

	switch namePerson {
	case "Jordan":
		fmt.Println("Hello Jordan")
	case "Kate":
		fmt.Println("Hello Kate")
	case "John":
		fmt.Println("Hello John")
	default:
		fmt.Println("I don't know you")
	}

	numberSwitch := 10
	switch {
	case numberSwitch > 1:
		fmt.Println("Number is greater than 1")
		// break
		fallthrough
	case numberSwitch < 11:
		fmt.Println("Number < 11")
	}

	aTmp := 3
	bTmp := 10
	if aTmp > 1 && bTmp > 5 {
		fmt.Println("True")
	}

	if aTmp > 1 || bTmp > 9 {
		fmt.Println("True")
	}

	if aTmp != 2 && b > 5 || a > 6 {
		fmt.Println("Yes")
	}

	// + - * / %
	var mathPerem float64 = 144
	result := math.Sqrt(mathPerem)
	fmt.Println(result)

	mathPerem = 25.2135
	result = math.Ceil(mathPerem)
	fmt.Println(result)

	mathPerem = 25.9999
	result = math.Floor(mathPerem)
	fmt.Println(result)

	mathPerem = 25.2135
	result = math.Round(mathPerem)
	fmt.Println(result)

}
