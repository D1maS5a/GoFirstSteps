package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("SAMPLE")

	var newString strings.Builder
	buffer := make([]byte, 4)
	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]
		newString.Write(chunk)
		fmt.Printf("Read %v bytes: %c \n", numBytes, chunk)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("%v\n", newString.String())

	source := strings.NewReader("SAMPLE")
	buffered := bufio.NewReader(source)
	newStringBuf, e := buffered.ReadString('\n')
	if e == io.EOF {
		fmt.Println(newStringBuf)
	} else {
		fmt.Println("Something went wrong...")
	}

	scannerMy := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 5)
	for scannerMy.Scan() {
		lines = append(lines, scannerMy.Text())
	}
	if scannerMy.Err() != nil {
		fmt.Println(scannerMy.Err())
	}
	fmt.Printf("Line  count: %v\n", lines)
	for _, line := range lines {
		fmt.Printf("Line: %v \n", line)
	}

	buf := bytes.NewBufferString("")
	numBytesMy, errorMy := buf.WriteString("SaMPLE")
	if errorMy != nil {
		fmt.Println(errorMy)
	} else {
		fmt.Printf("Wrote %v bytes: %c\n", numBytesMy, buf)
	}
}
