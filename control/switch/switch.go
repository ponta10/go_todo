package main

import "fmt"

func main() {
	n := 5
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	switch n1 := 2; n1 {
		case 1,2:
			fmt.Println("1 or 2")
		case 3,4:
			fmt.Println("3 or 4")
	}

	n3 := 6
	switch {
	case n3 > 0 && n3 < 4:
		fmt.Println("0 < n < 4")
	case n3 > 3 && n3 < 7:
		fmt.Println("3 < n < 7")
	}

	//列挙型と演算子の混合パターンはエラーが起きる

}