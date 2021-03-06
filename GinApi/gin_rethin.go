package main


import (
	"fmt"
	"log"
	"encoding/json"
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	)

var session *r.Session

type Users struct {
	Id    int `json:"id"`
	FirstName  string `json:"firstname" binding:"required" form:"firstName" gorethink:"firstName"`
	LastName string `json:"lastname" binding:"required" form:"lastName" gorethink:"lastName"`
	Age int `json:"age" binding:"required" form:"age" gorethink:"age"`
}

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

	router := gin.Default()

	v1 := router.Group("api")
	{
		v1.POST("/users", PostUser)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
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


func GetUsers(c *gin.Context) {
	db := initDb()
	defer db.Close()

	//var users []Users
	cursor, err := r.Table("users").Run(db)
	defer cursor.Close()

	var hero []interface{}
	err = cursor.All(&hero)
	if err != nil { // should I be checking for EOF here and just trying a restart?
		log.Panic("getOneUser 2:",err)
		fmt.Println("nonil")
	}
	c.JSON(200, hero)
	//fmt.Println(hero)

}

func GetUser(c *gin.Context) {
	db := initDb()
	defer db.Close()

	id_temp := c.Params.ByName("id")
    fmt.Println("id:", id_temp)
    id, err := strconv.Atoi(id_temp)
    fmt.Println(reflect.TypeOf(id))

	cursor, err := r.Table("users").Get(id).Run(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	if id != 0 {
		var user Users
		cursor.One(&user)
		c.JSON(200, user)
	}else{
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
	}
}

func PostUser(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var user Users
	err:=c.Bind(&user)
	if err!=nil{
	    fmt.Println(err)
	}
	fmt.Println(user)
	if user.FirstName != "" && user.LastName != "" {
		r.Table("users").Insert(user).Run(db)
		c.JSON(201, gin.H{"success": user})
		return
	}else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
		return
	}

}

func DeleteUser(c *gin.Context) {
    db := initDb()
    defer db.Close()
    fmt.Println("DELETE OP!")
    id_temp := c.Params.ByName("id")
    fmt.Println("id:", id_temp)
    id, err := strconv.Atoi(id_temp)
	fmt.Println(reflect.TypeOf(id))
    cursor, err := r.Table("users").Get(id).Run(db)
    var hero []interface{}
    err = cursor.All(&hero)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("OUT!",err,";;;",hero,";;;",cursor)
    c.JSON(200, hero)
    if id != 0 {
        r.Table("users").Get(id).Delete().Run(db)
        //fmt.Println("IN!",err,";;;",cursor)
        c.JSON(200, gin.H{"success": "User #" + string(id) + " eliminado alv"})
    }else {
      c.JSON(404, gin.H{"error": "No se encuentra el usuario"})
      fmt.Println("=======", id)
    }
}



func UpdateUser( c *gin.Context) {
	db := initDb()
	defer db.Close()

	var user Users
	err:=c.Bind(&user)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(user)
	id := c.Params.ByName("id")
	fmt.Println("=======", id)

	if user.FirstName != "" && user.LastName != ""{
	if id != "" {
		result := Users{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName : user.LastName,
		Age : user.Age,
	}
	r.Table("users").Get(id).Update(user).Run(db)
	c.JSON(200, gin.H{"success": result})
	}else{
		c.JSON(404, gin.H{"error": "No se encontro al usuario"})
	}
	}else{
		c.JSON(422, gin.H{"error": "Campos vacios"})
	}
}

func printStr(v string) {
	fmt.Println(v)
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}
