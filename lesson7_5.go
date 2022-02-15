package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	// Something do
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
	// Consumer
	go consumer(ch, &wg)

	wg.Wait()
	close(ch)
}
