package main

import (
	"html/template"
	"net/http"

	"github.com/DanyZugz/Software-Generator/internal/services"
	"github.com/go-chi/chi/v5"
)

func Home(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}

func Func1(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Func1"))
}

func Func2(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Func2"))
}

func Func3(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Func3"))
}

func Func4(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Func4"))
}