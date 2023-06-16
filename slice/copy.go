package main 

import (
	"fmt"
)

func main() {
	sl := []int{100, 200}
	sl2 := sl

	//この書き方だとコピー元も上書きされる
	sl2[0] = 1000
	fmt.Println(sl)

	sl3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl4 := make([]int, 9, 10)
	//コピーに成功した数
	n := copy(sl4, sl3)

	fmt.Println(sl3, sl4, n)
}