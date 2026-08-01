package ids

import (
	"database/sql"
	"fmt"
)

func checkExists(conn *sql.DB, table, where string, arg interface{}) error {
	if conn == nil {
		return nil
	}
	if table == "" {
		return nil
	}
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s)", table, where)
	var exists bool
	err := conn.QueryRow(query, arg).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check ID existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("ID does not exist in %s.", table)
	}
	return nil
}
