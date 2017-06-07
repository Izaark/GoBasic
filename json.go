package main 
import (
	"net/http"
	"fmt"
	"encoding/json"
)
type Curso struct {
	User string  `json:"user"`
	Title string `json:"title"`
	NumeroVideos int `json:"videos"`
}
func main() {
	http.HandleFunc("/", format_json )
	http.ListenAndServe(":8000",nil)
}

type Cursos []Curso
func format_json(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Peticion !!!")
	cursos := Cursos{
		Curso{"Isaac", "Curso de Go", 20},
		Curso{"Francisco", "Python", 60},
		Curso{"Juan", "RoR", 100},
	}
	json.NewEncoder(w).Encode(cursos)
	}