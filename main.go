package main

import (
	"log"
	"net/http"
	"os"

	"github.com/c1freitas/reporting/db"
	"github.com/c1freitas/reporting/handlers"
)

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

	dbConnection := db.DBConnection{
		Host:     db.Host,
		Port:     db.Port,
		Dbname:   db.DataBaseName,
		User:     dbUser,
		Password: dbPassword,
	}

	err := dbConnection.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the Database: %v", err)
	}

	defer dbConnection.Close()

	log.Println("Registering Handler(s)")
	http.HandleFunc("/report", handlers.ReportHandler(&dbConnection))

	log.Printf("Starting Server on port %v", port)
	http.ListenAndServe(":"+port, nil)
}
