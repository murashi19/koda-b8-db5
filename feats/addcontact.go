package feats

import (
	"fmt"
	"minitask-db5/models"
	"minitask-db5/utils"

	"github.com/jackc/pgx/v5"
)

func AddContact(db *pgx.Conn) {
	utils.ClearScreen()
	fmt.Println("Create New Contact")

	data, err := inputContactData()
	if err != nil {
		fmt.Println(err)
		utils.EnterBack()
		return
	}

	contact, err := models.CreateContact(data, db)
	if err != nil {
		fmt.Printf("Failed to create contact: %v\n", err)
		utils.EnterBack()
		return
	}

	fmt.Println("Contact Successfully Saved!")
	fmt.Printf("ID: %d\n", contact.Id)
	fmt.Printf("Name: %s\n", contact.Name)
	fmt.Printf("Email: %s\n", contact.Email)
	fmt.Printf("Phone: %s\n", contact.Phone_number)

	utils.EnterBack()
}
