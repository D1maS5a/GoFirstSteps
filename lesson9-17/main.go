package main

import (
	"fmt"
	"math"
	"sort"
)

/*

 */

func first() {
	fmt.Println("Hello from first func")
}

func mySum(a int, b int) int {
	// fmt.Println(a + b)
	return a + b
}

func mathFunc(a int, b int) (sum int, sub int, mul int, div int) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return
}

func test(someFunc func(int) int) {
	fmt.Println(someFunc(25))
}

func testStr(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	// fmt.Println("New File!")
	// var names [3]string = {}
	// names[0] = "Jordan"
	// names[1] = "Emma"
	// names[2] = "Kate"

	names := [3]string{"John", "Kate", "Emma"}
	fmt.Println(names, len(names))

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	marks := [5]float64{11, 10, 12, 9, 11}

	var sum float64 = 0
	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}
	fmt.Println(sum)

	var result float64 = sum / float64(len(marks))
	fmt.Println(math.Round(result))

	array := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	fmt.Println(array)
	fmt.Println(array[1][2])

	slice := []int{3, 1, 2, 5, 7, 4}
	slice = append(slice, 0)
	slice[0] = 11
	sort.Ints(slice)
	fmt.Println(slice)

	sliceStr := []string{"v", "b", "a", "c", "d", "z"}
	sort.Strings(sliceStr)
	fmt.Println(sliceStr)

	sliceInt := []int{1, 4, 5}

	for _, el := range sliceInt {
		fmt.Printf("%d\n", el)
	}

	age := 30
	name := "John"
	value := 1000.476
	aaa := true
	fmt.Printf("My age is %d. My name is %s. I have %f. %t \n", age, name, value, aaa)

	// var money map[string]int = map[string]int{
	money := map[string]int{
		"dollars": 1000,
		"euros":   2000,
		"apples":  3,
	}

	money["dollars"] = 5000
	delete(money, "apples")
	// el, status := money["dollars"]
	el, status := money["ddd"]
	fmt.Println(el, status)

	fmt.Println(money)
	fmt.Println(money["dollars"])

	first()
	valueFunc := mySum(5, 9)
	fmt.Println(valueFunc)

	s, su, m, d := mathFunc(3, 5)
	fmt.Println(s, su, m, d)

	funcPerem := func(x int, y int) int {
		return x + y
	}
	sumInt := funcPerem(3, 4)
	fmt.Println(sumInt)

	square := func(x int) int {
		return x * x
	}
	test(square)

	// variable := testStr("Hello")
	// variable()
	testStr("Hello from fn!")()
}
