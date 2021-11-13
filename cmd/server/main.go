package main

import (
	"log"
	"net/http"
	"os"

	delivery "github.com/slatermorgan/go-care/playlists/deliveries/http"
)

func main() {
	port := os.Getenv("PORT")

	router, err := delivery.Routes()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Running on port: ", port)
	log.Panic(http.ListenAndServe(":"+port, router))
}
