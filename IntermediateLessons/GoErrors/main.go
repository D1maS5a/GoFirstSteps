package main

import (
	"errors"
	"fmt"
)

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("Cannot divide by zero")
	} else {
		return rhs / lhs, nil
	}
}

type DivError struct {
	a, b int
}

func (d *DivError) Error() string {
	return fmt.Sprintf("Cannot divide by zero: %d / %d", d.a, d.b)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivError{a, b}
	} else {
		return a / b, nil
	}
}

type UserError struct {
	Msg string
}

func (u *UserError) Error() string {
	return fmt.Sprintf("User error: %v", string(u.Msg))
}

// Пример функции, которая может возвращать ошибку
func div2(a, b int) (int, error) {
	if b == 0 {
		return 0, &UserError{"Cannot divide by zero"}
	}
	return a / b, nil
}

func main() {
	// answer, err := div(9, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("The answer is:", answer)

	// _, e := divide(9, 0)
	// if e != nil {
	// 	var InputError = UserError{"Cannot divide by zero"}
	// 	if errors.Is(e, &InputError) {
	// 		fmt.Println("InpEr", e)
	// 	} else {
	// 		fmt.Println("Other", e)
	// 	}
	// }

	_, e := div2(9, 0)
	if e != nil {
		var inputError *UserError
		if errors.As(e, &inputError) {
			fmt.Println("InpEr", e)
		} else {
			fmt.Println("Other", e)
		}
	}
}
