package main

import (
	"fmt"
	"strconv"
)
func main() {
    a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("none")
	}

	if b := 100; a > b {
		fmt.Println("yes")
	} else {
        fmt.Println("no")
	}

	var s string = "100"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", i)
}