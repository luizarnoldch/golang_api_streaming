package model

import (
	appError "main/utils/error"
	"main/utils/validation"
)

type User struct {
	ID           string `json:"ID,omitempty" dynamodbav:"ID,omitempty"`
	Name         string `json:"name,omitempty" dynamodbav:"name,omitempty"`
	Email        string `json:"email" dynamodbav:"email"`
	LastActivity string `json:"last_activity,omitempty" dynamodbav:"last_activity,omitempty"`
	LastUpdate   string `json:"last_update,omitempty" dynamodbav:"last_update,omitempty"`
	CreatedAt    string `json:"created_at,omitempty" dynamodbav:"created_at,omitempty"`
}

func (u *User) Validate() *appError.Error {
	if err := validation.ValidateUUID(u.ID); err != nil {
		return err
	}
	if err := validation.ValidateStringField(u.Name, 2); err != nil {
		return err
	}
	if err := validation.ValidateEmail(u.Email); err != nil {
		return err
	}
	if err := validation.ValidateDate(u.CreatedAt); err != nil {
		return err
	}
	return nil
}
