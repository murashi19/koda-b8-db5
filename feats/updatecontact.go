package feats

import (
	"fmt"
	"minitask-db5/models"
	"minitask-db5/utils"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

func EditContact(db *pgx.Conn) {
	utils.ClearScreen()
	fmt.Println("Update Contact")
	id, err := strconv.Atoi(utils.Input("Enter Contact ID: "))
	if err != nil {
		fmt.Println("Invalid Contact ID")
		return
	}
	name := strings.TrimSpace(utils.Input("Enter your Name: "))
	email := strings.TrimSpace(utils.Input("Enter your Email: "))
	phone := strings.TrimSpace(utils.Input("Enter your Phone Number: "))

	data := models.Contact{
		Id:           id,
		Name:         name,
		Email:        email,
		Phone_number: phone,
	}

	updateInput, err := models.UpdateContact(data, db)
	if err != nil {
		fmt.Println("Failed Update Contact")
		utils.EnterBack()
		return
	}

	fmt.Println("Contact Successfuly Saved!")
	fmt.Printf("ID: %d\n", updateInput.Id)
	fmt.Printf("Name: %s\n", updateInput.Name)
	fmt.Printf("Email: %s\n", updateInput.Email)
	fmt.Printf("Phone Number: %s\n", updateInput.Phone_number)
	time.Sleep(2 * time.Second)
	utils.EnterBack()
}
