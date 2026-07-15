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
	rows, err := conn.Query(context.Background(), `
		SELECT id, name, email, phone_number, created_at, updated_at
		FROM contacts
	`)
	if err != nil {
		return nil, fmt.Errorf("get all contacts: %w", err)
	}
	defer rows.Close()

	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[Contact])
	if err != nil {
		return nil, fmt.Errorf("collect contacts: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return contacts, nil
}

func CreateContact(data Contact, conn *pgx.Conn) (Contact, error) {
	rows, err := conn.Query(context.Background(), `
		INSERT INTO contacts (name, email, phone_number)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, phone_number, created_at, updated_at
	`, data.Name, data.Email, data.Phone_number)
	if err != nil {
		return Contact{}, fmt.Errorf("insert contact: %w", err)
	}
	defer rows.Close()

	contact, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Contact])
	if err != nil {
		return Contact{}, fmt.Errorf("collect inserted contact: %w", err)
	}

	return contact, nil
}

func UpdateContact(data Contact, conn *pgx.Conn) (Contact, error) {
	rows, err := conn.Query(context.Background(), `
		UPDATE contacts
		SET
			name = $1,
			email = $2,
			phone_number = $3,
			updated_at = NOW()
		WHERE id = $4
		RETURNING id, name, email, phone_number, created_at, updated_at
	`, data.Name, data.Email, data.Phone_number, data.Id)
	if err != nil {
		return Contact{}, fmt.Errorf("update contact: %w", err)
	}
	defer rows.Close()

	contact, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Contact])
	if err != nil {
		return Contact{}, fmt.Errorf("collect updated contact: %w", err)
	}

	return contact, nil
}

func DeleteContact(id int, conn *pgx.Conn) error {
	commandTag, err := conn.Exec(context.Background(), `
		DELETE FROM contacts
		WHERE id = $1
	`, id)

	if err != nil {
		return fmt.Errorf("delete contact: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("contact with id %d not found", id)
	}

	return nil
}
