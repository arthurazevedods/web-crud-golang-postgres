package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/arthurazevedods/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //pega uma informação da url
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		preco := r.FormValue("preco")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")

		idConvertedToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para Int:", err)
		}
		precoConvertedToFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco para float:", err)
		}

		quantidadeConvertedToInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade para Int:", err)
		}

		models.AtualizarProduto(idConvertedToInt, nome, descricao, precoConvertedToFloat, quantidadeConvertedToInt)
	}

	http.Redirect(w, r, "/", 301)
}
