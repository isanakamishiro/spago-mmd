package main

import (
	"log"
	"net/http"
	"os"
)

//go:generate rm -Rf dist/
//go:generate sh -c "cd frontend && spago deploy -tinygo ../dist"
//go:generate cp -Rf frontend/assets frontend/favicon.ico frontend/serviceworker.js dist/

func main() {
	http.Handle("/", http.FileServer(http.Dir("./dist")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
