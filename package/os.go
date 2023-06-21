package main 

import (
	"fmt"
	"os"
	"log"
)

func main() {
	//任意の時点で終了
	// os.Exit(1)
	fmt.Println("Starting")


	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}
}