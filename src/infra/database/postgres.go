package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	Connection *sql.DB
}

type PostgresConfiguration struct {
	host     string
	database string
	port     string
	user     string
	password string
}

func environmentConfiguration() *PostgresConfiguration {
	return &PostgresConfiguration{
		host:     os.Getenv("DB_HOST"),
		database: os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
	}
}

func NewPostgresClient() *PostgresClient {
	config := environmentConfiguration()

	var cs = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.host,
		config.port,
		config.user,
		config.database,
		config.password,
	)

	db, err := sql.Open("postgres", cs)

	if err != nil {
		log.Fatal(err)
	}

	return &PostgresClient{
		Connection: db,
	}
}
