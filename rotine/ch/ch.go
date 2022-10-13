package ch

import (
	"fmt"
	"time"
)

var ch = make(chan int, 10)

func download(url int) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url
}

var count int = 0

func Down() int {
	for i := 0; i < 3; i++ {
		go download(i)
	}
	for i := 0; i < 3; i++ {
		msg := <-ch
		count++
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
	return count
}
