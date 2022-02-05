package main

import (
	"fmt"
)

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 配列はサイズを変更できない
	var b [2]int = [2]int{300, 400}
	fmt.Println(b)
}
