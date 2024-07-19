package main

import (
	"errors"
	"fmt"
	"log"
)

// type name struct {
// 	A, B int
// }

// func (n *name) method() {
// 	fmt.Println("ok")
// 	fmt.Println(n.A)
// }

type AppError struct {
	Message string
	Err     error
}

func (ae *AppError) Error() string {
	return ae.Message
}

func main() {
	// n := &name{1, 2}
	// n = nil
	// n.method()

	// defer func() {
	// 	fmt.Println("Ok")
	// }()
	// panic("something goes wrong")

	divide(4, 0)
	fmt.Println("after panic")
}

func divide(a, b int) {
	defer func() {
		var appErr *AppError
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("This is App Handle panic", err)
				} else {
					fmt.Println("Other error handle Panic")
				}
			default:
				panic("Default go panic!")
			}
			log.Println("panic happend:", err)
		}
	}()
	fmt.Println(div(a, b))
}

func div(a, b int) int {
	if b == 0 {
		panic(fmt.Errorf("some error"))

		// panic(&AppError{
		// 	Message: "this is divide by zero custom error",
		// 	Err:     nil,
		// })

		// panic("vvv")
	}
	return a / b
}
