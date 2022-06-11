// chan real project main.go
package main

import (
	"fmt"
	"time"
)

func athletics(p chan int) {
	var athleticsgames []string = make([]string, 4)
	athleticsgames[0] = "running"
	athleticsgames[1] = "jumping"
	athleticsgames[2] = "swimming"
	athleticsgames[3] = "cycling"
	for i := 0; i < len(athleticsgames); i++ {
		fmt.Println(athleticsgames[i])
	}
	fmt.Println("end")
}
func school() {
	var strengthofstudents int = 600
	var strengthofteachers int = 60
	var noofrooms int = 80
	var noofnonteaching = 20
	var total int = strengthofteachers + strengthofstudents + noofnonteaching
	//total
	fmt.Println(total, "\n", noofrooms)
}

func main() {
	fmt.Println("start")
	p := make(chan int)

	 school()

	go athletics(p)

	time.Sleep(2)
	
//	p <- 1, 2, 3, 4
	/*
		c := make(chan string, 10)

		go publisher(c)

		go subscriber(c)

		time.Sleep(3)*/
}
/*
//Cloud collab Technologies question

func publisher(c chan string) {
	c <- "List of topics"

	c <- "select topics"
	c <- "Prefae"
	c <- "Index "
	c <- "Contents"
	c <- "Introduction"
	c <- "story"
	c <- "Climax"
	c <- "Summary"
	c <- "The end"
}

func subscriber(c chan string) {
	for i := 1; i < 10; i++ {
		if i == 5 {
			fmt.Println(i, <-c)
		}
	}
}
*/