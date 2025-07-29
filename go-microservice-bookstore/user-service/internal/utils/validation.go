package utils

import (
	"errors"
	"strings"

	"github.com/pratikchigani/user-service/internal/models"
)

func ValidateUser(u models.User) error {
	if strings.TrimSpace(u.ID) == "" {
		return errors.New("user ID is required")
	}
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("user name is required")
	}
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("user email is required")
	}
	return nil
}
