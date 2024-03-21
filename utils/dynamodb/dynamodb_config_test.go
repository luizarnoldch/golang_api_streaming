package dynamodb_test

import (
	"context"
	"testing"

	dynamodb_util "main/utils/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/suite"
)

type DynamoDBUtilSuite struct {
	suite.Suite
	client *dynamodb.Client
	ctx    context.Context
}

func (suite *DynamoDBUtilSuite) SetupSuite() {
	suite.ctx = context.TODO()
	var err error
	suite.client, err = dynamodb_util.GetLocalDynamoDBClient(suite.ctx)
	suite.NoError(err)
}

func (suite *DynamoDBUtilSuite) TestGetDynamoDBAWSClient() {
	client, err := dynamodb_util.GetDynamoDBAWSClient(suite.ctx)
	suite.NoError(err)
	suite.IsType(&dynamodb.Client{}, client)
}

func (suite *DynamoDBUtilSuite) TestGetLocalDynamoDBClient() {
	client, err := dynamodb_util.GetLocalDynamoDBClient(suite.ctx)
	suite.NoError(err)
	suite.IsType(&dynamodb.Client{}, client)
}

func (suite *DynamoDBUtilSuite) TestGetLocalEndpoint() {
	endpoint, err := dynamodb_util.GetLocalEndpoint("dynamodb", "us-west-1")
	suite.NoError(err)
	suite.Equal("http://localhost:8000", endpoint.URL, "Endpoint URL should be 'http://localhost:8000'")
}

func TestDynamoDBUtilSuite(t *testing.T) {
	suite.Run(t, new(DynamoDBUtilSuite))
}
