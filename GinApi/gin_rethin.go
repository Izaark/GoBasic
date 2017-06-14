package main

import (
    "fmt"
    "log"
    "encoding/json"
    r "gopkg.in/gorethink/gorethink.v3"
    "github.com/gin-gonic/gin"
    //"reflect"
)

var session *r.Session

type Users struct {
    Id    int `json:"id" binding:"required" form:"id" gorethink:"id"`
   FirstName  string `json:"firstname" binding:"required" form:"firstName" gorethink:"firstName"`
   LastName string `json:"lastname" binding:"required" form:"lastName" gorethink:"lastName"`
   Age int `json:"age" binding:"required" form:"age" gorethink:"age"`
}

func initDb() *r.Session {
    var err error
    session, err := r.Connect(r.ConnectOpts{
        Address: "192.168.0.4:28015",
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
        v1.DELETE("/users/:id", DeleteUser)
    }

   router.Run(":1337")
}

func createTable(session *r.Session) {
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
    } else {
        c.JSON(422, gin.H{"error": "Fields are empty"})
        return
    }

}

func DeleteUser(c *gin.Context) {
    db := initDb()
    defer db.Close()

    id := c.Params.ByName("id")
    fmt.Println("=======", id)
    if id != id {
      r.Table("users").Get(id).Delete().Run(db)
        c.JSON(200, gin.H{"success": "User #" + id + " eliminado alv"})
    }else {
      c.JSON(404, gin.H{"error": "No se encuentra el usuario"})
      fmt.Println("=======", id)
}
}


func printStr(v string) {
    fmt.Println(v)
}

func printObj(v interface{}) {
    vBytes, _ := json.Marshal(v)
    fmt.Println(string(vBytes))
}

