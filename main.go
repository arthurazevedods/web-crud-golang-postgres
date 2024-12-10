package main

import (
	"html/template"
	"net/http"

	"github.com/arthurazevedods/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)
}
