package main

// mk project main.go
/*package main

import (
	"fmt"


	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type reservation struct {
	gorm.Model
	Date         string `json:"date"`
	Day          string `json:"day"`
	Availability string `json:"availablity"`
	Price        string `json:"price"`
	Book         string
}

func main() {
	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/manikanta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("The connection is opened", db, err)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&reservation{})
	var price_800 = reservation{}
	var price_1200 = reservation{}
	r := gin.Default()
	r.POST("/reserve", func(c *gin.Context) {
		price_800 = reservation{Date: "27/02/2021", Day: "wednesday", Availability: "23 rooms", Price: "800/per room"}
		price_1200 = reservation{Date: "28/02/2021", Day: "thursday", Availability: "24 rooms", Price: "1200/per room"}
		result := db.Create(&price_800)
		result1 := db.Create(&price_1200)
		if result != nil {

		}
		if result1 != nil {

		}
	})
	r.GET("/show", func(c *gin.Context) {

		res := db.Find(&price_800)

		fmt.Println(res)
	})
	r.Run(":8080")
}
*/
/*
type Customer struct {
 FirstName        string `json:"first_name" bson:"first_name"`
 LastName         string `json:"last_name" bson:"last_name"`
 Email            string `json:"email" bson:"email"`
}
type Customers []Customer

type new_user struct {
 first_name     string
 last_name      string
 email          string
}
func main(){
	GetData()
}
function GetData(c *gin.Context){
 first_name := c.PostForm("first_name")
 last_name := c.PostForm("last_name")
 email := c.PostForm("email")
 reqBody := new(new_user)
 err := c.Bind(reqBody)
 if err != nil {
    fmt.Println(err)
 }
 customer.FirstName = first_name
 customer.LastName = last_name
 customer.Email = email
}
*/

//package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.Bind(&person) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	if c.BindJSON(&person) == nil {
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(200, "Success")
}
