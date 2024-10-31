package repository

import (
	"database/sql"
	"golang-database/model"
)


type CustomerRepository interface {
	CreateCustomers(customer *model.Customers, user *model.Users) error
	MonthlyDataCustomers() ([]model.GetMonthlyDataCustomers, error)
	ActiveCustomers() ([]model.Customers, error)
}

type RepositoryDB struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &RepositoryDB{DB: db}
}

func (c *RepositoryDB) CreateCustomers(customer *model.Customers, user *model.Users) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}

	defer func () {
		if err != nil {
			tx.Rollback()
		}
	}()

	userQuery := `INSERT INTO users (email, password, is_status) VALUES($1, $2, $3) RETURNING ID`

	err = tx.QueryRow(userQuery, user.Email, user.Password, user.IsStatus).Scan(&user.ID)
	if err != nil {
		return err
	}

	CustomerQuery := `INSERT INTO customers (name, address, created_at, user_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(CustomerQuery, customer.Name, customer.Address, customer.Created_at, user.ID).Scan(&customer.ID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

// Repository For Get data monthly customers

func (c *RepositoryDB) MonthlyDataCustomers() ([]model.GetMonthlyDataCustomers, error) {
	query := `
		select to_char(date_trunc('month', o.order_date), 'Month') as month,
			c.name as customer_name,
			count(o.id) as total_transaction
		from 
			orders o
		join 
			customers c on o.customer_id = c.id 
		where 
			o.deleted_at is null 
		group by 
			month, c.name
		order by     
			total_transaction desc`

	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []model.GetMonthlyDataCustomers

	for rows.Next() {
		var transaction model.GetMonthlyDataCustomers
		err := rows.Scan(&transaction.Month, &transaction.CustomerName, &transaction.Transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (c *RepositoryDB) ActiveCustomers() ([]model.Customers, error) {
	query := `
		select 
			c.name
		from
			customers c 
		join
			users u on c.user_id = u.id 
		where 
			u.is_status = true`

	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []model.Customers

	for rows.Next() {
		var customer model.Customers
		err := rows.Scan(&customer.Name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}