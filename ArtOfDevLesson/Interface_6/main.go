package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	var a InterfaceHere
	sh := structHere{1, 2}
	os := otherStruct{2, 3}

	a = &sh
	fmt.Println(a.Sum())

	a = os
	fmt.Println(a.Sum())

	var i InterfaceHere = otherStruct{3, 3}
	fmt.Println(i.Sum())

	// u := NewUserMy("", "", "")
	// u.Block()

	fmt.Printf("(%v, %T)\n", i, i)

	var os1 *structHere
	var inter InterfaceHere
	inter = os1
	fmt.Printf("(%v, %T)\n", inter, inter)

	if inter == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	var aMy interface{}
	aMy = "jelly"
	fmt.Println(aMy)
	fmt.Printf("(%v, %T)\n", aMy, aMy)
	aMy = 42
	fmt.Printf("(%v, %T)\n", aMy, aMy)

	var aNew interface{} = "hello"
	s := aNew.(string)
	fmt.Println(s)

	s, ok := aNew.(string)
	fmt.Println(s, ok)

	f, ok := aNew.(float32)
	if ok {
		// use f
	}
	fmt.Println(f, ok)

	// g := aNew.(float64) // panic
	// fmt.Println(g)

	// type switch (как выше лучше не писать XD)
	var aaa interface{} = 3.14 //true //"hello"
	switch aaa.(type) {
	case int:
		fmt.Println("aaa is int")
	case string:
		fmt.Println("aaa is string")
	case bool:
		fmt.Println("aaa is bool")
	default:
		fmt.Printf("unknown type %T\n", aaa)
	}

	// var json map[string]interface{}
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}
