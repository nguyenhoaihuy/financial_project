package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	DB *sql.DB
}

func NewDBManager(dsn string) (*DBManager, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &DBManager{DB: db}, nil
}

func (m *DBManager) Close() {
	m.DB.Close()
}
