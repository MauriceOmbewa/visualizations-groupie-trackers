package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Athooh/groupie-tracker/internal/api"
)

func main() {
	err := api.InitData()
	if err != nil {
		log.Fatalf("Error initializing data: %v", err)
	}

	// Handle the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			api.IndexHandler(w, r)
		} else {
			api.RenderError(w, http.StatusNotFound, "Page not found")
		}
	})

	http.HandleFunc("/artists", api.ArtistsHandler)
	http.HandleFunc("/artist/", api.ArtistDetailHandler)
	http.HandleFunc("/search", api.SearchHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on port localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
