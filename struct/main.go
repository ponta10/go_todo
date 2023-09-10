package main 

import (
	"fmt"
)

type User struct {
	Name string
	Age int
	// X, Y int
}

// User型の値を引数として受け取り、その値のコピーに対して変更を加えています。この関数は、引数として渡されたオリジナルのUserオブジェクトには影響を与えません。
func updateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// User型のポインタを引数として受け取り、そのポインタが指すオブジェクトに直接変更を加えています。この関数は、引数として渡されたUserオブジェクト自体を変更します。
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

	/* 
	変わる
	関数に渡されるのはUserオブジェクトのアドレス
    */
	updateUser2(user8)

	fmt.Println(user1, user8)

	var p *int
	x := 42
    p = &x
	// これはメモリ上のxのアドレスを出力します
	fmt.Println(p)
	// これはポインタpが指すアドレスに格納されている値
	fmt.Println(*p)
}