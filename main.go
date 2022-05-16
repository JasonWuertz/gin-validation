package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type params struct {
	Start *time.Time `json:"start" form:"start" binding:"required_without=End,omitempty,lt|ltfield=End"`
	End   *time.Time `json:"end" form:"end" binding:"required_without=Start,omitempty,gt|gtfield=Start"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		params := params{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
			return
		}
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
	r.Run()
}
