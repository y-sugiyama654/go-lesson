package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Yuta" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get Out!")
	}
}

func main() {
	var yuta Human = &Person{"Yuta"}
	//var dog Dog = Dog{"Pochi"}
	// yuta.Say()
	DriveCar(yuta)
	//DriveCar(dog)
}
