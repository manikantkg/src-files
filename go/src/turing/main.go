// turing project main.go
/*
package main

import (
	"fmt"
)

func defFooStart() {
	fmt.Println("defFooStart() executed")
}
func defFooEnd() {
	fmt.Println("defFooEnd() executed")
}
func defMain() {
	fmt.Println("defMain() executed")
}
func foo() {
	fmt.Println("foo() executed")

	defer defFooStart()

	panic("panic from foo()")

	defer defFooEnd()
	fmt.Println("foo() done")
}
func main() {
	fmt.Println("main() started")
	defer defMain()
}


package main

import (
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()

		}()
	}
	wg.Wait()
	Println(len(m))
}


package main

func main() {
	var m := map[string]int  //A
	m["a"] = 1

	if v := m["b"]; v != nil {   //B
		Println(v)
	}
}


package main

import (
	"fmt"
	"runtime"
)

func squares(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}
func main() {
	c := make(chan int, 3)
	go squares(c)

	fmt.Print(runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Print(runtime.NumGoroutine())

}
*/

package solution

import (
	"fmt"
	"sort"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	//A  = []int {-3, 0, 1, 3, 4}
	b := sort.Ints(A)
	fmt.Println(A)
	return 1
}
func main() {
	Solution([]int{-3, 0, 1, 3, 4})
}
