package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MYSQL_USER     = "MYSQL_USER"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_PORT     = "MYSQL_PORT"
	MYSQL_DATABASE = "MYSQL_DATABASE"
)

func NewMySQLConnection() (*sql.DB, error) {
	user := os.Getenv(MYSQL_USER)
	password := os.Getenv(MYSQL_PASSWORD)
	host := os.Getenv(MYSQL_HOST)
	port := os.Getenv(MYSQL_PORT)
	database := os.Getenv(MYSQL_DATABASE)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Testa a conexão
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
