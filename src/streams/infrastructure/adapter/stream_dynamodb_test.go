package adapter_test

import (
	"context"
	"main/src/streams/domain/model"
	"main/src/streams/domain/repository"
	"main/src/streams/infrastructure/adapter"
	"main/src/streams/infrastructure/configuration"

	"testing"

	dynamodbUtils "main/utils/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type StreamDynamoDBSuite struct {
	suite.Suite
	tableName                   string
	init_streams 				[]model.Stream
	dynamoClient                *dynamodb.Client
	dynamoDBLocalInfrastructure repository.StreamRepository
}

func (suite *StreamDynamoDBSuite) SetupSuite() {
	ctx := context.TODO()
	client, err := dynamodbUtils.GetLocalDynamoDBClient(ctx)
	suite.NoError(err)
	suite.dynamoClient = client

	table_name := configuration.GetDynamoDBStreamTable()
	suite.Equal("Test_Stream_Table", table_name)
	suite.tableName = table_name

	stream_infrastructure := adapter.NewStreamDynamoDBRepository(ctx, client,table_name)
	suite.dynamoDBLocalInfrastructure = stream_infrastructure

	configuration.CreateLocalDynamoDBStreamTable(ctx, suite.dynamoClient, suite.tableName)

	streams := []model.Stream{
        {
            ID:        uuid.NewString(),
            Name:      "test_name_1",
            Cost:      11.00,
            StartDate: "2022-01-01T15:04:05Z",
            EndDate:   "2023-12-01T15:04:05Z",
        },
        {
            ID:        uuid.NewString(),
            Name:      "test_name_2",
            Cost:      12.00,
            StartDate: "2022-01-02T15:04:05Z",
            EndDate:   "2023-12-02T15:04:05Z",
        },
        {
            ID:        uuid.NewString(),
            Name:      "test_name_3",
            Cost:      13.00,
            StartDate: "2022-01-03T15:04:05Z",
            EndDate:   "2023-12-03T15:04:05Z",
        },
    }
	suite.init_streams = streams
	exists, err := configuration.DescribeStreamTable(ctx, suite.dynamoClient, suite.tableName)
	suite.NoError(err)
	suite.True(exists)
}

func (suite *StreamDynamoDBSuite) TearDownSuite() {
	for _, stream := range suite.init_streams {
		err := suite.dynamoDBLocalInfrastructure.DeleteStream(stream.ID)
		suite.Nil(err)
	}
}

func (suite *StreamDynamoDBSuite) TestCreateStreamSuccessful() {
	for _, stream := range suite.init_streams {
		_, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
		suite.Nil(err)
	}
}

func (suite *StreamDynamoDBSuite) TestGetAllStreamSuccessful() {
    createdStreams, err := suite.dynamoDBLocalInfrastructure.GetAllStream()
    suite.Nil(err)
    suite.NotEmpty(createdStreams)
	suite.Equal(len(createdStreams), len(suite.init_streams))
}

func (suite *StreamDynamoDBSuite) TestGetStreamByIdSuccessful() {
	stream := suite.init_streams[0]
	retrievedStream, err := suite.dynamoDBLocalInfrastructure.GetStreamById(stream.ID)
	suite.Nil(err)
	suite.Equal(stream.ID, retrievedStream.ID)
	suite.Equal(stream.Name, retrievedStream.Name)
	suite.Equal(stream.Cost, retrievedStream.Cost)
	suite.Equal(stream.StartDate, retrievedStream.StartDate)
	suite.Equal(stream.EndDate, retrievedStream.EndDate)
}

func (suite *StreamDynamoDBSuite) TestUpdateStreamNameSuccessful() {
    old_stream := suite.init_streams[0]
	old_stream.Name = "new_name_test"

	suite.NotEqual(suite.init_streams[0], old_stream.Name)

	updatedStream, err := suite.dynamoDBLocalInfrastructure.UpdateStreamById(old_stream.ID, &old_stream)
	suite.Nil(err)
	suite.NotNil(updatedStream)
	suite.NotEmpty(updatedStream)

	suite.Equal(old_stream.Name, updatedStream.Name)
}

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamDynamoDBSuite))
}