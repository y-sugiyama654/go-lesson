package main

import "fmt"

func main() {
	b := []byte{89, 85, 84, 65}
	// アスキーコードの72と73が出力される
	// https://www.ascii-code.com/
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("SUGIYAMA")
	fmt.Println(c)
	fmt.Println(string(c))
}
