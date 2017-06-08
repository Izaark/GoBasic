package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
) 

func main() {
	//endpoint()
	rout()

}// listen and serve on 0.0.0.0:8080


func rout() {
	router := gin.Default()
	router.POST("/post", user)
	router.Run(":8000")

}

func user(c *gin.Context) {

		id := c.Query("id")
		name := c.Query("name")
		last_name := c.Query("last_name")
		age := c.DefaultQuery("age", "20")

		c.JSON(200, gin.H{
			"id": id,
			"name": name,
			"last_name": last_name,
			"age":age,
		})
		fmt.Printf("id: %s; name: %s; last_name: %s; age: %s; ", id, name, last_name,age)
	}


func endpoint(){
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", )
		v1.POST("/submit", )
		v1.POST("/read", )
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", )
		v2.POST("/submit", )
		v2.POST("/read", )
	}

	router.Run(":8000")
}