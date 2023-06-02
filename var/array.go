package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"a", "b", "c"}
	fmt.Println(arr2)

	fmt.Println(arr2[1])

	fmt.Println(len(arr2))

	var sli1 []string = []string{"a", "b","c", "d", "e", "f", "g"}
	fmt.Println(sli1)

	var newSlice = append(sli1, "h")
	fmt.Println(newSlice)

	arr := [...] string{"Golang", "Java"}
    slice := arr[0:1]

    fmt.Println(slice) //=> Golang　
    fmt.Println(len(slice)) //=> 1 //sliceの要素数は１つなので1を返す。
    fmt.Println(cap(slice)) //=> 2 //sliceの生成元の要素数は２つなので2を返す。
}