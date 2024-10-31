package repository

import (
	"database/sql"
	"golang-database/model"
)

type DriverRepository interface {
	CreateAccountDriver(driver *model.Drivers, user *model.Users) error
	HighestGetOrder() ([]model.GetOrder, error)
}

type RepositoryDriver struct {
	DB *sql.DB
}

func NewDriverRepository(db *sql.DB) DriverRepository {
	return &RepositoryDriver{DB: db}
}

func (d *RepositoryDriver) CreateAccountDriver(driver *model.Drivers, user *model.Users) error {
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			tx.Rollback()
		}
	}()

	userQuery := `INSERT INTO users(email, password, is_status) VALUES($1, $2, $3) RETURNING ID`

	err = tx.QueryRow(userQuery, user.Email, user.Password, user.IsStatus).Scan(&user.ID)
	if err != nil {
		return err
	}

	driverQuery := `INSERT INTO drivers(name, address, user_id, created_at) VALUES($1, $2, $3, $4) RETURNING ID`

	err = tx.QueryRow(driverQuery, driver.Name, driver.Address, user.ID, driver.Created_at).Scan(&driver.ID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (god *RepositoryDriver) HighestGetOrder() ([]model.GetOrder, error) {
	query := `
		select 
			d.name,
			count(o.id) as total_orders,
			to_char(date_trunc('month', order_date), 'Month') AS month
		from
			drivers d 
		join
			orders o on d.id = o.driver_id
		where 
			o.deleted_at is null
		group by 
			d.id, month
		order by
			month, total_orders desc`

	rows, err := god.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orderDrivers []model.GetOrder

	for rows.Next() {
		var orderDriver model.GetOrder
		err := rows.Scan(&orderDriver.Name, &orderDriver.TotalOrder, &orderDriver.Month)
		if err != nil {
			return nil, err
		}
		orderDrivers = append(orderDrivers, orderDriver)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orderDrivers, nil
}