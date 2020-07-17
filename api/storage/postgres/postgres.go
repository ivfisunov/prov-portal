package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"time"

	"xxx.ru/ds-dit/api/storage"
)

type postgres struct {
	db *sql.DB
}

// New returns database connection pool
func New() (storage.Services, error) {
	var (
		user   = os.Getenv("PG_USERNAME")
		pass   = os.Getenv("PG_PASSWORD")
		host   = os.Getenv("PG_HOST")
		port   = os.Getenv("PG_PORT")
		dbName = os.Getenv("PG_DATABASE")
	)

	if runtime.GOOS != "linux" {
		host = "host.docker.internal"
	}

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, pass, host, port, dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &postgres{db}, nil
}

func (p *postgres) Close() {
	p.db.Close()
}
