package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/msanatan/gopractice/clicontacts/models"
)

type SQLite struct {
	DB *sql.DB
}

func (s *SQLite) InitDB() error {
	sqlitedb, err := sql.Open("sqlite3", "./contacts.db")
	if err != nil {
		return err
	}
	s.DB = sqlitedb // Set up the struct's DB
	return s.createContactsTable()
}

func (s *SQLite) createContactsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS contacts (
			id INTEGER PRIMARY KEY AUTOINCREMENT
			name TEXT
			email TEXT
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := s.DB.Exec(query)
	return err
}

func (s *SQLite) CreateContact(contact models.Contact) error {
	query := `INSERT INTO contacts (name, email) VALUES (?, ?)`
	_, err := s.DB.Exec(query, contact.Name, contact.Email)
	return err
}