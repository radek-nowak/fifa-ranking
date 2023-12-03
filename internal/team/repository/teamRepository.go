package repository

import (
	"database/sql"
)

type TeamRepository struct {
	db *sql.DB
}
