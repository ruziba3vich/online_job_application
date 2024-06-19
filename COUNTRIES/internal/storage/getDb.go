package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	return nil
}
