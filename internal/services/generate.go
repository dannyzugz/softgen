package services

import "github.com/DanyZugz/Software-Generator/internal/models"

// Chi Projects

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

	GenerateFiles(data, "/cmd/api/server.go", "chi/chi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChi.tmpl")
}

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
	GenerateFiles(data, "/cmd/api/server.go", "chi/chiUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChiUi.tmpl")
}

func GenChiPost(data models.ProjectData) { // Generate Chi Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "chi/chiDbPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChi.tmpl")
}

func GenChiPostUi(data models.ProjectData) { // Generate Chi Project with PostgreSQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "chi/chiUiPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChiUi.tmpl")
}

func GenChiMsql(data models.ProjectData) { // Generate Chi Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "chi/chiDbMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChi.tmpl")
}

func GenChiMsqlUi(data models.ProjectData) { // Generate Chi Project with MySQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "chi/chiUiMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "chi/handlerChiUi.tmpl")
}

// Gin Projects

func GenGin(data models.ProjectData) { // Generate basic Gin Project

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "gin/gin.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGin.tmpl")
}

func GenGinUi(data models.ProjectData) { // Generate Gin Project with basic Ui

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
	GenerateFiles(data, "/cmd/api/server.go", "gin/ginUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGinUi.tmpl")
}

func GenGinPost(data models.ProjectData) { // Generate Gin Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "gin/ginDbPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGin.tmpl")
}

func GenGinPostUi(data models.ProjectData) { // Generate Gin Project with PostgreSQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "gin/ginUiPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGinUi.tmpl")
}

func GenGinMsql(data models.ProjectData) { // Generate Gin Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "gin/ginDbMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGin.tmpl")
}

func GenGinMsqlUi(data models.ProjectData) { // Generate Gin Project with MySQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "gin/ginUiMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "gin/handlerGinUi.tmpl")
}

// Gorilla Mux Projects

func GenMux(data models.ProjectData) { // Generate basic Gorilla Mux Project

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "mux/mux.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMux.tmpl")
}

func GenMuxUi(data models.ProjectData) { // Generate Gorilla Mux Project with basic Ui

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
	GenerateFiles(data, "/cmd/api/server.go", "mux/muxUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxUi.tmpl")
}

func GenMuxPost(data models.ProjectData) { // Generate Gorilla Mux Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "mux/muxDbPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxDb.tmpl")
}

func GenMuxPostUi(data models.ProjectData) { // Generate Gorilla Mux Project with PostgreSQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "mux/muxUiPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxUi.tmpl")
}

func GenMuxMsql(data models.ProjectData) { // Generate Gorilla Mux Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "mux/muxDbMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxDb.tmpl")
}

func GenMuxMsqlUi(data models.ProjectData) { // Generate Gorilla Mux Project with MySQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "mux/muxUiMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "mux/handlerMuxUi.tmpl")
}

// Fiber Projects

func GenFib(data models.ProjectData) { //Generate basic Fiber Project

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "fib/fib.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

func GenFibUi(data models.ProjectData) { // Generate Fiber Project with basic Ui

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
	GenerateFiles(data, "/cmd/api/server.go", "fib/fibUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFibUi.tmpl")
}

func GenFibPost(data models.ProjectData) { // Generate Fiber Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "fib/fibDbPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

func GenFibPostUi(data models.ProjectData) { // Generate Fiber Project with PostgreSQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "fib/fibUiPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFibUi.tmpl")
}

func GenFibMsql(data models.ProjectData) { // Generate Fiber Project with MySQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "fib/fibDbMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFib.tmpl")
}

func GenFibMsqlUi(data models.ProjectData) { // Generate Fiber Project with MySQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "fib/fibUiPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "fib/handlerFibUi.tmpl")
}

// Http Projects

func GenHttp(data models.ProjectData) { // Generate basic Http Project

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "http/server.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}

func GenHttpUi(data models.ProjectData) { // Generate Http Project with basic Ui

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
	GenerateFiles(data, "/cmd/api/server.go", "http/serverUi.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handlerUi.tmpl")
}

func GenHttpPost(data models.ProjectData) { // Generate Http Project with PostgreSQL DB

	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "http/httpDbPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}

func GenHttpPostUi(data models.ProjectData) { // Generate Http Project with PostgreSQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "http/UiPost.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handlerUi.tmpl")
}

func GenHttpMsql(data models.ProjectData) { // Generate Http Project with MySQL DB
	projectname := data.ProjectName

	GeneratePro(projectname)

	GenerateFolder(projectname, "bin")
	GenerateFolder(projectname, "cmd")
	GenerateFolder(projectname, "cmd/api")
	GenerateFolder(projectname, "internal")
	GenerateFolder(projectname, "internal/models")
	GenerateFolder(projectname, "internal/services")
	GenerateFolder(projectname, "test")

	GenerateFiles(data, "/cmd/api/server.go", "http/httpDbMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handler.tmpl")
}

func GenHttpMsqlUi(data models.ProjectData) { // Generate Http Project with MySQL DB and UI

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
	GenerateFiles(data, "/cmd/api/server.go", "http/UiMysq.tmpl")
	GenerateFiles(data, "/cmd/api/handlers.go", "http/handlerUi.tmpl")
}
