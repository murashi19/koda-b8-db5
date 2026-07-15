package feats

import (
	"fmt"
	"minitask-db5/models"
	"minitask-db5/utils"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func EditContact(db *pgx.Conn) {
	utils.ClearScreen()
	fmt.Println("Update Contact")

	id, err := strconv.Atoi(utils.Input("Enter Contact ID: "))
	if err != nil {
		fmt.Printf("Invalid contact ID: %v\n", err)
		utils.EnterBack()
		return
	}

	data, err := inputContactData()
	if err != nil {
		fmt.Println(err)
		utils.EnterBack()
		return
	}

	data.Id = id

	contact, err := models.UpdateContact(data, db)
	if err != nil {
		fmt.Printf("Failed to update contact: %v\n", err)
		utils.EnterBack()
		return
	}

	fmt.Println("Contact Successfully Updated!")
	fmt.Printf("ID: %d\n", contact.Id)
	fmt.Printf("Name: %s\n", contact.Name)
	fmt.Printf("Email: %s\n", contact.Email)
	fmt.Printf("Phone: %s\n", contact.Phone_number)

	utils.EnterBack()
}
