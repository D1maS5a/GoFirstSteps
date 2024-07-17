package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup")
	res := m.Run()
	fmt.Println("tear-down")
	os.Exit(res)
}

func main() {

}

func Multiple(x, y int) int {
	return x * y
}

func Add(x, y int) int {
	return x + y
}
