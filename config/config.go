package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "2030525_unecom"
	dbPass := "hDie244dq"
	dbName := "2030525_dbecom"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(207.246.248.54)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
