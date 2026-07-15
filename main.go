package main

import (
	"fmt"
	"minitask-db5/lib"
	"minitask-db5/models"
	"strconv"
)

func main() {

	var err error
	db, err := lib.Conn()
	if err != nil {
		fmt.Println("Gagal terhubung ke database:", err)
	}

	fmt.Println("Berhasil terhubung ke database!")

	contacts, _ := models.GetAllContact(db)
	if err != nil {
		fmt.Printf("Gagal mengambil data kontak: %v", err)
	}

	fmt.Println("----- Welcome to Contact List App -----")
	fmt.Println("\n1. List Contacts")
	fmt.Println("2. Add Contact")
	fmt.Println("3. Update Contact")
	fmt.Println("4. Delete Contact")

	choose, err := strconv.Atoi(lib.Input("\nChoose Menu: "))
	if err != nil {
		fmt.Println("Invalid Choose Menu, Please Enter a Number")
	}

	switch choose {
	case 1:
		lib.ClearScreen()
		fmt.Println("\n--- DAFTAR KONTAK ---")
		if len(contacts) == 0 {
			fmt.Println("Tidak ada data kontak.")
			lib.EnterBack()
		} else {
			for _, c := range contacts {
				fmt.Printf(" ID: %d\n Nama: %s\n Email: %s\n Telepon: %s\n", c.Id, c.Name, c.Email, c.Phone_number)
			}
			lib.EnterBack()
		}
	case 2:

	}

}
