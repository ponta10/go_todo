package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main() {
	m := map[int]User{
		1: {Name: "John", Age: 10},
		2: {Name: "Bob", Age: 20},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{Name: "John", Age: 10}: "USA",
		{Name: "Bob", Age: 20}: "France",
	}

	fmt.Println(m2)
}