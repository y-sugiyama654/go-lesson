package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(610)
	if result == "ok" {
		fmt.Println("Answer!!", "great")
	}

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("grest2")
	}

	num := 9
	if num%2 == 0 {
		fmt.Println("2で割り切れます！")
	} else if num%3 == 0 {
		fmt.Println("3で割り切れます！")
	} else {
		fmt.Println("2で割り切れません！")
	}

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}
