package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseRecord struct {
}

type IExporter interface {
	Export(record DatabaseRecord) error
}

type FakeExporter struct {
}

func (e *FakeExporter) Export(record DatabaseRecord) error {
	log.Println("Exporting to fake", record)
	return nil
}

type PostgresExporter struct {
	db *sql.DB
}

func NewPostgresExporter(connectionString string) (*PostgresExporter, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &PostgresExporter{db: db}, nil
}

func (e *PostgresExporter) Export(record DatabaseRecord) error {
	e.db.Exec("INSERT INTO records (data) VALUES ($1)", record)
	return nil
}
