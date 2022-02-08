package main

import (
	"fmt"
)

func main() {
	l := []string{"PHP", "Java", "JavaScript"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// lengeを使えばもっと簡単に書ける
	for i, v := range l {
		fmt.Println(i, v)
	}

	// iを使わない場合
	for _, v := range l {
		fmt.Println(v)
	}

	// map(連想配列)の場合
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// キーのみ取り出したい場合
	for k := range m {
		fmt.Println(k)
	}

	// バリューだけ取り出したい場合
	for _, v := range m {
		fmt.Println(v)
	}
}
