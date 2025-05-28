package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
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

	if user == "" || password == "" || host == "" || port == "" || database == "" {
		log.Println("Erro: uma ou mais variáveis de ambiente do MySQL estão vazias.")
		return nil, fmt.Errorf("variáveis de ambiente de conexão com MySQL não configuradas corretamente")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
	log.Printf("Conectando ao MySQL com DSN: %s\n", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com MySQL: %w", err)
	}

	// Testa a conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao testar conexão com MySQL: %w", err)
	}

	log.Println("Conexão com MySQL estabelecida com sucesso.")
	return db, nil
}
