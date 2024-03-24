package service

import (
	"main/src/users/domain/model"
	appError "main/utils/error"
)

type UserService interface {
	CreateUser(user *model.User) (*model.User, *appError.Error)
	GetUserByID(userID string) (*model.User, *appError.Error)
	UpdateUserByID(userID string, user *model.User) (*model.User, *appError.Error)
	GetAllUser() ([]model.User, *appError.Error)
	DeleteUser(userID string) *appError.Error
}