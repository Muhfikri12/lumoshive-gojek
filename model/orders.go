package model

import (
	"database/sql"
	"time"
)

type Orders struct {
	ID int
	Name string
	CustomerID Customers
	DriverID Drivers
	Created_at time.Time
	Updated_at sql.NullTime
	Deleted_at sql.NullTime
}

type MonthlyOrders struct {
	Month string
	TotalOrder int
}

type HighestOrderByCity struct {
	City string
	TotalOrders int
}

type HighestOrderByTime struct {
	Time string
	TotalOrders int
}

