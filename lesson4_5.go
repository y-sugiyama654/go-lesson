package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {
	/*
		foo()
		// main関数の処理が終わった後に呼び出される
		defer fmt.Println("world")
		fmt.Println("hello")
	*/

	/*
		fmt.Println("run")
		// deferは後から実行される
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/

	file, _ := os.Open("./lesson4_5.go")
	defer file.Close()
	data := make([]byte, 1000)
	file.Read(data)
	fmt.Println(string(data))
}
