package model

import (
	appError "main/utils/error"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID        string `json:"ID,omitempty" dynamodbav:"ID,omitempty"`
	Name      string `json:"name" dynamodbav:"name"`
	Email     string `json:"email" dynamodbav:"email"`
	CreatedAt string `json:"created_at,omitempty" dynamodbav:"created_at,omitempty"`
}

func (u *User) Validate() *appError.Error {
	if err := ValidateUUID(u.ID); err != nil {
		return err
	}
	if err := ValidateStringField(u.Name, 2); err != nil {
		return err
	}
	if err := ValidateEmail(u.Email); err != nil {
		return err
	}
	if err := ValidateDate(u.CreatedAt); err != nil {
		return err
	}
	return nil
}

func ValidateUUID(id string) *appError.Error {
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	if !uuidRegex.MatchString(id) {
		return appError.NewValidationError("Invalid ID: The ID must be a valid UUID.")
	}
	return nil
}

func ValidateStringField(field string, minLength int) *appError.Error {
	if len(strings.TrimSpace(field)) < minLength {
		return appError.NewValidationError("Field length is less than the required minimum")
	}
	return nil
}

func ValidateEmail(email string) *appError.Error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return appError.NewValidationError("Invalid email format")
	}
	return nil
}

func ValidateDate(dateStr string) *appError.Error {
	_, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return appError.NewValidationError("Invalid date format, must be RFC3339")
	}
	return nil
}