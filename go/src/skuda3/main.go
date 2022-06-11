// skuda3 project main.go
package main

import (
	//"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	id              int    `json:"id"`
	name            string `json:"name"`
	options         string `json:"options"`
	correct_options string `json:"correct_options"`
	quiz            string `json:"quiz"`
	points          int `json:"points"`
}

var db *gorm.DB

func main() {
	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/manikanta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var details Person
	r := gin.Default()

	r.GET("/post", func(c *gin.Context) {
		db.Save(&details)
		details{"id":1, name: "go", options: "language", correct_options: "correct",quiz:"quiz",points:100},

		//c.String(200, "Hello, World!")
		c.JSON(http.StatusCreated, gin.H{"id": details})
	})
	db.Find(&details)
	db.AutoMigrate(&details)
	r.Run(":8080")
}

/*
{
    "id": ,                     ?--------- number
    "name": "",
    "options": "",
    "correct_option": ,
    "quiz": ,
    "points": ,
} */
