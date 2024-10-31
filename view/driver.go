package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateDriverMenu() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("Nama tidak boleh kosong!")
		return "","","",""
	}
	

	fmt.Print("Masukkan Alamat: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)
	if address == "" {
		fmt.Println("Alamat tidak boleh kosong!")
		return "","","",""
	}
	

	fmt.Print("Masukkan Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		fmt.Println("Email tidak boleh kosong!")
		return "","","",""
	}
	

	fmt.Print("Masukkan Kata Sandi: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	if password == "" {
		fmt.Println("Password tidak boleh kosong!")
	}
	

	return name, address, email, password
}