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

<<<<<<< HEAD
type Users struct {
    Id    int `json:"id" binding:"required" form:"id" gorethink:"id"`
   FirstName  string `json:"firstname" binding:"required" form:"firstName" gorethink:"firstName"`
   LastName string `json:"lastname" binding:"required" form:"lastName" gorethink:"lastName"`
   Age int `json:"age" binding:"required" form:"age" gorethink:"age"`
=======
type Person struct {
    id    int `gorethink:"id"`
   firstName  string `gorethink:"name"`
   lastName string `gorethink:"lastname"`
   age int `gorethink:"age"`
>>>>>>> 3052489c016f09dab4849a831bb843475d786b5a
}

func initDb() *r.Session {
    var err error
    session, err := r.Connect(r.ConnectOpts{
<<<<<<< HEAD
        Address: "192.168.0.4:28015",
=======
        Address: "192.168.0.6:28015",
>>>>>>> 3052489c016f09dab4849a831bb843475d786b5a
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
<<<<<<< HEAD
        v1.POST("/users", PostUser)
        v1.GET("/users", GetUsers)
        v1.DELETE("/users/:id", DeleteUser)
=======
        //v1.POST("/users", PostUser)
        v1.GET("/users", GetUsers)
        //v1.DELETE("/users/:id", DeleteUser)
>>>>>>> 3052489c016f09dab4849a831bb843475d786b5a
    }

   router.Run(":1337")
}

func createTable(session *r.Session) {
<<<<<<< HEAD
    result, err := r.TableCreate("users").RunWrite(session)
=======
    result, err := r.TableCreate("personas").RunWrite(session)
>>>>>>> 3052489c016f09dab4849a831bb843475d786b5a
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
<<<<<<< HEAD
    cursor, err := r.Table("users").Run(db)
    defer cursor.Close()

   var hero []interface{}
    err = cursor.All(&hero)
=======
    cursor, err := r.Table("personas").GetAll(2).Run(db)
    defer cursor.Close()

   var hero map[string]interface{}
    err = cursor.One(&hero)
>>>>>>> 3052489c016f09dab4849a831bb843475d786b5a
    if err != nil { // should I be checking for EOF here and just trying a restart?
        log.Panic("getOneUser 2:",err)
        fmt.Println("nonil")
    }
    c.JSON(200, hero)
<<<<<<< HEAD
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


=======
    fmt.Println(">>",hero["firstName"],"<<")

}

>>>>>>> 3052489c016f09dab4849a831bb843475d786b5a
func printStr(v string) {
    fmt.Println(v)
}

func printObj(v interface{}) {
    vBytes, _ := json.Marshal(v)
    fmt.Println(string(vBytes))
<<<<<<< HEAD
}

=======
}
>>>>>>> 3052489c016f09dab4849a831bb843475d786b5a
