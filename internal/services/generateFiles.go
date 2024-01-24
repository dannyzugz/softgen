package services

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path"

	"github.com/dannyzugz/Software-Generator/internal/models"
)

func GeneratePro(projectname string) {

	dir := "./generatedProjects/" + projectname

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Println(err)
		}

	}
}

func GenerateFolder(directory string, newfolder string) {

	dir := "./generatedProjects/" + directory + "/" + newfolder

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Println(err)
		}

	}
}

func GenerateFiles(data models.ProjectData, dir string, temp string) {

	projectname := data.ProjectName
	var processed bytes.Buffer

	if temp == "html.tmpl" || temp == "css.tmpl" || temp == "js.tmpl" {
		tmpl, err := template.New(temp).ParseFiles("./internal/templates/ui-templates/" + temp)

		if err != nil {
			log.Println(err)
		}

		err = tmpl.Execute(&processed, nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		tmpl, err := template.New(path.Base(temp)).ParseFiles("./internal/templates/" + temp)

		if err != nil {
			log.Println(err)
		}

		err = tmpl.Execute(&processed, data)
		if err != nil {
			log.Println(err)
		}
	}

	directory := "./generatedProjects/" + projectname + dir
	SaveToFile(directory, processed.String())
}

func GenerateBasics(data models.ProjectData) {

	projectname := data.ProjectName

	// All folder for the project
	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "internal/config")
	GenerateFolder(projectname, "database")
	GenerateFolder(projectname, "test")

	// Configuration for the project

	GenerateFiles(data, "/.env", "env.tmpl")
	GenerateFiles(data, "/internal/config/config.go", "config.tmpl")
	GenerateFiles(data, "/go.mod", "mod.tmpl")
}

func GenerateUI(data models.ProjectData) {

	projectname := data.ProjectName

	GenerateFolder(projectname, "ui")
	GenerateFolder(projectname, "ui/html")
	GenerateFolder(projectname, "ui/static")

	GenerateFiles(data, "/ui/html/index.html", "html.tmpl")
	GenerateFiles(data, "/ui/static/main.css", "css.tmpl")
	GenerateFiles(data, "/ui/static/main.js", "js.tmpl")
}
