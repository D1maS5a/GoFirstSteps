package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println("Call pointer")
}

func main() {
	pointsMap := map[string]Point{
		"b": {
			Y: 13,
			X: 15,
		},
	}
	otherPointsMap := make(map[int]Point)

	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])
	otherPointsMap[1] = Point{1, 2}
	fmt.Println(otherPointsMap)
	fmt.Println(otherPointsMap[1])

	var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{
			"a": {1, 2},
		}
		fmt.Println("oneMorePointsMap is nil")
	} else {
		oneMorePointsMap["a"] = Point{1, 2}
	}
	fmt.Println(oneMorePointsMap["a"].X)
	fmt.Println(oneMorePointsMap["a"].Y)
	oneMorePointsMap["a"].method()

	key := "d"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key = %s exist in map \n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key = %s not in map \n", key)
		fmt.Println(value)
		fmt.Println(Point{})
	}

	pointsMapMy := map[string]int{
		"xx": 123,
		"yy": 345,
	}

	p1 := Point{}
	mapstructure.Decode(pointsMapMy, &p1)
	fmt.Println(p1)

	for k, v := range pointsMapMy {
		fmt.Println(k)
		fmt.Println(v)
	}
}
