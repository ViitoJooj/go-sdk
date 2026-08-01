package ids

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
)

var uuidv7Regex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

func UUIDv7(id string, table string, conn *sql.DB) error {
	if id == "" {
		return errors.New("UUIDv7 cannot be empty.")
	}
	if len(id) != 36 {
		return fmt.Errorf("UUIDv7 must be exactly 36 characters. (current %d)", len(id))
	}
	if !uuidv7Regex.MatchString(id) {
		return errors.New("invalid UUIDv7 format.")
	}
	if conn != nil && table != "" {
		return checkExists(conn, table, "id = $1", id)
	}
	return nil
}
