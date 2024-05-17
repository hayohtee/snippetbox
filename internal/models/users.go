package models

import (
	"database/sql"
	"time"
)

// User a type to hold individual user.
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// UserModel a type which wraps database connection pool.
type UserModel struct {
	DB *sql.DB
}

// Insert - inserts a new record to the "users" table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate - verify whether a user exists with the
// provided email address and password. This will return
// the relevant user ID if they do or an error.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exists - check if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
