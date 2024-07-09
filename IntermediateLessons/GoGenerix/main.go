package main

//go mod init main
//go mod tidy
import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Max(1, 5))
	var x, y uint = 2, 6
	fmt.Println(Max(x, y))

	fmt.Println(Max("abc", "zxc"))
	// fmt.Println(quote.Go())

	list := []int{1, 2, 4, 8}
	fmt.Println(IsContains(3, list))
	fmt.Println(IsContains(4, list))
	list2 := []string{"apple", "banana", "orange"}
	fmt.Println(IsContains("apple", list2))
	fmt.Println(IsContains("lalala", list2))

	fmt.Println(Sum(list))

	// sum := func(x, y int) int { return x + y }
	mul := func(x, y int) int { return x * y }
	res := Reduce(list, mul, 1)
	fmt.Println(res)

	// Ternary(cond?, x any, y any)
	// cond ? x : y
	fmt.Println(MaxTer(1, 2))
	fmt.Println(MaxTer(8, 4))

	//Приближение (без ~ работать не будет так как не "чистый" int32)
	type MyInt int32
	nums := []MyInt{MyInt(1), MyInt(2), MyInt(3)}
	fmt.Println(nums)
}

func MaxTer(x, y int) int {
	return TernaryMy(x > y, x, y)
}

func TernaryMy[T any](cond bool, x T, y T) T {
	if cond {
		return x
	}
	return y
}

// Map(list, func) list
// Filter(list,func(x) bool) list
func Reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	// 0 + el1 + el2 + ...
	// 1 * el1 * el2 * ...
	for _, el := range list {
		init = accumulator(init, el)
	}
	return init
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](list []T) T {
	var res T
	for _, n := range list {
		res += n
	}
	return res
}

func IsContains[T comparable](n T, list []T) bool {
	for _, i := range list {
		if i == n {
			return true
		}
	}
	return false
}

// вариант создания своего джинерика
type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

// Ordered
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}

	// comparable
	// if x == y {
	// 	return x
	// }
	return y
}

// Приближение
type Integer32 interface {
	~int32 | ~uint32
}

func SumNumbers[T Integer32](arr []T) T {
	var sum T
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
