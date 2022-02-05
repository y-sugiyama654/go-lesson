package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	//  %f : 10進記数法
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	//  %d : 整数値
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d\n", i, i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
