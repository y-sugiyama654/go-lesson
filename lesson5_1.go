package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)

	var a int = 654
	// 代入した数値
	fmt.Println(a)
	// アドレス
	fmt.Println(&a)
	// 代入した数値
	fmt.Println(*&a)
	// アドレス
	fmt.Println(&*&a)

	/*
		var n int = 100
		fmt.Println(n)
		fmt.Println(&n)

		var p *int = &n
		fmt.Println(p)
		fmt.Println(*p)
	*/
}
