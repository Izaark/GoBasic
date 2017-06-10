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

type Person struct {
    id    int `gorethink:"id"`
   firstName  string `gorethink:"name"`
   lastName string `gorethink:"lastname"`
   age int `gorethink:"age"`
}

func initDb() *r.Session {
    var err error
    session, err := r.Connect(r.ConnectOpts{
        Address: "192.168.0.6:28015",
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
        //v1.POST("/users", PostUser)
        v1.GET("/users", GetUsers)
        //v1.DELETE("/users/:id", DeleteUser)
    }

   router.Run(":1337")
}

func createTable(session *r.Session) {
    result, err := r.TableCreate("personas").RunWrite(session)
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
    cursor, err := r.Table("personas").GetAll(2).Run(db)
    defer cursor.Close()

   var hero map[string]interface{}
    err = cursor.One(&hero)
    if err != nil { // should I be checking for EOF here and just trying a restart?
        log.Panic("getOneUser 2:",err)
        fmt.Println("nonil")
    }
    c.JSON(200, hero)
    fmt.Println(">>",hero["firstName"],"<<")

}

func printStr(v string) {
    fmt.Println(v)
}

func printObj(v interface{}) {
    vBytes, _ := json.Marshal(v)
    fmt.Println(string(vBytes))
}