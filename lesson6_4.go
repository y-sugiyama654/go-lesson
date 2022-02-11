package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	// fmt.Println("%T %v\n", i, i)
	// fmt.Println("%T %v\n", 1, 1)
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
