package main

import (
	"archive/zip"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"path/filepath"

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

	projectname := chi.URLParam(r, "name")

	// Create a temporary file to store the ZIP file
	zipFile, err := os.CreateTemp("", "")
	if err != nil {
		log.Println("el error es cuando crear el file temporal")
		log.Fatal(err)
	}

	// defer os.Remove(zipFile.Name())

	// Create a writer for the ZIP file
	zipWriter := zip.NewWriter(zipFile)

	dir := "/generatedProjects/" + projectname

	// Compress the folder into the ZIP file
	folderPath := filepath.Join(".", dir)
	err = services.CompressFolder(folderPath, zipWriter)
	if err != nil {
		log.Println("el error es cuando comprime la carpeta")
		log.Fatal(err)
	}

	// Close the ZIP file writer
	err = zipWriter.Close()
	if err != nil {
		log.Println("el error es cuando cierra file writer")
		log.Fatal(err)
	}

	// Send ZIP file as HTTP response
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", projectname))
	http.ServeFile(w, r, zipFile.Name())

}

func CreateHttpProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Simple Http"))

}

func CreateDddProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DDD Arch"))
}
