package main

import (
	"fmt"
	"log"
	"net/http"
)

// webPort is the port on which the HTTP server will listen.
const webPort = "80"

type Config struct{}

func main() {
	app := Config{}
	// Log a message indicating that the broker service is starting.
	log.Printf("Starting broker service on port %s\n", webPort)

	// Create an HTTP server.
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
		// GET "/"
	}

	// Start the HTTP server.
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}
