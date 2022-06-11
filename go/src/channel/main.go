// channel project main.go
/*
package main

import (
	"fmt"
)

func tracks() {
	fmt.Println("tracks added")
	for i := 0; i < 5; i++ {
		fmt.Println("tracks", i)
	}
}
func main() {

	p := make(chan string)

	go func() {
		p <- "open_VLC, Play, pause, Stop"
	}()
	pla := <-p
	fmt.Println(pla)

	v := make(chan string)
	go func() {
		v <- "Vehicle Service date is tomorrow"
	}()
	date := <-v
	fmt.Println(date)

	b := make(chan int)
	go func() {
		b <- 256
	}()
	values := <-b
	fmt.Println(values)

	//USING MORE CHANNELS IN ONE GOROUTINE

	vlc := make(chan string)
	var songselect chan string = make(chan string)
	var pausesong chan string = make(chan string)
	var volume chan string = make(chan string)
	var tracknumber chan int = make(chan int)
	var stop chan string = make(chan string)

	go func() {
		vlc <- "functioning correctly"
		songselect <- "PLZ choose a song"
		pausesong <- "Pause a song"
		volume <- "Adjust Volume"
		tracknumber <- 54
		stop <- "Stop the music"

	}()

	go tracks()
	app := <-vlc
	fmt.Println(app)
	song := <-songselect
	pause := <-pausesong
	vol := <-volume
	track := <-tracknumber
	stp := <-stop

	fmt.Println(song)
	fmt.Println(pause)
	fmt.Println(vol)

	fmt.Println("tracknumber =", track)
	fmt.Println(stp)

}
*/

// Go program to illustrate
// the concept of Goroutine
package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(5 * time.Second)
		fmt.Println(str)
	}
}

func main() {
	time.Sleep(2)

	go display("Welcome") // Calling Goroutine

	display("GeeksforGeeks") // Calling normal function

}
