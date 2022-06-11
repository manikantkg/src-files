// RTC rest api project main.go
package main

import (
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"

	"fmt"

	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

type Rtc struct {
	Title     string `json: "title"`
	Completed int    `json: "completed"`
	Name      string `json: "name"`
}

type rtc []Rtc


func empRtc(w http.ResponseWriter, r *http.Request) {
		articles := Articles{

		Article{Title: "Test Title", Desc: "Test Description", Content: "Manikanta"},

		Article{Title: "Exam Timetable", Desc: "Schedule", Content: "Final Round"},
	}
	fmt.Println("End point Hit : All Articles End point")
	json.NewEncoder(w).Encode(articles)
}
	}

func main() {
	
	dsn := "root:ap02BL1426*@tcp(localhost:3306)/manikanta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("The connection is opened", db, err)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Rtc{})

	r := gin.Default()

	r.GET("search", func(c *gin.Context) {

	})
	r.Run()

}
