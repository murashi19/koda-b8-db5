package feats

import (
	"fmt"
	"minitask-db5/models"
	"minitask-db5/utils"
	"strings"
)

func inputContactData() (models.Contact, error) {
	name := strings.TrimSpace(utils.Input("Enter your Name: "))
	email := strings.TrimSpace(utils.Input("Enter your Email: "))
	phone := strings.TrimSpace(utils.Input("Enter your Phone Number: "))

	if name == "" {
		return models.Contact{}, fmt.Errorf("name cannot be empty")
	}

	if email == "" {
		return models.Contact{}, fmt.Errorf("email cannot be empty")
	}

	if phone == "" {
		return models.Contact{}, fmt.Errorf("phone number cannot be empty")
	}

	return models.Contact{
		Name:         name,
		Email:        email,
		Phone_number: phone,
	}, nil
}
