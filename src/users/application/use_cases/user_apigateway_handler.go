package usecases

import (
	"context"
	"log"
	"main/src/users/application/service"
	"main/src/users/domain/model"
	"main/src/users/infrastructure/adapter"
	dynamodbUtils "main/utils/dynamodb"
	appError "main/utils/error"
)

type UserUseCases struct {
	Ctx       context.Context
	TableName string
}

func (user *UserUseCases) UpdateUserByID(userID string, req *model.User) (*model.User, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(user.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from UpdateUserByID API Function")
		return nil, appError.NewUnexpectedError(err.Error())
	}

	userInfrastructure := adapter.NewUserDynamoDBRepository(user.Ctx, dynamoDBClient, user.TableName)
	userService := service.NewUserService(userInfrastructure)

	return userService.UpdateUserByID(userID, req)
}

func (user *UserUseCases) GetUserByID(userID string) (*model.User, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(user.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from GetUserByID API Function")
		return nil, appError.NewUnexpectedError(err.Error())
	}

	userInfrastructure := adapter.NewUserDynamoDBRepository(user.Ctx, dynamoDBClient, user.TableName)
	userService := service.NewUserService(userInfrastructure)

	return userService.GetUserByID(userID)
}

func (user *UserUseCases) GetAllUser() ([]model.User, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(user.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from GetAllUser API Function")
		return nil, appError.NewUnexpectedError(err.Error())
	}

	userInfrastructure := adapter.NewUserDynamoDBRepository(user.Ctx, dynamoDBClient, user.TableName)
	userService := service.NewUserService(userInfrastructure)

	return userService.GetAllUser()
}

func (user *UserUseCases) CreateUser(req *model.User) (*model.User, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(user.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from CreateUser API Function")
		return nil, appError.NewUnexpectedError(err.Error())
	}

	userInfrastructure := adapter.NewUserDynamoDBRepository(user.Ctx, dynamoDBClient, user.TableName)
	userService := service.NewUserService(userInfrastructure)

	return userService.CreateUser(req)
}

func (user *UserUseCases) DeleteUser(userID string) *appError.Error {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(user.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from CreateStream API Function")
		return appError.NewUnexpectedError(err.Error())
	}

	userInfrastructure := adapter.NewUserDynamoDBRepository(user.Ctx, dynamoDBClient, user.TableName)
	userService := service.NewUserService(userInfrastructure)

	return userService.DeleteUser(userID)
}
