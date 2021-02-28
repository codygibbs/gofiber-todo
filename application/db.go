package application

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

// Open new database connection
func open() (*sql.DB, error) {
	connSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		getEnvInt("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	conn, err := sql.Open("postgres", connSource)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func getEnvInt(name string) int {
	val, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		val = 0
	}

	return val
}
