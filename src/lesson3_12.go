package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["orange"] = 500
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// panickになる
	// var m3 map[string]int
	// m3["pc"] = 5000
	// fmt.Println(m3)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
