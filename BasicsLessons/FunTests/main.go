package main

import "fmt"

func double(nums []int) []int {
	// res := make([]int, 0, len(nums))
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = num * 2
	}

	// var res []int
	// 1
	// 2
	// 4
	// 4
	// 8	5 итераций 4 создания массива с копированием массива
	// for _, num := range nums {
	// 	res = append(res, num)
	// }
	return res
}

func main() {
	var list []int
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0)

	list = []int{}
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0)

	list = []int{1, 2, 3, 4, 8}
	fmt.Println(double(list))
}
