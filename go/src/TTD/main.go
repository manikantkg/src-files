// TTD project main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	phoneno string `form:"phone"`
	amount  int    `form:"amt"`
}

func main() {
	route := gin.Default()
	route.POST("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.phoneno)
		log.Println(person.amount)
	}
	c.String(200, "Success")
}

/*
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TTD struct {
	Date string `form:"date"`
}

const (
	layoutISO = "2021-02-10"
)

func checkdate(c *gin.Context) {

	var ttd TTD

	if err := c.ShouldBind(&ttd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, _ := time.Parse(layoutISO, ttd.Date)

	fmt.Println(value.Day(), value.Month(), value.Year())

	if value.Day() != 10 && value.Day() != 16 && value.Day() != 24 && value.Day() != 23 && value.Day() != 9 {

		c.JSON(http.StatusUnauthorized, gin.H{"status": "Tickets Not Available for this day"})
		return

	}

	if value.Month() != 2 && value.Month() != 1 && value.Month() != 3 {

		c.JSON(http.StatusNotFound, gin.H{"status": "Tickets Not Open for this month"})
		return
	}

	if value.Year() != 2021 {

		c.JSON(http.StatusNotFound, gin.H{"status": "Not a valid year"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Number of Tickets available": 20, "Time Slot":

}

func main() {
	fmt.Println("Hello World!")
}
*/
