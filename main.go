package main

import (
	"fmt"
	"hcpb-api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	fmt.Println(db.QueryAllCalls())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/call", insertCallHandler)
	r.Run(":7777")
}
