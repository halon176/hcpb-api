package main

import (
	"fmt"
	"hcpb-api/db"
	"log"

	"github.com/gin-gonic/gin"
)

func insertCallHandler(c *gin.Context) {
	var call db.Call
	c.BindJSON(&call)

	fmt.Println(call)
	/*
	num_calls, err := db.CountCallsByChatID(call.ChatID)
	if err != nil {
		log.Println(err)
		c.JSON(500, nil)
		return
	}
	fmt.Println(num_calls)
	*/
	err := db.InsertCall(call)
	if err != nil {
		log.Println(err)
		c.JSON(500, nil)
	} else {
		c.JSON(201, nil)
	}
}

func getLastCallsHandler(c *gin.Context) {
	calls, err := db.QueryLastCalls()
	if err != nil {
		log.Println(err)
		c.JSON(500, nil)
		return
	}
	c.JSON(200, calls)
}
