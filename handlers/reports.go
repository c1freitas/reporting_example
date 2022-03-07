package handlers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/c1freitas/reporting/db"
)

/*
 * reportHandler is the http endpoint. returns the CSV file or an error
 */
func ReportHandler(db *db.DBConnection) http.HandlerFunc {
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
