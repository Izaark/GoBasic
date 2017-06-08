package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
) 

func main() {
router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe

router.GET("/welcome", guest)
	router.Run(":8080")

}// listen and serve on 0.0.0.0:8080


func Ok(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"message2": "pong2",
			"message3": "pong3",
		})
	}


func User(u *gin.Context) {

	u.JSON(200, gin.H{
		"name:":"isaac",
		"last_name:":"lopez",
		"age:":"22",
		})

}

func guest(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		age := c.Query("age")

		c.JSON(http.StatusOK, gin.H{
			"firstname": firstname,
			"lastname": lastname,
			"age": age,
		})
		//c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	}
