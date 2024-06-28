package main

import (
	"fmt"
	"log"
	"strconv"
)

func printNums(nums ...int) {
	fmt.Println(nums)
}

func squareVal(v int) {
	fmt.Println(&v)
	square := v * v
	v = square
}

func squareAdd(v *int) {
	fmt.Println(v)
	sqr := *v * *v
	*v = sqr
}

type People struct {
	Name    string
	Surname string
	Age     int
}

func (p People) IncAge() { // нужно p *People
	fmt.Printf("%p\n", &p)
	p.Age++
}

func main() {
	// Range: 0 through 255.
	// type uint8 uint8

	floatNum := 10.
	floatNumType2 := float64(10)
	fmt.Println(floatNum)
	fmt.Println(floatNumType2)
	fmt.Println()

	// byte is an alias for uint8
	// type byte = uint8
	name := "masha"
	fmt.Println(name[0])
	fmt.Println(string(name[1]))
	fmt.Println(string(name[2]))
	fmt.Println(string(name[3]))
	fmt.Println(string(name[4]))

	// rune is an alias for int32
	// type rune = int32
	symbol := rune('я')
	fmt.Println(symbol)

	x, yStr := 1, "2"
	yStrFloat := "2.5"

	yFloat, er := strconv.ParseFloat(yStrFloat, 64)
	if er != nil {
		log.Fatal(er)
	}
	yInt, err := strconv.Atoi(yStr)
	if err != nil {
		log.Fatal(err)
	}
	z := x + yInt
	m := 1.25 + yFloat
	fmt.Println(z)
	fmt.Println(m)

	printNums(z, x, yInt)

	a := 4
	fmt.Println(&a)
	squareVal(a)
	squareAdd(&a)
	fmt.Println(a)

	// var pointer *int
	// pointer = nil
	// fmt.Println(*pointer)	PANIC

	vasya := People{"Вася", "Пупкин", 20}
	fmt.Printf("%p\n", &vasya)
	vasya.IncAge()
	fmt.Println(vasya)

	//slice
	s := []int{73, 98, 86, 61, 96}
	s[4] = 555
	fmt.Println(len(s), cap(s))
	// s[5] = 999 	PANIC	index out of range [5]
	s = append(s, 999)
	fmt.Println(len(s), cap(s))

	s2 := make([]int, 5, 500)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	s3 := s[0:2]
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	s3 = s[1:3]
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	s4 := s[:2:2]
	s4 = append(s4, 444)
	fmt.Println(s)
	fmt.Println(len(s4), cap(s4))

	s4 = s[0:2]
	s4 = append(s4, 444)
	fmt.Println(s)
}
