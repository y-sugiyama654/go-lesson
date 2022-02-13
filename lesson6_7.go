package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func main() {
	yuta := Person{"Yuta", 27}
	fmt.Println(yuta)
}
