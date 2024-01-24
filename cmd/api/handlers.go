package main

import (
	"archive/zip"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"path/filepath"

	"github.com/dannyzugz/Software-Generator/internal/models"
	"github.com/dannyzugz/Software-Generator/internal/services"
	"github.com/go-chi/chi/v5"

	"encoding/json"
)

func Home(w http.ResponseWriter, r *http.Request) {

	const ENTRY_POINT_PATH = "./ui/base.html"
	const MANIFEST_PATH = "./ui/react/dist/manifest.json"

	////////// DANI REVISA ESTO ///////////

	content, err := os.ReadFile(MANIFEST_PATH)
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

func CreateVueProject(w http.ResponseWriter, r *http.Request) { // Vue js
	w.Write([]byte("Vue Project"))
}

func CreateReactProject(w http.ResponseWriter, r *http.Request) { // React js
	w.Write([]byte("React Project"))
}

func CreateChiProjectUi(w http.ResponseWriter, r *http.Request) { // Chi Router
	w.Write([]byte("Chi Project"))

	var data models.ProjectData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func CreateApiProject(w http.ResponseWriter, r *http.Request) {

	var data models.ProjectData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	router := data.RouterName
	db := data.DbName
	ui := data.Ui

	if ui {
		switch db {
		case "postgresql":
			w.Write([]byte("Project with UI and Postrgres"))
			switch router {
			case "chi":
				w.Write([]byte("Chi Project with UI"))
				services.GenChiPostUi(data)
			case "gin":
				w.Write([]byte("Gin Project"))
				services.GenGinPostUi(data)
			case "fiber":
				w.Write([]byte("fiber Project"))
				services.GenFibPostUi(data)
			case "gorilla":
				w.Write([]byte("gorilla Project"))
				services.GenMuxPostUi(data)
			case "http":
				w.Write([]byte("http Project"))
				services.GenHttpPostUi(data)
			default:
				w.Write([]byte("Error, is not exist that project"))
			}
		case "mysql":
			w.Write([]byte("Project with UI and MySql"))
			switch router {
			case "chi":
				w.Write([]byte("Chi Project with UI"))
				services.GenChiMsqlUi(data)
			case "gin":
				w.Write([]byte("Gin Project"))
				services.GenGinMsqlUi(data)
			case "fiber":
				w.Write([]byte("fiber Project"))
				services.GenFibMsqlUi(data)
			case "gorilla":
				w.Write([]byte("gorilla Project"))
				services.GenMuxMsqlUi(data)
			case "http":
				w.Write([]byte("http Project"))
				services.GenHttpMsqlUi(data)
			default:
				w.Write([]byte("Error, is not exist that project"))
			}
		default:
			switch router {
			case "chi":
				w.Write([]byte("Chi Project with UI"))
				services.GenChiUi(data)
			case "gin":
				w.Write([]byte("Gin Project"))
				services.GenGinUi(data)
			case "fiber":
				w.Write([]byte("fiber Project"))
				services.GenFibUi(data)
			case "gorilla":
				w.Write([]byte("gorilla Project"))
				services.GenMuxUi(data)
			case "http":
				w.Write([]byte("http Project"))
				services.GenHttpUi(data)
			default:
				w.Write([]byte("Error, is not exist that project"))
			}
		}
	} else {
		switch db {
		case "postgresql":
			w.Write([]byte("Project with  Postrgres"))
			switch router {
			case "chi":
				w.Write([]byte("Chi Project with UI"))
				services.GenChiPost(data)
			case "gin":
				w.Write([]byte("Gin Project"))
				services.GenGinPost(data)
			case "fiber":
				w.Write([]byte("fiber Project"))
				services.GenFibPost(data)
			case "gorilla":
				w.Write([]byte("gorilla Project"))
				services.GenMuxPost(data)
			case "http":
				w.Write([]byte("http Project"))
				services.GenHttpPost(data)
			default:
				w.Write([]byte("Error, is not exist that project"))
			}
		case "mysql":
			w.Write([]byte("Project with UI and MySql"))
			switch router {
			case "chi":
				w.Write([]byte("Chi Project with UI"))
				services.GenChiMsql(data)
			case "gin":
				w.Write([]byte("Gin Project"))
				services.GenGinMsql(data)
			case "fiber":
				w.Write([]byte("fiber Project"))
				services.GenFibMsql(data)
			case "gorilla":
				w.Write([]byte("gorilla Project"))
				services.GenMuxMsql(data)
			case "http":
				w.Write([]byte("http Project"))
				services.GenHttpMsql(data)
			default:
				w.Write([]byte("Error, is not exist that project"))
			}
		default:
			switch router {
			case "chi":
				w.Write([]byte("Chi Project with UI"))
				services.GenChi(data)
			case "gin":
				w.Write([]byte("Gin Project"))
				services.GenGin(data)
			case "fiber":
				w.Write([]byte("fiber Project"))
				services.GenFib(data)
			case "gorilla":
				w.Write([]byte("gorilla Project"))
				services.GenMux(data)
			case "http":
				w.Write([]byte("http Project"))
				services.GenHttp(data)
			default:
				w.Write([]byte("Error, is not exist that project"))
			}
		}
	}

}

func CreateDddProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DDD Arch"))
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
		log.Fatal(err)
	}

	// Create a writer for the ZIP file
	zipWriter := zip.NewWriter(zipFile)

	dir := "/generatedProjects/" + projectname

	// Compress the folder into the ZIP file
	folderPath := filepath.Join(".", dir)
	err = services.CompressFolder(folderPath, zipWriter)
	if err != nil {
		log.Fatal(err)
	}

	// Close the ZIP file writer
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send ZIP file as HTTP response
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", projectname))
	http.ServeFile(w, r, zipFile.Name())

}
