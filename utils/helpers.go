package utils

import (
	"database/sql"
	"fmt"
)

func CheckDuplicate(db *sql.DB, tableName, columnName, value string) error {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", tableName, columnName)

	var count int
	err := db.QueryRow(query, value).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking duplicates: %v", err)
	}

	if count > 0 {
		return fmt.Errorf("found an existing ID in the database: %s", value)
	}

	return nil
}
