package main

import "fmt"

//最初に呼び出される
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}