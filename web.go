package main 
import (
	"net/http"
	// /  "io"
	//"fmt"
)
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000",nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Peticion !!!")
	//io.WriteString(w,"Hola mundo")
	http.ServeFile(w,r,"index.html")
}