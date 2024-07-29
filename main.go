package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func main() {
	connection.InitDb()
	defer func() {
		if err := connection.DB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()
	mux := config.Routes()
	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
