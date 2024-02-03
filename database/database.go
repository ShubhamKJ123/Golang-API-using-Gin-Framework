// database/database.go
package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}
