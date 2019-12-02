package pq

import (
	"github.com/jackc/pgconn"
)

func Duplicated(err error) bool {
	// unique_violation
	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
		return true
	}

	return false
}