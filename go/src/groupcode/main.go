// groupcode project main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Id   string `json:"Id"`
	Name string `json: "Name"`
}

type Edetail struct {
	Desg    string `json:"Desg"`
	Address string
	Salary  string
}

var db *gorm.DB

func main() {
	var eemploy []Employee

	var detail []Edetail
	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/anirudh?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//fmt.Println("The connection is opened", db, err)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Edetail{})

	r := gin.Default()

	r.POST("/id", func(c *gin.Context) {

		//Employees := Employee{"1", "Sumalatha"}
		var person []Employee = []Employee{ ///

			Employee{Id: "1", Name: "Sumalatha"},

			Employee{Id: "2", Name: "Sharat"},
			Employee{Id: "3", Name: "Sai"},
			Employee{Id: "4", Name: "Manikanta"},
		}

		var pdetails []Edetail = []Edetail{ //

			//Edetails := Edetail{"software engineer", "Bangalore", "6lakhs"}

			Edetail{Desg: "Developer", Address: "Hyd", Salary: "7lakhs"},
			Edetail{Desg: "Developer", Address: "BGLR", Salary: "8lakhs"},
			Edetail{Desg: "Developer", Address: "DLH", Salary: "9lakhs"},
		}

		db.Create(&person)
		db.Create(&pdetails)

		//db.Where("Id=?", "1").Delete(&Employee{})
	})
	r.GET("/get", func(c *gin.Context) {
		db.Find(&eemploy)
		//db.Find(&detail)
		c.JSON(http.StatusOK, gin.H{"Employee details": eemploy})
		//c.JSON(http.StatusCreated, gin.H{"Details of the employee": detail})
	})

	r.GET("/getdetail", func(c *gin.Context) {

		db.Find(&detail)

		c.JSON(http.StatusCreated, gin.H{"Details of the employee": detail})
	})
	r.Run(":8085")
}
