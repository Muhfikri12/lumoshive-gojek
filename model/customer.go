package model

import (
	"database/sql"
	"time"
)

type Customers struct {
	ID int
	Name string
	Address string
	Created_at time.Time
	Updated_at sql.NullTime
	Deleted_at sql.NullTime
}

type GetMonthlyDataCustomers struct {
	Month string
	CustomerName string
	Transaction int
}
