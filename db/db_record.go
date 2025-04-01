package db

import (
	"database/sql"
	"time"
)

type DIDRecord struct {
	DID       string
	Document  string
	Hash      string
	Owner     string
	CreatedAt time.Time
}

type PostgresDB struct {
	db *sql.DB
}
