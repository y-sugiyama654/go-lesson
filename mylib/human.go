package mylib

import "fmt"

// 大文字：Public
// 小文字：private
type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human")
}
