package config

import (
	"database/sql"
	"fmt"
	"strconv"
)

func DbConnection() (*sql.DB, error) {

	dbServer := GetEnv("RBE_DB_SERVER", "localhost")
	dbPort := GetEnv("RBE_DB_PORT", "5432")
	dbUser := GetEnv("RBE_DB_USER", "rbe")
	dbPassword := GetEnv("RBE_DB_PASSWORD", "123456")
	dbName := GetEnv("RBE_DB_NAME", "rbe")
	dbSSLMode := GetEnv("RBE_DB_SSLMODE", "disable")
	dbMaxConns, _ := strconv.Atoi(GetEnv("RBE_DB_MAX_CONNS", "2"))
	dbMaxIdleConns, _ := strconv.Atoi(GetEnv("RBE_DB_MAX_IDLE_CONNS", "2"))

	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbServer, dbPort, dbUser, dbPassword, dbName, dbSSLMode)
	db, err := sql.Open("postgres", connectString)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)

	return db, nil
}
