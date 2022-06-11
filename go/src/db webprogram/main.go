// db webprogram project main.go
package main

import (
	"fmt"
	"log"

	//"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"*/

type UserModel struct {
	Id      int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func main() {
	dsn := "root:ap02BL1426*@tcp(localhost:3306)/manikanta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//fmt.Println("The connection is opened", db, err)
	if err != nil {
		panic("failed to connect database")
	}

	/*
		db, err := gorm.Open("mysql", "root:ap02BL1426*@tcp(localhost:3306)/manikanta?charset=utf8&parseTime=True")
		if err != nil {
			log.Panic(err)
		}
		log.Println("Connection Established") */
	// db.DropTableIfExists(&UserModel{})
	db.AutoMigrate(&UserModel{})

	user := &UserModel{Name: "John", Address: "New York"}
	newUser := &UserModel{Name: "Martin", Address: "Los Angeles"}
	fmt.Println(newUser, user)
	log.Println("end of code")

	/*
		//To insert or create the record.
		//NOTE: we can insert multiple records too
		db.Debug().Create(user)
		//Also we can use save that will return primary key
		db.Debug().Save(newUser)

		//Update Record
		db.Debug().Find(&user).Update("address", "California")
		//It will update John's address to California

		// Select, edit, and save
		db.Debug().Find(&user)
		user.Address = "Brisbane"
		db.Debug().Save(&user)

		// Update with column names, not attribute names
		db.Debug().Model(&user).Update("Name", "Jack")

		db.Debug().Model(&user).Updates(
			map[string]interface{}{
				"Name":    "Amy",
				"Address": "Boston",
			})

		// UpdateColumn()
		db.Debug().Model(&user).UpdateColumn("Address", "Phoenix")
		db.Debug().Model(&user).UpdateColumns(
			map[string]interface{}{
				"Name":    "Taylor",
				"Address": "Houston",
			})
		// Using Find()
		db.Debug().Find(&user).Update("Address", "San Diego")

		// Batch Update
		db.Debug().Table("user_models").Where("address = ?", "california").Update("name", "Walker")

		// Select records and delete it
		db.Debug().Table("user_models").Where("address= ?", "San Diego").Delete(&UserModel{})

	*/
	/*
		db.Debug().Where("address = ?", "Los Angeles").First(&user)
		log.Println(user)
		db.Debug().Where("address = ?", "Los Angeles").Find(&user)
		log.Println(user)
		db.Debug().Where("address <> ?", "New York").Find(&user)
		log.Println(user)
		// IN
		db.Debug().Where("name in (?)", []string{"John", "Martin"}).Find(&user)
		log.Println(user)
		// LIKE
		db.Debug().Where("name LIKE ?", "%ti%").Find(&user)
		log.Println(user)*/
	// AND
	/*	db.Debug().Where("name = ? AND address >= ?", "Martin", "Los Angeles").Find(&user)
		log.Println(user)

			//Find the record and delete it
			db.Where("address=?", "Los Angeles").Delete(&UserModel{})

				// Select all records from a model and delete all
				db.Debug().Model(&UserModel{}).Delete(&UserModel{})
	*/
}
