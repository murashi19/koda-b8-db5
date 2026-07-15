package main

import (
	"context"
	"fmt"
	"minitask-db5/feats"
	"minitask-db5/lib"
	"minitask-db5/utils"
	"strconv"
	"time"
)

func main() {

	db, err := lib.Conn()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}

	defer db.Close(context.Background())

	fmt.Println("Successfully connected to database!")

	done := make(chan struct{})
	go utils.Loading(done, "Data is being processed...")

	time.Sleep(2 * time.Second)
	close(done)

	for {
		utils.ClearScreen()
		fmt.Println("----- Welcome to Contact List App -----")
		fmt.Println("\n1. List Contacts")
		fmt.Println("2. Add Contact")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("\n0. Exit")

		choose, err := strconv.Atoi(utils.Input("\nChoose Menu: "))
		if err != nil {
			fmt.Println("Invalid Choose Menu, Please Enter a Number")
			time.Sleep(1 * time.Second)
			continue
		}

		switch choose {
		case 1:
			feats.ListContacts(db)
		case 2:
			feats.AddContact(db)
		case 3:
			feats.EditContact(db)
		case 4:
			feats.DeletedContact(db)
		case 0:
			fmt.Println("Thank you for using Contact List App.")
			return
		default:
			fmt.Println("Please fill according to the menu")
			time.Sleep(1 * time.Second)
		}
	}

}
