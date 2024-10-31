package model

import (
	"database/sql"
	"time"
)

type Users struct {
	ID int
	Email string
	Password string
	IsStatus bool
	Created_at time.Time
	Updated_at sql.NullTime
	Deleted_at sql.NullTime
}