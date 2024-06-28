package main

import "fmt"

func main() {
	i := 0
LOOP:
	if i > 25 {
		goto END
	}
	fmt.Println(i)
	i++
	goto LOOP
END:
	fmt.Println("End ...")
}
