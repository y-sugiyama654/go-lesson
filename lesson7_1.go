package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		//time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go goroutine("World!", &wg)
	normal("Hello")

	wg.Wait()
}
