package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt") // Если не обрабатывать error то _
	defer file.Close()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	data := "Hello\nFrom\nGo\n\tTab"
	file.WriteString(data)

	fileData, _ := os.ReadFile(file.Name())
	fmt.Println(string(fileData))

	file.Close()
	e := os.Remove(file.Name())
	if e != nil {
		fmt.Println("Error", e)
		os.Exit(1)
	}
}
