package main

import (
	"log"
	"net/http"

	"github.com/andrewvo148/pkg/openapi"
	"github.com/andrewvo148/pkg/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	// Use middleware if needed
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve static files from the embedded file system
	r.Mount("/bookstore-spec/", http.StripPrefix("/bookstore-spec/", http.FileServer(http.FS(openapi.SwaggerUI))))
	r.Handle("/*", http.FileServer(http.FS(web.WebUI)))

	// Start the server on port 5555
	err := http.ListenAndServe(":5555", r)
	if err != nil {
		log.Fatal(err)
	}
}
