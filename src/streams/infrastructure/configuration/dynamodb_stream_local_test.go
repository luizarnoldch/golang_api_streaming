package configuration_test

import (
	"context"

	"os"
	"testing"

	"main/src/streams/infrastructure/configuration"
	dynamodbUtils "main/utils/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/suite"
)

type DynamoDBConfigSuite struct {
	suite.Suite
	client    *dynamodb.Client
	ctx       context.Context
	tableName string
}

func (suite *DynamoDBConfigSuite) SetupSuite() {
	suite.ctx = context.TODO()
	var err error
	suite.client, err = dynamodbUtils.GetLocalDynamoDBClient(suite.ctx)
	suite.NoError(err)
	suite.tableName = configuration.GetDynamoDBStreamTable()
}

func (suite *DynamoDBConfigSuite) TestGetDynamoDBStreamTableWithEnvSet() {
	testValue := "MyStreamTable"
	os.Setenv("STREAM_TABLE", testValue)
	defer os.Unsetenv("STREAM_TABLE")
	result := configuration.GetDynamoDBStreamTable()
	suite.Equal(testValue, result)
}

func (suite *DynamoDBConfigSuite) TestGetDynamoDBStreamTableWithoutEnvSet() {
	os.Unsetenv("STREAM_TABLE")
	result := configuration.GetDynamoDBStreamTable()
	suite.Equal("Test_Stream_Table", result)
}

func (suite *DynamoDBConfigSuite) TestDeleteLocalDynamoDBStreamTable() {
	tableName := "New_Stream_Table_Testing"
	err := configuration.CreateLocalDynamoDBStreamTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)

	exists, err := configuration.DescribeStreamTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)
    suite.True(exists, "DescribeStreamTable exists")

	err = configuration.DeleteLocalDynamoDBStreamTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)
}

func (suite *DynamoDBConfigSuite) TestDeleteTableNotExistsDynamoTable() {
	tableName := "NonExistentTable"
	err := configuration.DeleteLocalDynamoDBStreamTable(suite.ctx, suite.client, tableName)
	suite.Error(err)
	suite.Contains(err.Error(), "does not exist")
}

func (suite *DynamoDBConfigSuite) TestCreateLocalDynamoDBStreamTable() {
	tableName := "Stream_Table_Testing"
	err := configuration.CreateLocalDynamoDBStreamTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)

	exists, err := configuration.DescribeStreamTable(suite.ctx, suite.client, tableName)
	suite.NoError(err)
    suite.True(exists, "DescribeStreamTable exists")

}

func TestDynamoDBConfigSuite(t *testing.T) {
	suite.Run(t, new(DynamoDBConfigSuite))
}
