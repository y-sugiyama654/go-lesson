package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	boom := time.Tick(5 * time.Second)

OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom!")
			break OuterLoop
		default:
			fmt.Println(".")
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println("######################")
}
