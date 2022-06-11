// airtel project main.go
package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Chooseplan struct {
	value   int
	details string
}

var db *gorm.DB

func main() {

	var Plans Chooseplan
	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/manikanta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Chooseplan{})

	r := gin.Default()
	
	func pln(w http.ResponseWriter, r *http.Request){
		
	r.POST("/show", func(c *gin.Context) {
		var Plans []Chooseplan = []Chooseplan{
			Chooseplan{value: 99, details: "1 month validity"},
			Chooseplan{value: 199, details: "1 mnth validity with daily 1 gb data"},
			Chooseplan{value: 299, details: "1 month validity with daily 1.5gb data"},
			Chooseplan{value: 399, details: "1 mnth validity with daily  2gb data"},
		}
		db.Create(&Plans)
	})
	}
	r.GET("/showdetails", func(c *gin.Context) {
		db.Find(&Plans)

		c.JSON(http.StatusCreated, gin.H{"Plan details": Plans})

	})
	r.Run(":8086")
}
