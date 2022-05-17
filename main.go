package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type params struct {
	Start *time.Time `form:"start" binding:"required_without=End,omitempty,lt|ltfield=End"`
	End   *time.Time `form:"end" binding:"required_without=Start,omitempty,gt|gtfield=Start"`
}

func main() {
	r := gin.Default()

	// field names returned by the validation engine should use hte same name as the JSON property name
	// so that from the API callers perspective property names are what they'd expect them to be.
	// For example, without this the field Start would also be referred to as 'Start' rather than the lowercase 'start'.
	// This may not seem significant with the above example, but as your API contract and your data model diverge or you
	// chose simpler naming for your JSON fields this will prevent a lot of confusion.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("form"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	r.GET("/", func(c *gin.Context) {
		params := params{}
		// Try binding the query parameters to the parameter struct and perform validation
		if err := c.ShouldBind(&params); err != nil {
			// if there are errors pass them to our parsing helper method
			c.JSON(http.StatusBadRequest, gin.H{"errors": parseError(err)})
			return
		}
		// Assuming there were no errors we can perform our actual task!
		now := time.Now()
		if params.Start == nil {
			params.Start = &now
		}
		if params.End == nil {
			params.End = &now
		}
		c.JSON(http.StatusOK, fmt.Sprintf("%v", params.End.Sub(*params.Start)))
		return
	})
	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
