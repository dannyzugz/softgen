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

func ViewCreatedPro(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET Project"))
}

func CreateVueProject(w http.ResponseWriter, r *http.Request) { // Vue js
	w.Write([]byte("Vue Project"))
	projectname := chi.URLParam(r, "name")
	services.GeneratePro(projectname)

}

func CreateReactProject(w http.ResponseWriter, r *http.Request) { // React js
	w.Write([]byte("React Project"))
	projectname := chi.URLParam(r, "name")
	services.GeneratePro(projectname)

}

func CreateChiProject(w http.ResponseWriter, r *http.Request) { // Chi Router
	w.Write([]byte("Chi Project"))
	projectname := chi.URLParam(r, "name")

	services.GeneratePro(projectname)
	services.GeneratePro(projectname)
	services.GenerateFolder(projectname, "bin")
	services.GenerateFolder(projectname, "cmd")
	services.GenerateFolder(projectname, "cmd/api")
	services.GenerateFolder(projectname, "internal")
	services.GenerateFolder(projectname, "internal/data")
	services.GenerateFolder(projectname, "test")
	services.GenerateFolder(projectname, "ui")
	services.GenerateFolder(projectname, "ui/html")
	services.GenerateFolder(projectname, "ui/static")
}

func CreateMuxProject(w http.ResponseWriter, r *http.Request) { // Gorilla Mux
	w.Write([]byte("Gorilla Mux Project"))
	projectname := chi.URLParam(r, "name")

	services.GeneratePro(projectname)
	services.GenerateFolder(projectname, "bin")
	services.GenerateFolder(projectname, "cmd")
	services.GenerateFolder(projectname, "cmd/api")
	services.GenerateFolder(projectname, "internal")
	services.GenerateFolder(projectname, "internal/data")
	services.GenerateFolder(projectname, "test")
	services.GenerateFolder(projectname, "ui")
	services.GenerateFolder(projectname, "ui/html")
	services.GenerateFolder(projectname, "ui/static")
}

func CreateGinProject(w http.ResponseWriter, r *http.Request) { // Gin framework
	w.Write([]byte("Gin Project"))
	projectname := chi.URLParam(r, "name")
	services.GeneratePro(projectname)

}

func CreateCobraProject(w http.ResponseWriter, r *http.Request) { // Cobra framework
	w.Write([]byte("Gin Project"))
	projectname := chi.URLParam(r, "name")
	services.GeneratePro(projectname)

}
