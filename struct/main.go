package main 

import (
	"fmt"
)

type User struct {
	Name string
	Age int
	// X, Y int
}

func updateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func updateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)

	user1.Name = "ponta"
	user1.Age = 21
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)

	user3 := User{Name: "kanta", Age:30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	//エラー
	// user5 := User{30, "user5"}
	// fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	//ポインタ型(関数の引数)
	user8 := &User{}
	fmt.Println(user8)

	//変わらない→コピーに対して変更
	updateUser(user1)
	//変わる
	updateUser2(user8)

	fmt.Println(user1, user8)
}