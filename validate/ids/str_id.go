package ids

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
)

var strIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

func StrID(id string, table string, conn *sql.DB) error {
	if id == "" {
		return errors.New("ID cannot be empty.")
	}
	if len(id) > 255 {
		return fmt.Errorf("ID cannot exceed %d characters. (current %d)", 255, len(id))
	}
	if !strIDRegex.MatchString(id) {
		return errors.New("ID contains invalid characters.")
	}
	return checkExists(conn, table, "id = $1", id)
}
