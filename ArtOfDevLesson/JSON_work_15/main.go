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
	byt := []byte(`{"name":"Dima","age":80,"is_blocked":true,"books":[{"name":"BN","year":1990},{"name":"BN1","year":2001}]}`)
	// var dat map[string]interface{}
	var dat User
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	// fmt.Println(dat["books"].([]interface{})[1].(map[string]interface{})["name"]) // dat["name"]
	fmt.Println(dat.Books[1].Name)
}

func serialize() {
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
}
