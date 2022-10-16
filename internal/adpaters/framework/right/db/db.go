package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Adapter struct {
	db *sql.DB
}

const TABLE_NAME = "arith_history"

func NewAdapter(driverName, dataSource string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	// test connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging to database: %v", err)
	}

	return &Adapter{
		db: db,
	}, nil
}

func (dbAdapter Adapter) CloseDBConnection() {
	err := dbAdapter.db.Close()
	if err != nil {
		log.Fatalf("error closing database connection: %v", err)
	}
}

func (dbAdapter Adapter) AddToHistory(answer int32, operation string) error {
	queryStr, args, err := sq.Insert(TABLE_NAME).Columns(
		"date",
		"answer",
		"operation",
	).Values(
		time.Now(),
		answer,
		operation,
	).ToSql()

	if err != nil {
		return err
	}

	fmt.Printf("queryStr Created: %q with args: %v\n", queryStr, args)

	_, err = dbAdapter.db.Exec(queryStr, args...)
	if err != nil {
		return err
	}

	return nil
}
