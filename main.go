package main

import (
	"fmt"
	"golang-database/database"
	"golang-database/repository"
	"golang-database/service"
	"golang-database/utils"
	"golang-database/view"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	customerRepo := repository.NewCustomerRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	driverRepo := repository.NewDriverRepository(db)

	for {
		var option int
		utils.ClearScreen()
		fmt.Println("=== Selamat Datang ===")
		fmt.Println("1. Membuat Data Customer")
		fmt.Println("2. Customer Sering Berbelanja")
		fmt.Println("3. Membuat Data Driver")
		fmt.Println("4. Data Bulanan Pesanan")
		fmt.Println("5. Data Transaksi Berdasarkan Kota")
		fmt.Println("6. Data Transaksi Berdasarkan Waktu")
		fmt.Println("7. Customer Aktif")
		fmt.Println("8. Driver Order Terbanyak")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih opsi: ")

		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Input tidak valid. Silakan masukkan angka.")
			continue
		}

		switch option {
		case 1:
			utils.ClearScreen()
			name, address, email, password := view.CreateCustomerMenu()
			if name == "" || address == "" || email == "" || password == "" {
				fmt.Println("Input data tidak valid, proses dibatalkan.")
			} else {
				err = service.InputDataCustomer(db, name, address, email, password)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Data customer berhasil ditambahkan.")
			}

		case 2:
			utils.ClearScreen()
			err := service.GetDataMonthlyCustomerService(customerRepo)
			if err != nil {
				log.Println("Error:", err)
			}

		case 3:
			utils.ClearScreen()
			name, address, email, password := view.CreateDriverMenu()
			if name == "" || address == "" || email == "" || password == "" {
				fmt.Println("Input data tidak valid, proses dibatalkan.")
			} else {
				err = service.InputDataDriver(db, name, address, email, password)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Data Driver berhasil ditambahkan.")
			}

		case 4:
			utils.ClearScreen()
			err := service.GetDataMonthlyOrderService(orderRepo)
			if err != nil {
				log.Println("Error:", err)
			}

		case 5:
			utils.ClearScreen()
			err := service.HighestOrderByCityService(orderRepo)
			if err != nil {
				log.Println("Error:", err)
			}

		case 6:
			utils.ClearScreen()
			err := service.HighestOrderByTimeService(orderRepo)
			if err != nil {
				log.Println("Error:", err)
			}

		case 7:
			utils.ClearScreen()
			err := service.GetActiveCustomers(customerRepo)
			if err != nil {
				log.Println("Error:", err)
			}

		case 8:
			utils.ClearScreen()
			err := service.GetHighestOrderByDriver(driverRepo)
			if err != nil {
				log.Println("Error:", err)
			}

		case 0:
			utils.ClearScreen()
			fmt.Println("Keluar dari aplikasi.")
			return

		default:
			utils.ClearScreen()
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}

		// Menunggu pengguna menekan Enter sebelum kembali ke menu
		fmt.Println("\nTekan Enter untuk kembali ke menu...")
		fmt.Scanln() // Tunggu pengguna menekan Enter
	}
}

