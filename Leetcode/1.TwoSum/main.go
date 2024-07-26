package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))
	fmt.Println(twoSumMap1([]int{2, 5, 5, 11}, 10))
	fmt.Println(twoSumMap2([]int{2, 5, 5, 11}, 10))
}

// My anwser
func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)
	for i, num := range nums {
		for j, num2 := range nums { // можно улучшить время тупо убрав continue и изменить цикл на j := i + 1
			if j == i {
				continue
			}
			if num+num2 == target {
				res = append(res, i, j)
				return res // return []int{i, j}
			}
		}
	}
	return res
}

func twoSumMap1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
	}
	return nil
}

func twoSumMap2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
