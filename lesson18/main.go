package main

import (
	"fmt"
)

func change(str *string) {
	*str = "LOL"
}

type User struct {
	name     string
	age      int64
	password string
	score    []int
}

func changeStruct(u *User) {
	u.name = "Kate"
	u.age = 15
	u.password = "newPass"
}

func (u User) getName() string {
	return u.name
}

func (u *User) setName(newName string) {
	u.name = newName
}

func (u User) isElder() bool {
	if u.age < 18 {
		return false
	}
	return true
}

func (u *User) getHighScore() int {
	// high := u.score
	// sort.Ints(high)
	// return high[len(high)-1]

	// sort.Ints(u.score)
	// fmt.Println(u.score)
	// return u.score[len(u.score)-1]

	high := 0
	for _, sc := range u.score {
		if high < sc {
			high = sc
		}
	}
	return high
}

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiply() int
	Division() float64
	Substract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}
func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}
func (n Numbers) Division() float64 {
	return float64(n.num1) / float64(n.num2)
}
func (n Numbers) Substract() int {
	return n.num1 - n.num2
}

func main() {
	a := 5
	pointer := &a
	fmt.Println(&pointer)
	fmt.Println(*pointer)
	// fmt.Scan(&a)

	s := "LLL"
	var point *string = &s
	fmt.Println(s)
	fmt.Println(point)
	fmt.Println(&s)
	change(point)
	fmt.Println(s)

	num := 9
	b := &num
	fmt.Println(num, b)
	fmt.Println(*b)
	*b = 35
	fmt.Println(*b)

	////////
	// var user User = User{name: "John", age: 23, password: "pass"}
	user := User{"John", 23, "pass", []int{23, 67, 89, 102, 245, 11}}
	fmt.Println(user)
	user.name = "Denis"
	fmt.Println(user.name)
	changeStruct(&user)
	fmt.Println(user.name)
	fmt.Println(&user)

	user.setName("Den")
	fmt.Println(user.getName())
	user.age = 18
	if user.isElder() {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("Come back when you grow up")
	}
	fmt.Println(user.isElder())
	fmt.Println(user.getHighScore())
	// fmt.Println(user.score)

	var i NumbersInterface
	nu := Numbers{9, 8}
	i = nu
	fmt.Printf("Sum: %d \n", i.Sum())
	fmt.Printf("Sum: %f \n", i.Division())
	fmt.Printf("Sum: %d \n", i.Multiply())
	fmt.Printf("Sum: %d \n", i.Substract())
}
