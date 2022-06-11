// sirAPI project main.go
package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:ap02BL1426*@/registration?")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user{})
}

func main() {

	router := gin.Default()
	s1 := router.Group("/api")
	{
		s1.POST("/register", registerNewUser)
		s1.GET("/signin/:status", login)
	}

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Run(":8080")
}

type user struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func registerNewUser(c *gin.Context) {

	var userlist user

	if strings.TrimSpace(c.PostForm("username")) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Username cannot be empty"})
		return
	}
	if strings.TrimSpace(c.PostForm("password")) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Password cannot be empty"})
		return
	}
	db.First(&userlist, "username = ?", c.PostForm("username"))

	if userlist.Username == c.PostForm("username") {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User Name already used"})
		return
	}
	data := user{Username: c.PostForm("username"), Password: c.PostForm("password")}
	db.Save(&data)
	c.JSON(http.StatusOK, gin.H{"Message": "Created User successfully"})
	return
}

func login(c *gin.Context) {

	count := rand.Intn(1e4)
	session := sessions.Default(c)

	if c.Param("status") == "login" {
		session.Set("id", count)
		session.Set("email", "test@gmail.com")
		session.Save()
		sessionID := session.Get("user_id")
		if sessionID == nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "not authed"})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{"message": "User Sign In successfully", "session ID": count})
	}
	if c.Param("status") == "logout" {

		session.Clear()
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "User Sign out successfully"})

	}

}
