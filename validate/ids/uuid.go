package ids

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

func UUID(id uuid.UUID, table string, conn *sql.DB) error {
	if id == uuid.Nil {
		return errors.New("UUID cannot be nil.")
	}
	return checkExists(conn, table, "id = $1", id)
}
