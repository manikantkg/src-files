// mapdatacopy project main.go
/*
package main

import "fmt"

func main() {

	// Creating and initializing a map
	// Using shorthand declaration and
	// using map literals
	originalMap := make(map[string]int)
	originalMap["one"] = 1
	originalMap["two"] = 2
	originalMap["three"] = 3
	originalMap["four"] = 4
	originalMap["five"] = 5
	originalMap["six"] = 6
	originalMap["seven"] = 7
	originalMap["eight"] = 8
	originalMap["nine"] = 9

	// Creating empty map
	CopiedMap := make(map[string]int)

	/* Copy Content from Map1 to Map2*/
/*	for index, element := range originalMap {
		CopiedMap[index] = element
	}

	for index, element := range CopiedMap {
		fmt.Println(index, "=>", element)
	}
}


package main

import "fmt"

func main() {

		var name string = "mani"
		for i := 0; i <= 5; i++ {
			fmt.Println("manikanta")
		}

		for i := 0; i <= len(name); i += 2 {
			fmt.Println(name)
		}

	a := make(map[string]int)

	a["ram"] = 2
	a["lakshman"] = 3
	a["sita"] = 4
	a["bharata"] = 5
	b := make(map[string]int)

	for index, value := range a {
		b[index] = value

	}
	for index, value := range b {
		fmt.Println(index, value)

	}
}



package main
import "fmt"
type AA interface {
	name string
	age int
}

type team struct {
	game string
	players int
}

func
func main(){

}

//wait group
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go count("sheep")
	wg.Done()

	wg.Wait()
}
func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
	}
	time.Sleep(time.Microsecond * 500)
}


//channel
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}
}
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 600)
	}
	close(c)// ****TO ARAISE DEADLOCK CONDITION****
}
*/
package main

import (
	"fmt"
	//"time"
)

func main() {
	c := make(chan string)
	go count("mani", c)
	msg := <-c
	fmt.Println(msg)

}
func count(thing string, c chan string) {
	//for i := 1; i <= 5; i++ {
	c <- thing
	//time.Sleep(time.Millisecond * 500)
	//}
}

/*
//Pending
package main

import (
	"fmt"
	"time"
)

func main() {
	go name("mani", c)
}
func name(name string, c chan string) {

	c <- "manikanta"
	fmt.Println(name)
}
func age(age int, d chan int) {

	d <- 30
	time.Sleep(10)
}
*/
