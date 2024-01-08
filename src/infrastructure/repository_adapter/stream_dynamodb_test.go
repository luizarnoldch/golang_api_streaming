package repositoryadapter_test

import (
	"context"
	"main/src/domain/model"
	"main/src/infrastructure/configuration"
	repositoryadapter "main/src/infrastructure/repository_adapter"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StreamSuite struct {
	suite.Suite
	dynamoLocalClient *repositoryadapter.StreamDynamoDBRepository
}

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamSuite))
}

func (suite *StreamSuite) BeforeTest() {
	ctx := context.TODO()
	client, err := configuration.GetLocalDynamoDBClient(ctx)
	suite.NoError(err)
	table_name := configuration.GetDynamoDBStreamTable()
	suite.Equal("Test_Stream_Table", table_name)

	configuration.CreateLocalDynamoDBStreamTable(client, ctx, table_name)

	stream_infrastructure := repositoryadapter.NewStreamDynamoDBRepository(client, ctx, table_name)
	suite.dynamoLocalClient = stream_infrastructure
}

func (suite *StreamSuite) AfterTest() {
	ctx := context.TODO()
	client, err := configuration.GetLocalDynamoDBClient(ctx)
	suite.NoError(err)
	table_name := configuration.GetDynamoDBStreamTable()
	suite.Equal("Test_Stream_Table", table_name)

	configuration.DeleteLocalDynamoDBStreamTable(client, ctx, table_name)
}

func (suite *StreamSuite) TestStreamValidateSuccessful() {
	stream := model.Stream{}
	response := stream.Validate()
	suite.NoError(response)
}
