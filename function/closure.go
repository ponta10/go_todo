package main

import "fmt"

//クロージャー
func Later () func(string) string {
	var store string
	return func (next string) string {
		s := store
		store = next
		return s
	}
}

//ジェネレーター(クロジャーの応用)
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
    f := Later()
	fmt.Println(f("hello world"))
	fmt.Println(f("My"))
	fmt.Println(f("Name"))
	fmt.Println(f("Is"))
	fmt.Println(f("Go"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

}