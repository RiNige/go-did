package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Set up PostgreSQL Connection Parameter
// 修改的地方
// const (
// 	host     = "ftec5520.cuz8oe2wwvba.us-east-1.rds.amazonaws.com"
// 	port     = 5432
// 	user     = "ftec_5520"
// 	password = "ftec_5520_group11"
// 	dbname   = "postgres"
// )

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"          // 默认用户
	password = "ftec_5520_group11" // ❗输入你自己安装数据库时设置的密码
	dbname   = "postgres"          // 默认数据库名
)

func ConnectDB() *PostgresDB {
	// Set up Connection String Parameters, available at https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters
	// 修改的地方
	// psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s host=%s", user, dbname, password, host)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
	row := p.db.QueryRow("SELECT did, document, hash, owner FROM did_documents WHERE did = $1", did)
	var record DIDRecord
	err := row.Scan(&record.DID, &record.Document, &record.Hash, &record.Owner)
	return &record, err
}

func (p *PostgresDB) Close() error {
	err := p.db.Close()
	return err
}

// 修改的地方 新加VC
// ✅ 保存 VC（Verifiable Credential）数据
func (p *PostgresDB) SaveVC(vc VCRecord) error {
	query := `
		INSERT INTO verifiable_credentials (id, issuer, subject, claims, issuance_date, signature, raw)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO NOTHING
	`

	claimsJSON, err := json.Marshal(vc.Claims)
	if err != nil {
		return fmt.Errorf("failed to marshal claims: %v", err)
	}

	_, err = p.db.Exec(query,
		vc.ID,
		vc.Issuer,
		vc.Subject,
		claimsJSON,
		vc.IssuanceDate,
		vc.Signature,
		vc.Raw,
	)
	if err != nil {
		return fmt.Errorf("failed to insert VC: %v", err)
	}

	return nil
}

// ✅ 新增：GetVC - 根据 VC ID 查询数据库中的 VC
func (p *PostgresDB) GetVC(id string) (*VCRecord, error) {
	row := p.db.QueryRow(`
		SELECT id, issuer, subject, claims, issuance_date, signature, raw
		FROM verifiable_credentials
		WHERE id = $1
	`, id)

	var vc VCRecord
	var claimsJSON []byte
	var rawJSON []byte

	err := row.Scan(&vc.ID, &vc.Issuer, &vc.Subject, &claimsJSON, &vc.IssuanceDate, &vc.Signature, &rawJSON)
	if err != nil {
		return nil, err
	}

	// 解析 claims 和 raw 字段
	if err := json.Unmarshal(claimsJSON, &vc.Claims); err != nil {
		return nil, err
	}
	vc.Raw = rawJSON

	return &vc, nil
}
