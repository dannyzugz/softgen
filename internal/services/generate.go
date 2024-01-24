package services

import "github.com/dannyzugz/Software-Generator/internal/models"

// Chi Projects

func GenChi(data models.ProjectData) { // Generate basic Chi Project

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "chi/chi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChi.tmpl")
}

func GenChiUi(data models.ProjectData) { // Generate Chi Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "chi/chiUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChiUi.tmpl")
}

func GenChiPost(data models.ProjectData) { // Generate Chi Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "chi/chi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChi.tmpl")

}

func GenChiPostUi(data models.ProjectData) { // Generate Chi Project with PostgreSQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "chi/chiUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChiUi.tmpl")
}

func GenChiMsql(data models.ProjectData) { // Generate Chi Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "chi/chi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChi.tmpl")
}

func GenChiMsqlUi(data models.ProjectData) { // Generate Chi Project with MySQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "chi/chiUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChiUi.tmpl")
}

// Gin Projects

func GenGin(data models.ProjectData) { // Generate basic Gin Project

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data)

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gin/gin.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGin.tmpl")
}

func GenGinUi(data models.ProjectData) { // Generate Gin Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data)
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gin/ginUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGinUi.tmpl")
}

func GenGinPost(data models.ProjectData) { // Generate Gin Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gin/gin.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGin.tmpl")
}

func GenGinPostUi(data models.ProjectData) { // Generate Gin Project with PostgreSQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data)
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gin/ginUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGinUi.tmpl")
}

func GenGinMsql(data models.ProjectData) { // Generate Gin Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gin/gin.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGin.tmpl")
}

func GenGinMsqlUi(data models.ProjectData) { // Generate Gin Project with MySQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data)
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "gin/ginUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGinUi.tmpl")
}

// Gorilla Mux Projects

func GenMux(data models.ProjectData) { // Generate basic Gorilla Mux Project

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "mux/mux.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMux.tmpl")
}

func GenMuxUi(data models.ProjectData) { // Generate Gorilla Mux Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "mux/muxUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxUi.tmpl")
}

func GenMuxPost(data models.ProjectData) { // Generate Gorilla Mux Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "mux/mux.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMux.tmpl")
}

func GenMuxPostUi(data models.ProjectData) { // Generate Gorilla Mux Project with PostgreSQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "mux/muxUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxUi.tmpl")
}

func GenMuxMsql(data models.ProjectData) { // Generate Gorilla Mux Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "mux/mux.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMux.tmpl")
}

func GenMuxMsqlUi(data models.ProjectData) { // Generate Gorilla Mux Project with MySQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "mux/muxUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxUi.tmpl")
}

// Fiber Projects

func GenFib(data models.ProjectData) { //Generate basic Fiber Project

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "fib/fib.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

func GenFibUi(data models.ProjectData) { // Generate Fiber Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "fib/fibUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFibUi.tmpl")
}

func GenFibPost(data models.ProjectData) { // Generate Fiber Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "fib/fib.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

func GenFibPostUi(data models.ProjectData) { // Generate Fiber Project with PostgreSQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "fib/fib.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

func GenFibMsql(data models.ProjectData) { // Generate Fiber Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "fib/fib.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

func GenFibMsqlUi(data models.ProjectData) { // Generate Fiber Project with MySQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "fib/fib.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

// Http Projects

func GenHttp(data models.ProjectData) { // Generate basic Http Project

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "http/server.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}

func GenHttpUi(data models.ProjectData) { // Generate Http Project with basic Ui

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "main.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "http/server.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handlerUi.tmpl")
}

func GenHttpPost(data models.ProjectData) { // Generate Http Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "http/server.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}

func GenHttpPostUi(data models.ProjectData) { // Generate Http Project with PostgreSQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initPsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "http/server.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}

func GenHttpMsql(data models.ProjectData) { // Generate Http Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "http/server.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}

func GenHttpMsqlUi(data models.ProjectData) { // Generate Http Project with MySQL DB and UI

	projectname := data.ProjectName

	GeneratePro(projectname)
	GenerateBasics(data) // This function generate all folders and config for the project
	GenerateUI(data)

	GenerateFiles(data, "/main.go", "mainDb.tmpl")
	GenerateFiles(data, "/database/init.go", "db-templates/initMsql.tmpl")
	GenerateFiles(data, "/cmd/api/server.go", "http/server.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}
