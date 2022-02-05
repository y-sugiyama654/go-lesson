package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])

	n[2] = 100
	fmt.Println(n)

	var boad = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(boad)

	// スライスは可変式
	n = append(n, 123)
	fmt.Println(n)
}
