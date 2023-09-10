package main 

import (
	"fmt"
)

func Sum(s ...int) int {
	n := 0
	for _, value := range s {
		n += value
	}
	return n
}

func main() {
	sl := []string{"A", "B", "C", "D", "E", "F"}

	for _, value := range sl {
		fmt.Println(value)
	}

	fmt.Println(Sum(1, 2, 5))

	sl2 := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(sl2...))
}