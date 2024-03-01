package main

import (
	"net/http"

	"andrewvo148/pkg/examples/jobs/api/internal/web"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Mount the FileServer handler to serve static files from the "static" directory
	r.Mount("/", http.FileServer(http.FS(web.WebUI)))

	http.ListenAndServe(":8080", r)
}
