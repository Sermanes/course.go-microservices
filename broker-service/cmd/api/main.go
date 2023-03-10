package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort string = "80"

type Config struct {}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server {
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Panic(srv.ListenAndServe())	
}