package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	//並行処理
    go sub()

	for {
		fmt.Println("Main loop")
		time.Sleep(200 * time.Millisecond)
	}
}