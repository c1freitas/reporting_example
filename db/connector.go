package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	Host         = "db"
	Port         = 5432
	DataBaseName = "reporting"
)

type DBConnection struct {
	Host       string
	Port       int
	User       string
	Password   string
	Dbname     string
	Connection *sql.DB
}

type FlightSummary struct {
	Id                  string
	External_id         string
	External_mission_id string
	Hardware_id         string
	Nickname            string
	Meta                string
	Created_at          time.Time
	Updated_at          time.Time
}

/*
 * Connect creates a connection to the DB
 */
func (d *DBConnection) Connect() error {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.Dbname)

	// open database
	var err error
	d.Connection, err = sql.Open("pgx", psqlconn)
	if err != nil {
		log.Printf("Could not connect to database: %v", err)
		return err
	}

	// check db
	err = d.Connection.Ping()
	if err != nil {
		log.Printf("Could not ping database: %v", err)
		return err
	}

	log.Println("Connected!")
	return nil
}

/*
 * Query returns all the flight summary data
 */
func (d *DBConnection) Query() ([]FlightSummary, error) {
	rows, err := d.Connection.Query("SELECT id, external_id, external_mission_id, hardware_id, nickname, meta, created_at, updated_at FROM flight_summary")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var flightSummaries []FlightSummary

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var fs FlightSummary
		if err := rows.Scan(&fs.Id, &fs.External_id, &fs.External_mission_id, &fs.Hardware_id, &fs.Nickname, &fs.Meta, &fs.Created_at, &fs.Updated_at); err != nil {
			return flightSummaries, err
		}
		flightSummaries = append(flightSummaries, fs)
	}
	if err = rows.Err(); err != nil {
		return flightSummaries, err
	}
	return flightSummaries, nil
}

/*
 * Close closes the connection to the Database. Should only be used when shutting down the application
 */
func (d *DBConnection) Close() {
	d.Connection.Close()
}
