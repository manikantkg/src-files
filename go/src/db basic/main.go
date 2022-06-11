// db basic project main.go
package main

import (
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB

type golang struct {
	Title     string `json:"Title"`
	Completed int    `json:"Completed"`
	Name      string `json:"Name"`
}

func main() {

	dsn := "root:ap02BL1426*@tcp(localhost:3306)/manikanta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the Schema

	db.AutoMigrate(&golang{})

	r := gin.Default()

	r.GET("Json", func(c *gin.Context) {

		//user := golang{Title: "jvt", Completed: 7, Name: " manik"}

		user := golang{Title: "GO", Completed: 2, Name: " m"}

		//result := db.Model(&user).Where("Name=?", "manik").Update("Name", "Sharat") //SHARAT

		//	db.Model(&todo).Update("completed", completed)
		db.Create(&user)
		//	db.Save(&user)

		//db.Where(&user)

		//result1 := db.Model(&user).Where("Name = ?", true).Update("name", "hello") //SAI LEELA

		//if result != nil {

	})

	r.Run(":8081")
}
