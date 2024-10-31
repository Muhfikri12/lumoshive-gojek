package service

import (
	"database/sql"
	"fmt"
	"golang-database/model"
	"golang-database/repository"
)

func InputDataDriver(db *sql.DB, name, address, email, password string) error {

	driverRepo := repository.NewDriverRepository(db)
	user := model.Users{
		Email: email,
		Password: password,
		IsStatus: true,
	}

	driver := model.Drivers{
		Name: name,
		Address: address,
	}

	err := driverRepo.CreateAccountDriver(&driver, &user)
	if err != nil {
		return fmt.Errorf("gagal membuat data Driver: %v", err)
	}

	fmt.Println("Berhasil menambahkan data driver dengan ID", driver.ID)
	return nil
}

func GetHighestOrderByDriver(driverRepo repository.DriverRepository) error  {
	drivers, err := driverRepo.HighestGetOrder()
	if err != nil {
		return fmt.Errorf("gagal mengambil data customers: %v", err)
	}

	for _, driver := range drivers {
		fmt.Printf("\nName: %s,\nTotal Orderan: %d,\nBulan: %s\n", driver.Name, driver.TotalOrder, driver.Month)
		fmt.Println("=======================================")
	}

	return nil
}