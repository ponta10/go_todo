package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("Current time:", now)

    // time.Hourは時間のデュレーションを表す定数です。これを48倍することで、48時間を表現しています。now.Add関数は、指定されたデュレーションを現在の時刻に追加して新しい時刻を返します。
    future := now.Add(48 * time.Hour)
    fmt.Println("48 hours later:", future)
}