package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a user repository
type Users struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create inserts a new user into the database table
// Struct 	   Method Name				Return
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastUserID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastUserID), nil
}
