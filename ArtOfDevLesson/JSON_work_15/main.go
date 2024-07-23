package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocked bool   `json:"is_blocked"`
	Books     []Book `json:"books"`
}

type Book struct {
	Name string "json:\"name\""
	Year int    "json:\"year\""
}

func main() {
	var books []Book
	book1 := Book{
		Name: "BN",
		Year: 1990,
	}
	book2 := Book{
		Name: "BN",
		Year: 2001,
	}

	books = append(books, book1, book2)
	// sv := []string{"A", "B", "C"}
	// sv := map[string]interface{}{"field": "value", "field1": 123, "field2": true, "arr": []string{"A", "B", "C"}}
	sv := User{
		Name:      "Dima",
		Age:       80,
		IsBlocked: true,
		Books:     books,
	}
	boolVar, _ := json.Marshal(sv) // "something"
	fmt.Println(string(boolVar))
	// json.Unmarshal()
}

// 15 min
