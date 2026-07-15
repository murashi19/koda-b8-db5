package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Contact struct {
	Id           int
	Name         string
	Email        string
	Phone_number string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GetAllContact(conn *pgx.Conn) ([]Contact, error) {

	rows, err := conn.Query(context.Background(), `SELECT id, name, email, phone_number, created_at, updated_at FROM contacts`)
	if err != nil {
		fmt.Println("Cannot read a data")
	}

	defer rows.Close()

	contact, err := pgx.CollectRows(rows, pgx.RowToStructByName[Contact])

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return contact, nil
}

func CreateContact(data Contact, conn *pgx.Conn) (Contact, error) {
	rows, _ := conn.Query(context.Background(), `
		INSERT INTO contacts (name, email, phone_number) VALUES
		('$1', '$2', '$3')
		RETURNING id, name, email, phone_number, created_at, updated_at
	`, data.Id, data.Name, data.Email, data.Phone_number, data.CreatedAt, data.UpdatedAt)

	contact, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Contact])
	if err != nil {
		fmt.Println("Failed Colleting Data Rows")
	}
	return contact, nil
}
