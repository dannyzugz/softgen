package services

import (
	"bytes"
	"html/template"
	"log"
)

func GenerateFiles(projectname string, dir string, temp string) {

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
		tmpl, err := template.New(temp).ParseFiles("./internal/templates/" + temp)

		if err != nil {
			log.Println(err)
		}

		err = tmpl.Execute(&processed, nil)
		if err != nil {
			log.Println(err)
		}
	}

	directory := "./generatedProjects/" + projectname + dir
	SaveToFile(directory, processed.String())
}
