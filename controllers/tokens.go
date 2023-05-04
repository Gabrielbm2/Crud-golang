package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gocrud/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsTokens := models.BuscaTodosOsTokens()
	temp.ExecuteTemplate(w, "Index", todosOsTokens)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		preco := r.FormValue("preco")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		models.CriaNovoToken(nome, precoConvertidoParaFloat)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoToken := r.URL.Query().Get("id")
	models.DeletaToken(idDoToken)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoToken := r.URL.Query().Get("id")
	Token := models.EditaToken(idDoToken)
	temp.ExecuteTemplate(w, "Edit", Token)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		preco := r.FormValue("preco")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		models.AtualizaToken(idConvertidaParaInt, nome, precoConvertidoParaFloat)
	}
	http.Redirect(w, r, "/", 301)
}
