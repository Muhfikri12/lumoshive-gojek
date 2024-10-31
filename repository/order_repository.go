package repository

import (
	"database/sql"
	"golang-database/model"
)

type OrderRepository interface {
	GetMonthlyOrder() ([]model.MonthlyOrders, error)
	GetHighestOrderByCity() ([]model.HighestOrderByCity, error)
	GetHighestOrderByTime() ([]model.HighestOrderByTime, error)
}

type RepositoryOrder struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository  {
	return &RepositoryOrder{DB: db}
}

func (o *RepositoryOrder) GetMonthlyOrder() ([]model.MonthlyOrders, error) {
	query := `
		SELECT 
		to_char(date_trunc('month', order_date), 'Month') AS month,
		count(id) AS total_orders
		FROM
			orders
		WHERE
			deleted_at IS NULL 
		GROUP BY 
			date_trunc('month', order_date)
		ORDER BY 
			date_trunc('month', order_date)`

	rows, err := o.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactionOrders []model.MonthlyOrders

	for rows.Next() {
		var transactionOrder model.MonthlyOrders

		err := rows.Scan(&transactionOrder.Month, &transactionOrder.TotalOrder)
		if err != nil {
			return nil, err
		}
		transactionOrders = append(transactionOrders, transactionOrder)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactionOrders, nil
}

func (ho *RepositoryOrder) GetHighestOrderByCity() ([]model.HighestOrderByCity, error) {
	query := `
		SELECT
		split_part(address, ',', 2) as city,
		COUNT(o.id) as total_purchases
		FROM
			orders o 
		JOIN
			customers c on o.customer_id = c.id 
		WHERE 
			o.deleted_at IS NULL 
		GROUP BY 
			city
		ORDER BY 
			total_purchases DESC`

	rows, err := ho.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []model.HighestOrderByCity

	for rows.Next() {
		var order model.HighestOrderByCity

		err := rows.Scan(&order.City, &order.TotalOrders)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (ho *RepositoryOrder) GetHighestOrderByTime() ([]model.HighestOrderByTime, error) {
	query := `
		select  
			to_char(date_trunc('hour', created_at), 'HH24:MI:SS') as order_time,
			count(id) as total_orders
		from  
			orders
		where  
			deleted_at is null 
		group by  
			date_trunc('hour', created_at)
		order by  
			total_orders desc`

	rows, err := ho.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []model.HighestOrderByTime

	for rows.Next() {
		var order model.HighestOrderByTime

		err := rows.Scan(&order.Time, &order.TotalOrders)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

