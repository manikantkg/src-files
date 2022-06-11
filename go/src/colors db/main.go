// colors db project main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	type myForm struct {                         //NOT EXECUTED//
    Colors []string `form:"colors[]"`
}

<form action="/" method="POST">
    <p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>


func (formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
})

}


