package services

import (
	"log"
	"os"
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
