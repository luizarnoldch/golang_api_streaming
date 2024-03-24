package adapter_test

import (
	"context"
	"main/src/users/domain/model"
	"main/src/users/domain/repository"
	"main/src/users/infrastructure/adapter"
	"main/src/users/infrastructure/configuration"

	"testing"

	dynamodbUtils "main/utils/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserDynamoDBSuite struct {
	suite.Suite
	tableName                   string
	initUsers                   []model.User
	dynamoClient                *dynamodb.Client
	userRepository              repository.UserRepository
}

func (suite *UserDynamoDBSuite) SetupSuite() {
	ctx := context.TODO()
	client, err := dynamodbUtils.GetLocalDynamoDBClient(ctx)
	suite.NoError(err)
	suite.dynamoClient = client

	table_name := configuration.GetDynamoDBUserTable()
	suite.Equal("Test_User_Table", table_name)
	suite.tableName = table_name

	user_infrastructure := adapter.NewUserDynamoDBRepository(ctx, client, table_name)
	suite.userRepository = user_infrastructure

	configuration.CreateLocalDynamoDBUserTable(ctx, suite.dynamoClient, suite.tableName)

	suite.initUsers = []model.User{
		{
			ID:        uuid.NewString(),
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: "2022-01-01T00:00:00Z",
		},
		{
			ID:        uuid.NewString(),
			Name:      "Alice Smith",
			Email:     "alice.smith@example.com",
			CreatedAt: "2022-03-01T00:00:00Z",
		},
		{
			ID:        uuid.NewString(),
			Name:      "Bob Johnson",
			Email:     "bob.johnson@example.com",
			CreatedAt: "2022-04-01T00:00:00Z",
		},
	}

	exists, err := configuration.DescribeUserTable(ctx, suite.dynamoClient, suite.tableName)
	suite.NoError(err)
	suite.True(exists)
}

func (suite *UserDynamoDBSuite) TearDownSuite() {
	for _, user := range suite.initUsers {
		err := suite.userRepository.DeleteUser(user.ID)
		suite.Nil(err)
	}
}

func (suite *UserDynamoDBSuite) TestCreateUserSuccessful() {
	for _, user := range suite.initUsers {
		_, err := suite.userRepository.CreateUser(&user)
		suite.Nil(err)
	}
}

func (suite *UserDynamoDBSuite) TestGetAllUserSuccessful() {
	retrievedUsers, err := suite.userRepository.GetAllUser()
	suite.Nil(err)
	suite.NotEmpty(retrievedUsers)
	suite.Equal(len(suite.initUsers), len(retrievedUsers))
}

func (suite *UserDynamoDBSuite) TestGetUserByIDSuccessful() {
	testUser := suite.initUsers[0]
	retrievedUser, err := suite.userRepository.GetUserByID(testUser.ID)
	suite.Nil(err)
	suite.NotNil(retrievedUser)
	suite.NotEmpty(retrievedUser)
	suite.Equal(testUser.ID, retrievedUser.ID)
	suite.Equal(testUser.Name, retrievedUser.Name)
	suite.Equal(testUser.Email, retrievedUser.Email)
	suite.Equal(testUser.CreatedAt, retrievedUser.CreatedAt)
}

func (suite *UserDynamoDBSuite) TestUpdateUserByIDSuccessful() {
	testUser := suite.initUsers[0]
	testUser.Name = "Updated Name"
	testUser.Email = "updated.email@example.com"

	updatedUser, err := suite.userRepository.UpdateUserByID(testUser.ID, &testUser)
	suite.Nil(err)
	suite.NotNil(updatedUser)
	suite.NotEmpty(updatedUser)
	suite.Equal("Updated Name", updatedUser.Name)
	suite.Equal("updated.email@example.com", updatedUser.Email)
}

func TestUserDynamoDBSuite(t *testing.T) {
	suite.Run(t, new(UserDynamoDBSuite))
}
