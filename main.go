package main

import (
	"fmt"
	"go-lesson/mylib"
	"go-lesson/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()

	person := mylib.Person{Name: "Sugiyama", Age: 27}
	fmt.Println(person)
}
