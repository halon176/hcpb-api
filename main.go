package main

import (
	"hcpb-api/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/call", insertCallHandler)
	r.GET("/calls", getLastCallsHandler)
	r.Run(":7777")
}
