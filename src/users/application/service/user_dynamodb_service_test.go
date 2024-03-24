package service_test

import (
	"testing"

	userMock "main/mocks" // Update this path to the correct location of your mocks
	"main/src/users/application/service"
	"main/src/users/domain/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserServiceDynamoDBSuite struct {
	suite.Suite
	userRepository          *userMock.UserRepository
	userServiceApplication  service.UserService
}

const (
	MethodCreateUser       = "CreateUser"
	MethodGetUserByID      = "GetUserByID"
	MethodUpdateUserByID   = "UpdateUserByID"
	MethodGetAllUser       = "GetAllUser"
	MethodDeleteUser       = "DeleteUser"
)

func (suite *UserServiceDynamoDBSuite) SetupTest() {
	suite.userRepository = new(userMock.UserRepository)
	suite.userServiceApplication = service.NewUserService(suite.userRepository)
}

func (suite *UserServiceDynamoDBSuite) TestCreateUser() {
	newUser := &model.User{
		Name:      "Jane Doe",
		Email:     "jane.doe@example.com",
		CreatedAt: "2023-01-01T00:00:00Z",
	}
	suite.userRepository.On(MethodCreateUser, newUser).Return(newUser, nil)
	createdUser, err := suite.userServiceApplication.CreateUser(newUser)
	suite.Nil(err, "Creating a user should not return an error")
	suite.NotNil(createdUser, "Created user should not be nil")
	suite.NotEqual("", createdUser.ID, "Created user should have a non-empty ID")
	suite.Equal(newUser.Name, createdUser.Name, "Created user name should match the input")
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserServiceDynamoDBSuite) TestGetUserByID() {
	userID := uuid.NewString()
	expectedUser := &model.User{
		ID:   userID,
		Name: "Jane Doe",
	}
	suite.userRepository.On(MethodGetUserByID, userID).Return(expectedUser, nil)
	user, err := suite.userServiceApplication.GetUserByID(userID)
	suite.Nil(err, "Getting a user by ID should not return an error")
	suite.Equal(expectedUser.ID, user.ID, "The returned user ID should match the expected ID")
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserServiceDynamoDBSuite) TestUpdateUserByID() {
    userID := uuid.NewString()
    updatedUser := &model.User{
        ID:    userID,
        Name:  "Updated Name",
        Email: "updated@example.com",
    }
    suite.userRepository.On(MethodUpdateUserByID, userID, updatedUser).Return(updatedUser, nil)
    resultUser, err := suite.userServiceApplication.UpdateUserByID(userID, updatedUser)
    suite.Nil(err, "Updating a user should not return an error")
    suite.NotNil(resultUser, "Updated user should not be nil")
    suite.Equal(updatedUser.Name, resultUser.Name, "Updated user name should match")
    suite.Equal(updatedUser.Email, resultUser.Email, "Updated user email should match")
    suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserServiceDynamoDBSuite) TestGetAllUser() {
    expectedUsers := []model.User{
        {ID: uuid.NewString(), Name: "User 1", Email: "user1@example.com"},
        {ID: uuid.NewString(), Name: "User 2", Email: "user2@example.com"},
    }
    suite.userRepository.On(MethodGetAllUser).Return(expectedUsers, nil)
    users, err := suite.userServiceApplication.GetAllUser()
    suite.Nil(err, "Retrieving all users should not return an error")
    suite.Len(users, len(expectedUsers), "The number of retrieved users should match")
    suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserServiceDynamoDBSuite) TestDeleteUser() {
    userID := uuid.NewString()
    suite.userRepository.On(MethodDeleteUser, userID).Return(nil)
    err := suite.userServiceApplication.DeleteUser(userID)
    suite.Nil(err, "Deleting a user should not return an error")
    suite.userRepository.AssertExpectations(suite.T())
}

func TestUserServiceDynamoDBSuite(t *testing.T) {
	suite.Run(t, new(UserServiceDynamoDBSuite))
}
