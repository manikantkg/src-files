// chan project main.go
/*
package main

import (
	"fmt"
	"time"
)

func vehicles() {
	var brands []string = make([]string, 4, 5)
	brands = []string{"hero", "honda", "Bajaj", "Yamaha", "Activa", "duet"}

	/*brands[0] = "hero"
	brands[1] = "honda"
	brands[2] = "Bajaj"
	brands[3] = "Yamaha"
	brands[4] = "Activa"

	for i, v := range brands {
		fmt.Println(i, v)
	}
	fmt.Println("length is", len(brands), "capacity is ", cap(brands))
}

func servicedates(date string, varient string) {
	if date == "" {
		fmt.Println("HOLIDAY")
	}

	switch date {
	case "10th", "11th", "12th", "13th", "14th":

		break
	default:

		fmt.Println("NO SLOTS AVAILABLE ON THIS DATE")
	}
	switch varient {
	case "Splendor+", "Passion Pro", "Pulsar", "FZ", "Glamour":
		break
	default:
		fmt.Println("THIS VARIENT SERVICE NOT AVAILABLE")
	}

	fmt.Println(date, "of this month", "vehicle is", varient)
}
func main() {

	fmt.Println("Hello World!")

	go vehicles()

	//go servicedates("16th", "Pulsar")

	time.Sleep(3)

	fmt.Println("End")
}

package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go countCat(c)

	for i := 0; i < 5; i++ {
		message := <-c
		fmt.Println(message)
	}
}

func countCat(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "MANIK"
	}
}
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go countCat(c)

	for i := 0; i < 4; i++ {
		message := <-c
		fmt.Println(message)
	}
}

func countCat(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "Cat"
	}
}

/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go countCat(c, &wg)

	for i := 0; i < 4; i++ {
		message := <-c
		fmt.Println(message)
	}
	wg.Wait()
}

func countCat(c chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- "Cat"
	}
	wg.Done()
}
*/
