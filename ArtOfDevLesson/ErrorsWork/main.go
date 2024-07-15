package main

import (
	"fmt"
	"io"
	"strings"
)

// type error interface {
// 	Error() string
// }

type appError struct {
	Err    error
	Custom string
	Field  int
}

// type AppError interface {
// 	Error() string
// 	UnWrap() error
// }

func (e *appError) Error() string {
	// err := fmt.Errorf("new error %s", e.Custom)
	// fmt.Println(e.Custom)
	// return e.Err.Error()
	return fmt.Sprintf("%s error", e.Custom)
}

func (e *appError) UnWrap() error {
	return e.Err
}

func main() {
	// err := m()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	err := m()
	if err != nil {
		// ...
		fmt.Println(err.UnWrap())
		// return
	}
	fmt.Println("success")

	r := strings.NewReader("hello world")
	arr := make([]byte, 8)

	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d err = %v b = %v \n", n, err, arr)
		fmt.Printf("arr n bytes: %q \n", arr[:n])
		if err == io.EOF {
			break
		}
	}
}

func m() *appError /*error*/ {
	// return &AppError{
	// 	Err:    fmt.Errorf("my error"),
	// 	Custom: "value here",
	// 	Field:  89,
	// }
	return m2()
}

func m2() *appError /*error*/ {
	// implementation business logic
	return m3()
}

func m3() *appError /*error*/ {
	// implementation business logic
	// return fmt.Errorf("some error")
	// return nil
	return &appError{
		Err:    fmt.Errorf("internal error"),
		Custom: "something goes wrong",
	}
}

// Stringer - это
type name struct {
	A int
	B string
}

func (n *name) String() string {
	return n.B
}
