package db

import "database/sql"

type DIDRecord struct {
	DID      string
	Document string
	Hash     string
	Owner    string
}

type PostgresDB struct {
	db *sql.DB
}
