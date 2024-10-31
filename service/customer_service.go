package service

import (
	"database/sql"
	"fmt"
	"golang-database/model"
	"golang-database/repository"
	"time"
)

func InputDataCustomer(db *sql.DB, name, address, email, password string) error {

	customerRepo := repository.NewCustomerRepository(db)
	user := model.Users{
		Email:    email,
		Password: password,
		IsStatus: true,
	}
	customer := model.Customers{
		Name:       name,
		Address:    address,
		Created_at: time.Now(),
	}

	// Memasukkan data customer dan user
	err := customerRepo.CreateCustomers(&customer, &user)
	if err != nil {
		return fmt.Errorf("gagal menambahkan data customer: %v", err)
	}

	fmt.Println("Berhasil menambahkan data customer dengan Id", customer.ID)
	return nil
}

func GetDataMonthlyCustomerService(customerRepo repository.CustomerRepository) error  {
	customers, err := customerRepo.MonthlyDataCustomers()
	if err != nil {
		return fmt.Errorf("gagal mengambil data customers: %v", err)
	}

	for _, customer := range customers {
		// fmt.Println("---Data Customer dengan pembelian terbanyak setiap bulannya---")
		fmt.Printf("\nName: %s,\nTotal Transaksi: %d,\nBulan: %s\n", customer.CustomerName, customer.Transaction,customer.Month)
		fmt.Println("=======================================")
	}

	return nil
}

func GetActiveCustomers(customerRepo repository.CustomerRepository) error  {
	customers, err := customerRepo.ActiveCustomers()
	if err != nil {
		return fmt.Errorf("gagal mengambil data customers: %v", err)
	}

	for _, customer := range customers {
		// fmt.Println("---Data Customer dengan pembelian terbanyak setiap bulannya---")
		fmt.Printf("\nNama: %s\n", customer.Name)
		fmt.Println("=======================================")
	}

	return nil
}