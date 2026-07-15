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
	rows, err := conn.Query(context.Background(), `
		INSERT INTO contacts (name, email, phone_number) VALUES
		($1, $2, $3)
		RETURNING id, name, email, phone_number, created_at, updated_at
	`, data.Name, data.Email, data.Phone_number)
	if err != nil {
		fmt.Println("Failed Insert Query")
	}
	contact, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Contact])
	if err != nil {
		fmt.Println("Failed Colleting Data Rows")
	}
	return contact, nil
}

func UpdateContact(data Contact, conn *pgx.Conn) (Contact, error) {
	rows, err := conn.Query(context.Background(), `
		UPDATE contacts SET
		name = $1,
		email = $2,
		phone_number = $3,
		updated_at = NOW()
		WHERE id = $4
		RETURNING id, name, email, phone_number, created_at, updated_at
	`, data.Name, data.Email, data.Phone_number, data.Id)
	if err != nil {
		fmt.Println("Failed Update Query")
	}
	contact, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Contact])
	if err != nil {
		fmt.Println("Failed Colleting Data Rows")
	}
	return contact, nil
}

func DeleteContact(id int, conn *pgx.Conn) error {
	commandTag, err := conn.Exec(context.Background(), `
		DELETE FROM contacts
		WHERE id = $1
	`, id)
	if err != nil {
		fmt.Println("Failed Delete Query")
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("Contact not found")
	}
	if err != nil {
		fmt.Println("Failed Colleting Data Rows")
	}
	return nil
}
