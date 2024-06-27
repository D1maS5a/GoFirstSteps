package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Text to file")
	e := ioutil.WriteFile("text2.txt", data, 0600)

	if e != nil {
		fmt.Println("Error : ", e)
	}

	file_data, err := ioutil.ReadFile("text2.txt")
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println(string(file_data))

	dataOs := []byte("Text to OS file")
	erOs := os.WriteFile("text3.txt", dataOs, 0600)

	if erOs != nil {
		fmt.Println("Error : ", e)
	}

	fileDataOs, errorMes := os.ReadFile("text3.txt")
	if errorMes != nil {
		fmt.Println("Error : ", errorMes)
	}
	fmt.Println(string(fileDataOs))

	os.Remove("text2.txt")
}
