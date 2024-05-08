package main

import (
	"flag"
	"log"
	"net/http"

	"readingList.test/internal/models"
)

type application struct {
	readinglist *models.ReadinglistModel
}

func main() {

	addr := flag.String("addr", ":8081", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:8080/v1/books", "Endpoint for readinglist the web server")

	app := application{
		readinglist: &models.ReadinglistModel{Endpoint: *endpoint},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
