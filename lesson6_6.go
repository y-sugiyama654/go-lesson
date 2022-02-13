package main

import "fmt"

func do(i interface{}) {
	/*
		ii := i.(int)
		i = ii * 2
		fmt.Println(i)

		ss := i.(string)
		i = "Mr." + ss
		fmt.Println(i)
	*/

	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println("Mr." + v)
	default:
		fmt.Printf("I don't know %T\n", v)
	}

}

func main() {
	do(10)
	do("Yuta")
	do(true)
}
