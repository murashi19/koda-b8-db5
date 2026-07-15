package feats

import (
	"fmt"
	"minitask-db5/models"
	"minitask-db5/utils"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

func AddContact(db *pgx.Conn) {
	utils.ClearScreen()
	fmt.Println("Create New Contact")
	name := strings.TrimSpace(utils.Input("Enter your Name: "))
	email := strings.TrimSpace(utils.Input("Enter your Email: "))
	phone := strings.TrimSpace(utils.Input("Enter your Phone Number: "))

	data := models.Contact{
		Name:         name,
		Email:        email,
		Phone_number: phone,
	}

	createInput, err := models.CreateContact(data, db)
	if err != nil {
		fmt.Println("Failed Create New Contact")
		return
	}

	fmt.Println("Contact Successfuly Saved!")
	fmt.Printf("ID: %d\n", createInput.Id)
	fmt.Printf("Name: %s\n", createInput.Name)
	fmt.Printf("Email: %s\n", createInput.Email)
	fmt.Printf("Phone Number: %s\n", createInput.Phone_number)
	time.Sleep(2 * time.Second)
	utils.EnterBack()
}
