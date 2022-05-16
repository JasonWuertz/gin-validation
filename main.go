package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type params struct {
	Start *string `json:"start" form:"start" binding:"required_without=End,omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	End   *string `json:"end" form:"end" binding:"required_without=Start,omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		params := params{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("%v", params))
		return
	})
	r.Run()
}
