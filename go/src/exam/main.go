// exam project main.go
package main

import (
	//"net/http"
	//"net/http/httptest"
	//"testing"

	"github.com/gin-gonic/gin"
	//"github.com/stretchr/testify/assert"
)

/*
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

*/

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

/*
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"name":    "manikanta",
			"frnd": "Umesh yadav",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

//"fmt"
//"strings"


func str() {

	A := "aabbbcccab"
	fmt.Println(A)
	strings.Count()
	fmt.Println(A)

}



//Write Code :-
//Input : "aabbbcccab" ,
//Output: "aa2bbb3ccc3a1b1"

//Find Combination of pair elements which give 0,
// Input = [ -1, 0, 1, -2, 2, 3, 4, 5, 2, 3]

func arr() {

	array := []uint{-1, 0, 1, -2, 2, 3, 4, 5, 2, 3}

	for i := 0; i <= len(array); i++ {
		fmt.Println(array[i] + array[i+1])
		//fmt.Println(array[i])
	}

}
func main() {
	//str()
	DisableConsoleColor()
	//arr()
	//fmt.Println("Hello, playground")

}
func DisableConsoleColor() {

}
*/
