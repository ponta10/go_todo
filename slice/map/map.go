package main 

import (
	"fmt"
)

func main() {
	var m = map[string]int{"a": 100, "b": 200}
	fmt.Println(m)

	m2 := map[string]int{"c": 100, "d": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "japan"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["a"])

	//エラーハンドリング
	s, ok := m4[4]
	if !ok {
		fmt.Println("error")
	} else {
		fmt.Println(s, ok)
	}

	//削除
	delete(m4, 2)
	fmt.Println(m4)

	m5 := map[string]int{
		"Apple": 100,
		"Banana": 200,
	}

	for key,  val := range m5 {
		fmt.Println(key, val)
	}
}