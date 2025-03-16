package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var postgresDB *sql.DB

func ConnectPostgres(logger *zap.Logger) error {
	// TODO add connection config
	connStr := "postgresql://test:test@localhost/recipio?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("DB connection failed", zap.Error(err))
		return ErrDBNotConnected
	}
	if err := db.Ping(); err != nil {
		logger.Error("DB ping failed", zap.Error(err))
		return ErrDBNotConnected
	}
	postgresDB = db
	return nil
}

func GetPostgres() *sql.DB {
	return postgresDB
}
