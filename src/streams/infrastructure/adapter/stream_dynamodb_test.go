package adapter

import (
	"context"
	"main/src/streams/domain/model"
	"main/src/streams/domain/repository"
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

	stream_infrastructure := NewStreamDynamoDBRepository(client, ctx, table_name)
	suite.dynamoDBLocalInfrastructure = stream_infrastructure

	configuration.CreateLocalDynamoDBStreamTable(suite.dynamoClient, ctx, suite.tableName)

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
	createdStream, err := suite.dynamoDBLocalInfrastructure.GetStreamById(stream.ID)
	suite.Nil(err)
	suite.Equal(createdStream.ID, stream.ID)
	suite.Equal(createdStream.Name, stream.Name)
	suite.Equal(createdStream.Cost, stream.Cost)
	suite.Equal(createdStream.StartDate, stream.StartDate)
	suite.Equal(createdStream.EndDate, stream.EndDate)
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



// func (suite *StreamDynamoDBSuite) TestGetStreamByIdWithBadId() {
// 	stream_id := uuid.NewString()
// 	stream_bad_id := uuid.NewString()

// 	stream := model.Stream{
// 		ID:        stream_id,
// 		Name:      "test_name",
// 		Cost:      15.00,
// 		StartDate: "2022-01-01T15:04:05Z",
// 		EndDate:   "2023-01-01T15:04:05Z",
// 	}
// 	response := stream.Validate()
// 	if response != nil {
// 		suite.Fail("Validation failed", response.ToString())
// 	}

// 	_, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
// 	if err != nil {
// 		suite.Fail("Failed to create stream", err.ToString())
// 	}

// 	_, err = suite.dynamoDBLocalInfrastructure.GetStreamById(stream_bad_id)
// 	suite.NotNil(err, "Expected an error for bad ID")
// }

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamDynamoDBSuite))
}