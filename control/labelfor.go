package main

import "fmt"

func main() {
	//ラベルつきfor
	Loop: 
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
		}
	}

Looper:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j > 1 {
				continue Looper
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}
}