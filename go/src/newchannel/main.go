// newchannel project main.go
/*
package main

import (
	"fmt"
	"sync"
)

func fanin(ag *sync.WaitGroup, in chan string) {

	fmt.Println("FAN IN FUNCTION channel in  :", <-in)

	ag.Done()

}

func fanout(ag *sync.WaitGroup, out chan string) {

	fmt.Println("FAN OUT FUNCTION channel out :", <-out)

	ag.Done()
}

func passchannel(ag *sync.WaitGroup, in chan string, out chan string) {
	fmt.Println("passing channel")
	in <- "login"
	out <- "logout"
	defer ag.Done()
}

func main() {

	fmt.Println("main function starts")
	in := make(chan string)
	out := make(chan string)
	var ag sync.WaitGroup
	ag.Add(3)

	go fanin(&ag, in)
	go fanout(&ag, out)

	go passchannel(&ag, in, out)

	ag.Wait()

	fmt.Println("main ends")
}


package main

import (
	"fmt"
	"time"
)

func fanin(in chan string) {
	fmt.Println("successfully fan in:  ", <-in)
}

func fanout(out chan string) {
	fmt.Println("successsfully logout :  ", <-out)
}
func passchannel(in chan string, out chan string) {
	in <- "login"
	out <- "logout"
}

func main() {

	fmt.Println("main function starts")
	in := make(chan string)
	out := make(chan string)

	go fanin(in)
	go fanout(out)

	go passchannel(in, out)
	time.Sleep(3)

	fmt.Println("main ends")
}*/
package main

import (
	"fmt"
	"time"
)

func fanin(in chan string) {
	fmt.Println("successfully fan in:  ", <-in)
}

func fanout(out chan string) {
	fmt.Println("successsfully logout :  ", <-out)
}
func passchannel(in chan string, out chan string) {
	in <- "login"
	out <- "logout"
}

func main() {

	fmt.Println("main function starts")
	in := make(chan string)
	out := make(chan string)

	go fanin(in)
	go fanout(out)

	go passchannel(in, out)
	time.Sleep(3)

	fmt.Println("main ends")
}
