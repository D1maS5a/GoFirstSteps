package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(1111))
	fmt.Println(isPalindromeGPT(12321))
}

// func isPalindrome(x int) bool {
// 	str := strconv.Itoa(x)
// 	count := utf8.RuneCountInString(str)
// 	for i := 0; i <= count; i++ {
// 		for j := -1; j >= ; j-- {
// 			fmt.Println(i, j)
// 		}
// 	}
// 	return false
// }

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}

func isPalindromeGPT(x int) bool {
	// Преобразуем число в строку
	s := strconv.Itoa(x)
	// Используем двойной цикл для сравнения символов с начала и с конца строки
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
