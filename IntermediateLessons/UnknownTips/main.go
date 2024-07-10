package main

import (
	"fmt"
)

func main() {
	x := 2
	switch x { // switch i := isTest(a); i{	.. }
	case 1, 2, 3:
		fmt.Println(1, 2, 3)
		fallthrough
	case 10:
		fmt.Println("fall")
	case 4, 5, 6:
		fmt.Println(4, 5, 6)
	default:
		fmt.Println("Null)")
	}

	switch i := getUserRank("User"); i {
	case "User":
		fmt.Println(i)
	case "Admin":
		fmt.Println("Admin")
	}

	iNew := "Admin"
	switch {
	case iNew == "User":
		fmt.Println(iNew)
	case iNew == "Admin":
		fmt.Println("Admin")
	}

	//while
	i := 0
	for i < 10 {
		//..
		i++
	}

	// while (true)
	/*for {
		//..
	}*/

	// Анонимная структура - это структура, которая не содержит имени
	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}
	fmt.Println(Element)

	// Анонимные поля - это те поля, которые не содержат никакого имени
	type student struct {
		int
		string
		float64
	}
	value := student{123, "Bud", 8900.23}
	fmt.Println(value.int, value.string, value.float64)

	//array
	myArray := [...]int{7, 8, 9}
	fmt.Println(len(myArray))

	part1 := []int{1, 2, 3}
	part2 := []int{4, 5, 6}
	combined := append(part1, part2...)
	fmt.Println(combined)

	if j := 5; j < 10 {
		//..
	} else {
		//..
	}

	str := "user"
	if rank := getUserRank(str); rank == "admin" {
		//..
	} else if rank == "manager" {
		//..
	} else {
		//..
	}

	const (
		Online      = iota // 0 iota == serial
		Offline            // 1
		Maintenance        // 2
		Retired            // 3
	)

	// пропуск значений
	const (
		s0 = iota // 0
		_         // 1
		_         // 2
		s3        // 3
	)

	// начать с 3
	const (
		i3 = iota + 3 // 3
		i4            // 4
		i5            // 5
	)

	fmt.Println(North)
	north := North
	fmt.Println(north)

	surrounded := surround("this message", '(', ')')
	fmt.Println(surrounded)

	a, b, c := TwoParReturn()
	fmt.Println(a, b, c)

	sum := 0
	for sum < 1000 { // Типа while
		sum += 10
	}
	fmt.Println(sum)
	// Вечный цикл
	// for{

	// }
}

func getUserRank(rank string) string {
	return rank
}

const (
	North = iota
	East
	South
	West
)

type Direction byte

func (d Direction) Dir() string {
	switch d {
	case North:
		return fmt.Sprintf("North")
	case South:
		return fmt.Sprintf("South")
	case East:
		return fmt.Sprintf("East")
	case West:
		return fmt.Sprintf("West")
	default:
		return "other direction"
	}
}

func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d]
}

func surround(msg string, left, right rune) string {
	return fmt.Sprintf("%c%v%c", left, msg, right)
}

// Именованные выходные переменные делают инициализацию
func TwoParReturn() (a, c, b string) {
	a = "Hello"
	b = "World!"
	c = "all"
	// return a, b, c
	return // можно так изза именных входных но не стоит :)
}
