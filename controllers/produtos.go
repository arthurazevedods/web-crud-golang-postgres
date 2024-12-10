package controllers

import (
	"html/template"
	"net/http"

	"github.com/arthurazevedods/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)
}
