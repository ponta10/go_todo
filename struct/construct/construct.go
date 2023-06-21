package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

//コンストラクタ関数
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("usee1", 10)

	fmt.Println(user1)
	fmt.Println(*user1)
}