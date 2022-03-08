package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "s", "t", "w"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Yuta", 27},
		{"Fumio", 25},
		{"Shinzo", 60},
	}
	fmt.Println(i, s, p)

	// intのソート
	sort.Ints(i)
	fmt.Println(i)

	// stringのソート
	sort.Strings(s)
	fmt.Println(s)

	// pのNameのアルファベット順にソート
	sort.Slice(p, func(i, j int) bool {
		return p[i].Name < p[j].Name
	})
	fmt.Println(p)
}
