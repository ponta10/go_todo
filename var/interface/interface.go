package main

import "fmt"

func main() {
	// io.Readerやio.Writer
	//interface 初期値nil 
	var x interface{}
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)
	x = "aaa"
	fmt.Println(x)
	x = [3]int{1,2,3}
	fmt.Println(x)

	//interfaceとintでは計算できない
    // x = 2
	// fmt.Println(x + 3)
}