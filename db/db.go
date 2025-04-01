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

func ConnectDB() *PostgresDB {
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
	fmt.Printf("Connected to DB at %v:%v\n", host, port)

	// Configure connection pool for DB
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &PostgresDB{db: db}
}

func (p *PostgresDB) SaveDID(record DIDRecord) error {
	_, err := p.db.Exec(`
		INSERT INTO did_documents (did, document, hash, owner)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (did) DO UPDATE
		SET document = $2, hash = $3, updated_at = CURRENT_TIMESTAMP`,
		record.DID, record.Document, record.Hash, record.Owner,
	)
	return err
}

func (p *PostgresDB) GetDID(did string) (*DIDRecord, error) {
	row := p.db.QueryRow("SELECT did, document, hash, owner, created_at FROM did_documents WHERE did = $1", did)
	var record DIDRecord
	err := row.Scan(&record.DID, &record.Document, &record.Hash, &record.Owner, &record.CreatedAt)
	return &record, err
}

func (p *PostgresDB) Close() error {
	err := p.db.Close()
	return err
}
