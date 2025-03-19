package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Set up PostgreSQL Connection Parameter
const (
	host     = "ftec5520.cuz8oe2wwvba.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "ftec_5520"
	password = "ftec_5520_group11"
	dbname   = "postgres"
)

func ConnectDB() *sql.DB {
	// Set up Connection String Parameters, available at https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s host=%s", user, dbname, password, host)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Try out the DB connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to DB at %v:%v", host, port)

	// Configure connection pool for DB
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db
}
