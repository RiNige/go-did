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

// 修改的地方 新加 --- VC 数据结构 ---
type VCRecord struct {
	ID           string
	Issuer       string
	Subject      string
	Claims       map[string]string
	IssuanceDate string
	Signature    string
	Raw          []byte
}
