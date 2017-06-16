package controllers
import(
	"fmt"
	"log"
	"encoding/json"
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"../models"
	)
//var user = []models.User{}

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
		user :=[]models.Users{}
		cursor.One(&user)
		c.JSON(200, user)
	}else{
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
	}
}

func PostUser(c *gin.Context) {
	db := initDb()
	defer db.Close()

	user := models.Users{}
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

	user := models.Users{}
	err:=c.Bind(&user)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(user)
	id := c.Params.ByName("id")
	fmt.Println("=======", id)

	if user.FirstName != "" && user.LastName != ""{
	if id != "" {
		result := models.Users{
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