package main

import (
	"html/template"
	"net/http"

	"github.com/DanyZugz/Software-Generator/internal/services"
	"github.com/go-chi/chi/v5"

	"encoding/json"
	"io/ioutil"
	"log"
)

func Home(w http.ResponseWriter, r *http.Request) {

	// Preferiria sacar esto como const pero ando ninja en Go
	ENTRY_POINT_PATH := "./ui/base.html"
	MANIFEST_PATH := "./ui/react/dist/manifest.json"

	////////// DANI REVISA ESTO ///////////

	content, err := ioutil.ReadFile(MANIFEST_PATH)
	if err != nil {
		log.Fatal("Error with react manifest file: ", err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	type Statics struct {
		Css interface{}
		Jsx interface{}
	}

	sfiles := Statics{
		Css: payload["src/main.css"],
		Jsx: payload["src/main.jsx"],
	}

	//////////////////////////////////////////

	template, err := template.ParseFiles(ENTRY_POINT_PATH)

	if err != nil {
		panic(err)
	} else {
		template.Execute(w, sfiles)
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

	services.GenerateFolder(projectname, "bin")
	services.GenerateFolder(projectname, "cmd")
	services.GenerateFolder(projectname, "cmd/api")
	services.GenerateFolder(projectname, "internal")
	services.GenerateFolder(projectname, "internal/data")
	services.GenerateFolder(projectname, "test")
	services.GenerateFolder(projectname, "ui")
	services.GenerateFolder(projectname, "ui/html")
	services.GenerateFolder(projectname, "ui/static")

	services.GenerateFiles(projectname, "/ui/html/index.html", "html.tmpl")
	services.GenerateFiles(projectname, "/ui/static/main.css", "css.tmpl")
	services.GenerateFiles(projectname, "/ui/static/main.js", "js.tmpl")
	services.GenerateFiles(projectname, "/cmd/api/server.go", "chi.tmpl")
	services.GenerateFiles(projectname, "/cmd/api/handlers.go", "handlers.tmpl")

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

func Download(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download"))
}
