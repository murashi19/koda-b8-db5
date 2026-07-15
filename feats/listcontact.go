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
		fmt.Printf("Failed to get contacts: %v\n", err)
		utils.EnterBack()
		return
	}

	fmt.Println("\n--- CONTACT LIST ---")

	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		utils.EnterBack()
		return
	}

	for _, c := range contacts {
		fmt.Printf(
			"ID: %d\nName: %s\nEmail: %s\nPhone Number: %s\n",
			c.Id,
			c.Name,
			c.Email,
			c.Phone_number,
		)
		fmt.Println("--------------------------")
	}

	utils.EnterBack()
}
