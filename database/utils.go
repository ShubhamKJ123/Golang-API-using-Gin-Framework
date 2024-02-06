// database/utils.go
package database

import (
	"database/sql"
)

// Helper function to execute a SQL query and return the result.
func QueryDB(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Helper function to execute a SQL statement and return the result.
func ExecDB(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}
