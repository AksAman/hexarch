package pgdb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AksAman/hexarch/utils"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	logger = utils.InitializeLogger("adapters.frameworks.right.db")
}

type Adapter struct {
	db *sql.DB
}

const TABLE_NAME = "arith_history"

func NewAdapter(driverName, dataSource string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// test connection
	logger.Debugf("pinging")
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging to database: %v", err)
	}
	logger.Debugf("db ping successful")

	return &Adapter{
		db: db,
	}, nil
}

func (dbAdapter Adapter) CloseDBConnection() error {
	err := dbAdapter.db.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %v", err)
	}
	return nil
}

func (dbAdapter Adapter) AddToHistory(answer int32, operation string) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	queryStr, args, err := psql.Insert(TABLE_NAME).Columns(
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

	logger.Debugf("queryStr Created: %q with args: %v", queryStr, args)

	_, err = dbAdapter.db.Exec(queryStr, args...)
	if err != nil {
		return err
	}

	return nil
}
