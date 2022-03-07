package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func reportHandler(w http.ResponseWriter, req *http.Request) {
	// ctx := req.Context()
	log.Println("server: hello handler started")
	defer log.Println("server: hello handler ended")

	fmt.Fprintf(w, "hello v2\n")
}

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatalln("PORT value not set correctly")
	}
	dbUser := os.Getenv("POSTGRES_USER")
	if len(dbUser) == 0 {
		log.Fatalln("POSTGRES_USER value not set correctly")
	}
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if len(dbPassword) == 0 {
		log.Fatalln("POSTGRES_PASSWORD value not set correctly")
	}

	log.Println("Registering Handler(s)")
	http.HandleFunc("/report", reportHandler)

	log.Printf("Starting Server on port %v", port)
	http.ListenAndServe(":"+port, nil)
}
