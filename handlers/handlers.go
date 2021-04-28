package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func Index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
