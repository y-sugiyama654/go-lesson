package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// チャンネルに100を追加
	ch <- 100
	fmt.Println(len(ch))
	// チャンネルに200を追加
	ch <- 200
	fmt.Println(len(ch))

	close(ch)

	for c := range ch {
		fmt.Println(c)
	}

	// // チャンネルから100を取り出す
	// x := <-ch
	// fmt.Println(x)
	// fmt.Println(len(ch))
	// // チャンネルに300を追加
	// ch <- 300
	// fmt.Println(len(ch))
}
