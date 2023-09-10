package main

import "fmt"

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 名前付きの戻り値resultがあり、return文はそのresultの現在の値を返します。
func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func CallFunc(f func()){
	f()
}

func main() {
    i := Plus(1, 2)

	fmt.Println(i)
	i2, i3 := Div(9, 4)
	fmt.Println(i2,i3)

	i4, _ := Div(10, 2)
	fmt.Println(i4)

	i5 := Double(2000)
	fmt.Println(i5)

	Noreturn()

	f := func(x, y int) int {
		return x - y
	}
	i6 := f(3,4)
    fmt.Println(i6)

	i7 := func(x, y int) int {
		return x - y
	}(2,5)
	fmt.Println(i7)

	CallFunc(func () {
		fmt.Println("I'm a function")
	})
}