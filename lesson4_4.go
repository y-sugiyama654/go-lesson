package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac", os)
	case "windows":
		fmt.Println("Windows", os)
	default:
		fmt.Println("Default", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 24:
		fmt.Println("Good Night!")
	default:
		fmt.Println("Morning!")
	}
}
