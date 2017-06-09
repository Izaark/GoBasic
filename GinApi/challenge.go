package main
import (
     "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func InitDb() *gorm.DB {
   
    db, err := gorm.Open("sqlite3", "./data.db")
    db.LogMode(true)
    if err != nil {
        panic(err)
    }

    if !db.HasTable(&Users{}) {
        db.CreateTable(&Users{})
        db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
    }

    return db
}

type Users struct {
    Id int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Firstname string `gorm:"not null" form:"firstname" json:"firstName"`
    Lastname  string `gorm:"not null" form:"lastname" json:"lastName"`
    Age int `gorm:"not null" form:"age" json:"age"`
}

func main() {
    router := gin.Default()

    v1 := router.Group("api")
    {
        v1.POST("/users", PostUser)
        v1.GET("/users", GetUsers)
        v1.DELETE("/users/:id", DeleteUser)
    }

    router.Run(":8080")
}

func PostUser(c *gin.Context) {
    db := InitDb()
    defer db.Close()

    var user Users
    c.Bind(&user)

    if user.Firstname != "" && user.Lastname != "" {
        db.Create(&user)
        c.JSON(201, gin.H{"success": user})
    } else {
        c.JSON(422, gin.H{"error": "Fields are empty"})
    }

}

func GetUsers(c *gin.Context) {
    db := InitDb()
    defer db.Close()

    var users []Users
    db.Find(&users)

    c.JSON(200, users)

}

func DeleteUser(c *gin.Context) {
	  db := InitDb()
	  defer db.Close()

	  id := c.Params.ByName("id")
	  var user Users
	  db.First(&user, id)

	  if user.Id != 0 {

	  	db.Delete(&user)
	  	c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	  } else {
	  	c.JSON(404, gin.H{"error": "User not found"})
}
  
}




