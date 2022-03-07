package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/c1freitas/reporting/db"
)

/*
 * reportHandler is the http endpoint. returns the CSV file or an error
 */
func reportHandler(db *db.DBConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// ctx := req.Context()
		log.Println("server: hello handler started")
		defer log.Println("server: hello handler ended")
		summary, err := db.Query()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v", err)
		}

		buf := &bytes.Buffer{}
		wr := csv.NewWriter(buf)

		for _, fi := range summary {
			s := []string{fi.Id, fi.External_id, fi.External_mission_id, fi.Hardware_id, fi.Nickname, fi.Meta, fi.Created_at.String(), fi.Updated_at.String()}
			wr.Write(s)
		}

		wr.Flush()

		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
		w.Write(buf.Bytes())
	}
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
	http.HandleFunc("/report", reportHandler(&dbConnection))

	log.Printf("Starting Server on port %v", port)
	http.ListenAndServe(":"+port, nil)
}
