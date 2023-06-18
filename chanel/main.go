package main

import "fmt"

func main() {
    // var ch1 chan int

    // //受信専用のチャネル
    // var ch2 <-chan int

    // //送信専用
    // var ch3 chan<- int

    ch3 := make(chan int, 5)
    
    ch3 <- 1
    fmt.Println(len(ch3))

    ch3 <- 2
    ch3 <- 3
    fmt.Println(len(ch3))

    //1個ずつデータ量が減っていく
    i := <- ch3
    fmt.Println(i)
    fmt.Println(len(ch3))

    i2 := <- ch3
    fmt.Println(i2)

    //deadlockになる
    ch3 <- 1
    ch3 <- 2
    ch3 <- 3
    ch3 <- 4
    ch3 <- 5
    ch3 <- 6
}