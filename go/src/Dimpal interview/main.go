// Dimpal interview project main.go
package main

import "sync"

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	go feeder(ch, &wg)
	go catcher(ch, &wg)
	close(ch)
	wg.Wait()
}

func catcher(ch <-chan int, wg *sync.WaitGroup) {
	for {
		foo, ok := <-ch
		if !ok {
			println("done")
			return
		}
		println(foo)
	}
}

func feeder(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		ch <- i
	}
}

/*

Q) What data does Gin internal context contains?
Q) How to copy data from one map to another?


*/
