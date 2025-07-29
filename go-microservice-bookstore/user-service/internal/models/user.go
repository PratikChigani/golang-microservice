// internal/models/user.go
package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/pratikchigani/user-service/internal/db"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserConflict = errors.New("user already exists")
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers() ([]User, error) {
	rows, err := db.DB.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if users == nil {
		users = []User{}
	}
	return users, nil
}

func GetUserByID(id string) (*User, error) {
	var u User
	err := db.DB.QueryRow(`SELECT id, name, email FROM users WHERE id = ?`, id).
		Scan(&u.ID, &u.Name, &u.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

func CreateUser(u User) error {
	_, err := GetUserByID(u.ID)
	if err == nil {
		return ErrUserConflict
	} else if !errors.Is(err, ErrUserNotFound) {
		return fmt.Errorf("check existence: %w", err)
	}

	_, err = db.DB.Exec(`INSERT INTO users (id, name, email) VALUES (?, ?, ?)`, u.ID, u.Name, u.Email)
	return err
}

func UpdateUser(id string, u User) error {
	res, err := db.DB.Exec(`UPDATE users SET name = ?, email = ? WHERE id = ?`, u.Name, u.Email, id)
	if err != nil {
		return fmt.Errorf("update users [%s]: %w", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected for update: %w", err)
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	log.Printf("[DB] User [%s] updated", id)
	return nil
}

func DeleteUser(id string) error {
	res, err := db.DB.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}
