package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

type Users []*User

func main() {
	user1 := User{Name: "user1", Age:10}
	user2 := User{Name: "user2", Age:20}
	user3 := User{Name: "user3", Age:30}
	user4 := User{Name: "user4", Age:40}

	users := Users{}

	//アドレスが表示される
	users = append(users,&user1, &user2, &user3, &user4)

	fmt.Println(users)
	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user2, &user3)
	fmt.Println(users2)
	for _, u2 := range users2 {
		fmt.Println(*u2)
	}
}