package main

import (
	"fmt"
	"./controllers"
	"./models"
	)

var session *r.Session

func initDb() *r.Session {
	var err error
	session, err := r.Connect(r.ConnectOpts{
	Address: "192.168.0.2:28015",
	})

	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
		return session
	}

	res, err := r.Expr("Hello World").Run(session)
	if err != nil {
		log.Fatalln(">>>>>>WSW>>>>")
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
	res.Close()

	return session

}
func main() {
	createTable(initDb())

	router,Global := gin.Default()

	v1 := router.Group("api")
	{
		v1.POST("/users", controllers.PostUser)
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/users/:id", controllers.GetUser)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}
	
	router.Run(":1337")
}

func createTable(session *r.Session){
	result, err := r.TableCreate("users").RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}
	printStr("*** Create table result: ***")
	printObj(result)
	printStr("\n")
}

func printStr(v string) {
	fmt.Println(v)
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}

