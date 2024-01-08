package services

import "github.com/DanyZugz/Software-Generator/internal/models"

func GenChiUi(data models.ProjectData) { // Generate Chi Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")
	GenerateFolder(projectname, "ui")
	GenerateFolder(projectname, "ui/html")
	GenerateFolder(projectname, "ui/static")

	GenerateFiles(data, "/ui/html/index.html", "html.tmpl")
	GenerateFiles(data, "/ui/static/main.css", "css.tmpl")
	GenerateFiles(data, "/ui/static/main.js", "js.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "chi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "handlerMux.tmpl")
}

func GenChi(data models.ProjectData) { // Generate basic Chi Project

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "chiBasic.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "handlerChiBasic.tmpl")
}

func GenGinUi(data models.ProjectData) { //Generate Gin Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")
	GenerateFolder(projectname, "ui")
	GenerateFolder(projectname, "ui/html")
	GenerateFolder(projectname, "ui/static")

	GenerateFiles(data, "/ui/html/index.html", "html.tmpl")
	GenerateFiles(data, "/ui/static/main.css", "css.tmpl")
	GenerateFiles(data, "/ui/static/main.js", "js.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gin.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "handlerGin.tmpl")
}

func GenGin(data models.ProjectData) { //Generate basic Gin Project

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "ginBasic.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "handlerGin.tmpl")
}

func GenMuxUi(data models.ProjectData) { //Generate GorillaMux Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")
	GenerateFolder(projectname, "ui")
	GenerateFolder(projectname, "ui/html")
	GenerateFolder(projectname, "ui/static")

	GenerateFiles(data, "/ui/html/index.html", "html.tmpl")
	GenerateFiles(data, "/ui/static/main.css", "css.tmpl")
	GenerateFiles(data, "/ui/static/main.js", "js.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gorilla.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "handlerGin.tmpl")
}
