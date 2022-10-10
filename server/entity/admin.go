package entity

import (
	"database/sql"
	"time"
)

type Admin struct {
	ID        string
	Token     sql.NullString
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
