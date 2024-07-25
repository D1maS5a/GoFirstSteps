package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))
}

func twoSum(nums []int, target int) []int {
	// sum := 0
	res := make([]int, 0, 2)
	for i, num := range nums {
		for j, num2 := range nums {
			// for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			// num2 := nums[j]
			if num+num2 == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}
