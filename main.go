package main

import (
	"log"
	"net/http"

	"github.com/andrewvo148/pkg/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()

	// Use middleware if needed
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve static files from the embedded file system
	r.Handle("/*", http.FileServer(http.FS(web.SwaggerUI)))

	// Start the server on port 5555
	err := http.ListenAndServe(":5555", r)
	if err != nil {
		log.Fatal(err)
	}
}
