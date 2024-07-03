package main

import "fmt"

func main() {
	x := 2
	switch x {
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
}
