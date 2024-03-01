package main

import (
	"fmt"
	"hcpb-api/db"

	"github.com/gin-gonic/gin"
)

func insertCallHandler(c *gin.Context) {
	var call db.Call
	c.BindJSON(&call)

	fmt.Println(call)
	num_calls, err := db.CountCallsByChatID(call.ChatID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(num_calls)
	if num_calls > 500 {
		c.JSON(400, gin.H{
			"error": "Too many calls",
		})
		return
	}

	err = db.InsertCall(call)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(201, gin.H{
			"message": "Call inserted",
		})
	}
}
