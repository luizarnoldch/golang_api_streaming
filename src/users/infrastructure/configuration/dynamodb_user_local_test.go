package configuration_test

import (
	"context"

	"os"
	"testing"

	"main/src/users/infrastructure/configuration"
	dynamodbUtils "main/utils/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/suite"
)

type DynamoDBUserConfigSuite struct {
	suite.Suite
	client    *dynamodb.Client
	ctx       context.Context
	tableName string
}

func (suite *DynamoDBUserConfigSuite) SetupSuite() {
	suite.ctx = context.TODO()
	var err error
	suite.client, err = dynamodbUtils.GetLocalDynamoDBClient(suite.ctx)
	suite.NoError(err)
	suite.tableName = configuration.GetDynamoDBUserTable()
}

func (suite *DynamoDBUserConfigSuite) TestGetDynamoDBStreamTableWithEnvSet() {
	testValue := "MyUserTable"
	os.Setenv("USER_TABLE", testValue)
	defer os.Unsetenv("USER_TABLE")
	result := configuration.GetDynamoDBUserTable()
	suite.Equal(testValue, result)
}

func (suite *DynamoDBUserConfigSuite) TestGetDynamoDBStreamTableWithoutEnvSet() {
	os.Unsetenv("USER_TABLE")
	result := configuration.GetDynamoDBUserTable()
	suite.Equal("Test_User_Table", result)
}

func (suite *DynamoDBUserConfigSuite) TestDeleteLocalDynamoDBStreamTable() {
	tableName := "New_User_Table_Testing"
	err := configuration.CreateLocalDynamoDBUserTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)

	exists, err := configuration.DescribeUserTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)
	suite.True(exists, "DescribeUserTable exists")

	err = configuration.DeleteLocalDynamoDBUserTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)
}

func (suite *DynamoDBUserConfigSuite) TestDeleteTableNotExistsDynamoTable() {
	tableName := "NonExistentTable"
	err := configuration.DeleteLocalDynamoDBUserTable(suite.ctx, suite.client, tableName)
	suite.Error(err)
	suite.Contains(err.Error(), "does not exist")
}

func (suite *DynamoDBUserConfigSuite) TestCreateLocalDynamoDBStreamTable() {
	tableName := "User_Table_Testing"
	err := configuration.CreateLocalDynamoDBUserTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)

	exists, err := configuration.DescribeUserTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)
	suite.True(exists, "DescribeUserTable exists")

}

func TestDynamoDBConfigSuite(t *testing.T) {
	suite.Run(t, new(DynamoDBUserConfigSuite))
}
