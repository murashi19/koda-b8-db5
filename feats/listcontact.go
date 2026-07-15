package feats

import (
	"fmt"
	"minitask-db5/models"
	"minitask-db5/utils"

	"github.com/jackc/pgx/v5"
)

func ListContacts(db *pgx.Conn) {
	utils.ClearScreen()
	contacts, err := models.GetAllContact(db)
	if err != nil {
		fmt.Printf("Gagal mengambil data kontak: %v", err)
	}
	fmt.Println("\n--- DAFTAR KONTAK ---")
	if len(contacts) == 0 {
		fmt.Println("Tidak ada data kontak.")
		utils.EnterBack()
	} else {
		for _, c := range contacts {
			fmt.Printf(" ID: %d\n Nama: %s\n Email: %s\n Phone Number: %s\n", c.Id, c.Name, c.Email, c.Phone_number)
			fmt.Println("--------------------------")
			fmt.Println("")
		}
		utils.EnterBack()
	}
}
