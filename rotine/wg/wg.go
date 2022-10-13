package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func download() {
	fmt.Println("Task Start")
	time.Sleep(time.Second)
	wg.Done()
	fmt.Println("Task End")
	count += 1
}

var wg sync.WaitGroup
var count int = 0

func WaitGroup() int {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download()
	}
	wg.Wait()
	fmt.Println("All done")
	return count
}
