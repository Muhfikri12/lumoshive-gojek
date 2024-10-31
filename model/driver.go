package model

import (
	"database/sql"
	"time"
)

type Drivers struct {
	ID int
	Name string
	Address string
	Created_at time.Time
	Updated_at sql.NullTime
	Deleted_at sql.NullTime
}

type GetOrder struct {
	Name string
	Month string
	TotalOrder int
}