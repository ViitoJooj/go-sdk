package ids

import (
	"database/sql"
	"errors"
)

func IntID(id int, table string, conn *sql.DB) error {
	if id <= 0 {
		return errors.New("ID must be a positive integer.")
	}
	return checkExists(conn, table, "id = $1", id)
}
