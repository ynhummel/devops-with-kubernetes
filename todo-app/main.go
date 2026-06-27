package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8001"
	}

	log.Printf("Server started in port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
