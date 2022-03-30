package main

import (
	"log"
	"net/http"
	"time"

	"todolist/dbconnection"
	"todolist/routes"

	"github.com/rs/cors"
)

func main() {

	r := routes.AllRoutes()
	handler := cors.AllowAll().Handler(r)

	server := &http.Server{
		Addr:         dbconnection.Port,
		Handler:      handler,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Listening on port ", dbconnection.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}
