package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sri2103/htmx_go/internal/config"
	"github.com/sri2103/htmx_go/internal/pkg/database/scripts"
)

type DB struct {
	Conn *sql.DB
}

func NewDatabase(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	_, err = db.Exec(scripts.UserTable)

	if err != nil {
		return nil, fmt.Errorf("failed to create user table: %w", err)
	}

	_, err = db.Exec(scripts.TodosTable)
	if err != nil {
		return nil, fmt.Errorf("failed to create todos table: %w", err)
	}
	return db, nil
}

func ConnectSQL(config *config.AppConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.DBName,
		config.DB.User,
		config.DB.Password,
		config.DB.SslMode,
	)
	db, err := NewDatabase(dsn)
	if err != nil {
		fmt.Println(err.Error(),"Error for creating database")
		panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{Conn: db}, nil

}
