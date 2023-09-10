package main

import "fmt"

type Circle struct {
    Radius float64
}

// 値レシーバを使用するメソッド
// 値そのものを変更することはできない
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// ポインタレシーバを使用するメソッド
// 値そのものを変更することができる
func (c *Circle) DoubleRadius() {
    c.Radius *= 2
}

func main() {
	// Radius: 10のオブジェクトを生成
	c := Circle{Radius: 10}
	fmt.Println(c) // {10}
	c.DoubleRadius()
	fmt.Println(c) // {20}
	println(c.Radius) // 2.000000e+001 (20)
	println(c.Area()) // +1.256000e+003 (1256  20 * 20 * 3.14)
}