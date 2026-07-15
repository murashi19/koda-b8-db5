package feats

import (
	"fmt"
	"minitask-db5/models"
	"minitask-db5/utils"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

func DeletedContact(db *pgx.Conn) {
	utils.ClearScreen()
	fmt.Println("Delete Contact")
	id, err := strconv.Atoi(utils.Input("Enter Contact ID: "))
	if err != nil {
		fmt.Println("Invalid Contact ID")
		return
	}

	err = models.DeleteContact(id, db)
	if err != nil {
		fmt.Println("Failed Update Contact")
		utils.EnterBack()
		return
	}

	fmt.Println("Contact Successfuly Deleted ID")
	time.Sleep(2 * time.Second)
	utils.EnterBack()
}
