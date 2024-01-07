package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "./ui/react/")) //Configuration to serve statics files from ./ui/static/ directory
	FileServer(r, "/static", filesDir)

	r.Get("/", Home)
	r.Get("/generator/create", ViewCreatedPro)
	r.Get("/generator/download/{name}", Download)

	r.Post("/generator/create/js/vue", CreateVueProject)
	r.Post("/generator/create/js/react", CreateReactProject)
	r.Post("/generator/create/ddd", CreateDddProject)

	r.Post("/generator/create/go/api", CreateApiProject)
	r.Post("/generator/create/go/cobra", CreateCobraProject)

	/* r.Route("/generator/create/go", func(r chi.Router) {
		r.Post("/simpleui/chi", CreateChiProjectUi)
		r.Post("/simpleui/gin", CreateGinProjectUi)
		r.Post("/simpleui/gorilla", CreateMuxProjectUi)
		r.Post("/simpleui/http", CreateHttpProjectUi)

		r.Post("/chi", CreateChiProject)
		r.Post("/gin", CreateGinProject)
		r.Post("/gorilla", CreateMuxProject)
		r.Post("/http", CreateHttpProject)

		r.Post("/cobra", CreateCobraProject)
	}) */

	log.Println("initializing server at http://127.0.0.1:3000")
	http.ListenAndServe(":3000", r)
}

// Handling fileserver in Chi-router

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
