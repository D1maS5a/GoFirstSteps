package main

import (
	"fmt"
	"regexp"
)

func IsValudEmail(addr string) bool {
	re, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("Failed to compile regex")
	} else {
		return re.Match([]byte(addr))
	}
}

func main() {
	fmt.Println(IsValudEmail("email@example.com"))
}
