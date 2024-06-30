package main

/*Советы
1. Проверять слайс на пустоту с помощью len(list) == 0
2. По возможности, аллоцировать память для слайса
3. Если хотим изменить переданный слайс, создаем его копию
4. Результат append присваивать той же переменной
*/
import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Слайс это
	// pointer: *array - ссылка на первый элемент массива
	// length: int 	   - длинна массива
	// capacity: int   - емкость массива
	// Изначально в массиве может храниться муссор но при append мы перезаписываем в последний индекс длины значение
	// len(12) cap(12) -> делаем слайс [4:7] и будет с 4 по 6 элемент len(3) cap(9) -> взяли слайс [6:9] и будет с 6 по 8 элемент len(3) cap(7)

	/*
		// 1 совет
		var list []int
		fmt.Println(list == nil)
		fmt.Println(len(list) == 0)

		list = []int{}
		fmt.Println(list == nil)
		fmt.Println(len(list) == 0)

		list = []int{1, 2, 3, 4, 8}
		fmt.Println(double(list))

		list2 := []int{1, 2, 3, 4}

		fmt.Println("before", list2)
		handle(list2[:1])
		fmt.Println("after", list2)
	*/

	/*
		// 4 совет
		list := make([]int, 4, 5)
		list2 := append(list, 1)
		list[0] = 5
		list2[0] = 9
		fmt.Println(list)
		fmt.Println(list2)
	*/

}

func double(nums []int) []int {
	// 2 совет
	// res := make([]int, 0, len(nums))
	res := make([]int, len(nums)) // в таком случае append делать нельзя так как массив тупо увеличится в два раза т.к будут вписаны в массив дефолтные значения
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

func handle(list []int) {
	// 3 совет
	newList := make([]int, len(list))
	copy(newList, list)
	newList = append(newList, 5)
	fmt.Println("append", newList)

	// list[1] = 123

	// [1], 2, 3, 4
	// len == 1
	// append
	// [1, 2], 3, 4 -> 5 [1, 5], 3, 4
	// len == 2

	// list = append(list, 5)
	// fmt.Println("append", list)
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)

	b = digitRegexp.Find(b)
	res := make([]byte, len(b))
	copy(res, b)
	//append

	return res
}
