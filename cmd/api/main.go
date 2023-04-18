package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 3000

type application struct {
	Domain string
}

func main() {
	fmt.Println("Yallo shmorld!")
	var app application
	app.Domain = "sxxxxxxxxxxxx.com"

	log.Printf("Starting the server on port %d ...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes()); err != nil {
		log.Fatal(err)
	}
}
