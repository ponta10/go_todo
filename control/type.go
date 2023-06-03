package main

import "fmt"

func anything(a interface{}) {
	fmt.Println(a)
	switch v := a.(type) {
		case string:
			fmt.Println(v + "17")
		case int:
			fmt.Println(v + 10000)
	}
}

func main() {
	//型スイッチ

	anything("aaaa")
	anything(1)

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 2)
	// fmt.Println(x + 1)

	//変換に失敗するとfalseがかえる
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	//if文
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x isint")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don(t know)")
	}

	//switch
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know)")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "is a bool")
	case int:
		fmt.Println(v, "is an int")
	case string:
		fmt.Println(v, "is a string")
	default:
		fmt.Println(v, "is an unknown type")
}
}