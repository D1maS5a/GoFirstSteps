package main

import (
	"fmt"
	"regexp"
)

var EmailExpr *regexp.Regexp

func init() {
	fmt.Println("Init")
	compiled, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("Failed to compile regular expression")
	}
	EmailExpr = compiled
	fmt.Println("Regular expression compiled successfully")
}

func isValidEmail(addr string) bool {
	return EmailExpr.Match([]byte(addr))
}

func main() {
	fmt.Println("First")
	fmt.Println(isValidEmail("invalid"))
	fmt.Println(isValidEmail("valid@example.com"))
	fmt.Println(isValidEmail("invalid@example"))
}
