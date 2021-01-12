package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	student "./exercices"
)

var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))
var templatesDir = os.Getenv("TEMPLATES_DIR")

func main() {
	fmt.Println("==>  * localhost:9000 *  <==")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/Ascii", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		ascii := r.Form["machaine"][0]

		//fmt.Println("test1")
		//fmt.Fprint(w, "student.Ascii(ascii)")
		fmt.Fprint(w, student.Ascii(ascii))

		//	fmt.Println("test2")
	})
	log.Fatal(http.ListenAndServe(":9000", nil))

}
