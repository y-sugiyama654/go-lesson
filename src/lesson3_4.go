package main

import "fmt"

const Pi = 3.141592653
const (
	Username = "test_user"
	Password = "hogehoge"
)

// 変数だとoverflowしてしまう
// var big int = 9223372036854775807 + 1

// 定数だとoverflowせずに代入できる
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(Big - 1)
}
