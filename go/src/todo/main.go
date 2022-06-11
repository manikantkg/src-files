// todo project main.go

package main

import (
	//"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// open db connection

	var err error

	dsn := "root:ap02BL1426*@tcp(127.0.0.1:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&todomodel{})
}

type todomodel struct {
	gorm.Model

	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// transformedTodo represents a formatted todo
type transformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json : "title"`
	Completed bool   `json : completed`
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{

		v1.POST("/post", CreateTodo)
		v1.GET("/get", fetchAllTodo)
		v1.GET("/:id", fetchsingleTodo)
		v1.PUT("/:id", UpdateTodo)
		v1.DELETE("/:id", deleteTodo)

	}

	router.Run()
}

func CreateTodo(c *gin.Context) {

	completed, _ := strconv.Atoi(c.PostForm("completed"))

	todo := todomodel{Title: c.PostForm("title"), Completed: completed}

	db.Save(&todo)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created"})

}

// fetch all todos

func fetchAllTodo(c *gin.Context) {

	var todos []todomodel

	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {

		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})

		return

	}

	for _, item := range todos {

		completed := false

		if item.Completed == 1 {

			completed = true

		} else {

			completed = false

		}

		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetch singleTodo fetch a single todo
func fetchsingleTodo(c *gin.Context) {

	var todo todomodel

	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {

		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo"})
		return
	}
	completed := false

	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})

}

// updateTodo update a todo

func UpdateTodo(c *gin.Context) {

	var todo todomodel

	todoID := c.Param("id")
	db.First(&todo, todoID)

	if todo.ID == 0 {

		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})

		return
	}
	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))

	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully"})

}

// delete Todo remove a todo

func deleteTodo(c *gin.Context) {

	var todo todomodel

	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no todo found!"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
