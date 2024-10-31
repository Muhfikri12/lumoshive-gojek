package service

import (
	"fmt"
	"golang-database/repository"
)

func GetDataMonthlyOrderService(orderRepo repository.OrderRepository) error {
	
	orders, err := orderRepo.GetMonthlyOrder()
	if err != nil {
		return fmt.Errorf("gagal mengambil data Transaksi: %v", err)
	}

	for _, order := range orders {
		fmt.Printf("\nBulan: %s,\nTotal Transaksi: %d,\n", order.Month, order.TotalOrder)
		fmt.Println("=======================================")
	}

	return nil
}

func HighestOrderByCityService(orderRepo repository.OrderRepository) error {

	orders, err := orderRepo.GetHighestOrderByCity()
	if err != nil {
		return fmt.Errorf("gagal mengambil data transaksi: %v", err)
	}

	for _, order := range orders {
		fmt.Printf("\nKota: %s, Total Transaksi: %d,\n", order.City, order.TotalOrders)
		fmt.Println("=======================================")
	}

	return nil
	
}

func HighestOrderByTimeService(orderRepo repository.OrderRepository) error {

	orders, err := orderRepo.GetHighestOrderByTime()
	if err != nil {
		return fmt.Errorf("gagal mengambil data transaksi: %v", err)
	}

	for _, order := range orders {
		fmt.Printf("\nWaktu: %s, Total Transaksi: %d,\n", order.Time, order.TotalOrders)
		fmt.Println("=======================================")
	}

	return nil
	
}