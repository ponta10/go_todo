package main

import "fmt"

const Pi = 3.14

const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

func main() {
	fmt.Println(Pi)

	fmt.Println(SiteName, URL)
}