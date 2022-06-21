package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "testing"
	dbURI := "(127.0.0.1:3306)"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp"+dbURI+"/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return db
}
