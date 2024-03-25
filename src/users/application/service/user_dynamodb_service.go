package service

import (
	"log"
	"main/src/users/domain/model"
	"main/src/users/domain/repository"
	appError "main/utils/error"
	"time"

	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (service *userServiceImpl) CreateUser(user *model.User) (*model.User, *appError.Error) {
	if user.ID != "" {
		user.ID = uuid.NewString()
	}
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.LastUpdate = time.Now().Format(time.RFC3339)
	user.LastActivity = time.Now().Format(time.RFC3339)
	if err := user.Validate(); err != nil {
		log.Printf("error while validating request from CreateStream Service: %v", err)
		return nil, err
	}
	return service.userRepo.CreateUser(user)
}

func (service *userServiceImpl) GetUserByID(userID string) (*model.User, *appError.Error) {
	return service.userRepo.GetUserByID(userID)
}

func (service *userServiceImpl) UpdateUserByID(userID string, user *model.User) (*model.User, *appError.Error) {
	return service.userRepo.UpdateUserByID(userID, user)
}

func (service *userServiceImpl) GetAllUser() ([]model.User, *appError.Error) {
	return service.userRepo.GetAllUser()
}

func (service *userServiceImpl) DeleteUser(userID string) *appError.Error {
	return service.userRepo.DeleteUser(userID)
}
