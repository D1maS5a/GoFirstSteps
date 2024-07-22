package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// readFile()
	// writeToFile()
	// appendToFile()
	// readFile()

	// ssh -t localhost top
	//c := exec.Command("top") // ну она на винде не запустится XD
	// c := exec.Command("ssh", "-t", "localhost", "top")
	c := exec.Command("cmd.exe", "/C", "type", "text.txt") // type -> cat если Linux (start для запуска уже в блокноте)
	/*
		В Windows для выполнения команды через exec.Command, параметр, который передается в exec.Command, должен включать саму команду и её аргументы.
		Команда type является встроенной командой оболочки cmd.exe, поэтому необходимо вызвать cmd.exe с параметром /C для выполнения встроенной команды.
	*/
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

func appendToFile() {
	f, err := os.OpenFile("appendFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString("- Dima -"); err != nil {
		panic(err)
	}

}

func writeToFile() {
	data := []byte("I am new.")
	err := os.WriteFile("text.txt", data, 0600) // ls -la
	if err != nil {
		fmt.Println(err)
	}

}

func readFile() {
	data, err := os.ReadFile("appendFile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
