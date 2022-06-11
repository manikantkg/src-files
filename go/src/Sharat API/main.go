// Sharat API project main.go
package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Lab struct {
	//gorm.Model
	ID            string `"primary_key"` //`json: id`
	Labname       string
	Certification string

	City     string
	Location string
	tests    []test `gorm:"ForeignKey:Custid"`
}

type test struct {
	gorm.Model
	Testname      string `primary_key`
	Totaltestdone int
	Rating        int
	Price         int
	Custid        int `gorm:"ForeignKey:CustId"`
}

var db *gorm.DB

func main() {

	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/manikanta?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect data base")

	}
	r := gin.Default()
	var patient = []Lab{
		{ID: "165", Labname: "vijaya diagnostic", Certification: "ISO", City: "hyderabad", Location: "lb nagar"},
		{ID: "166", Labname: "likitha diagnostic", Certification: "ISO", City: "hyderabad", Location: "chaithanyapuri"},
		{ID: "167", Labname: "medquest diagnostic", Certification: "ISO", City: "hyderabad", Location: "kothapet"},
		{ID: "168", Labname: "metro diagnostic", Certification: "ISO", City: "hyderabad", Location: "dilsukhnagar"},
	}
	r.GET("/test/:id", func(c *gin.Context) {
		db.Find(&patient)

		id := c.Param("id")
		for _, a := range patient {
			if a.ID == id {
				c.JSON(http.StatusOK, a)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "test not found"})
	})
	r.Run()
}
