package adapter

import (
	"context"
	"main/src/streams/domain/model"
	"main/src/streams/infrastructure/configuration"
	"main/src/streams/domain/repository"

	"testing"

	dynamodbUtils "main/utils/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type StreamDynamoDBSuite struct {
	suite.Suite
	tableName                   string
	dynamoClient                *dynamodb.Client
	dynamoDBLocalInfrastructure repository.StreamRepository
}

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamDynamoDBSuite))
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
}

func (suite *StreamDynamoDBSuite) SetupTest() {
	ctx := context.TODO()
	configuration.CreateLocalDynamoDBStreamTable(suite.dynamoClient, ctx, suite.tableName)
}

func (suite *StreamDynamoDBSuite) TearDownTest() {
	ctx := context.TODO()
	configuration.DeleteLocalDynamoDBStreamTable(suite.dynamoClient, ctx, suite.tableName)
}

func (suite *StreamDynamoDBSuite) TestCreateStreamAndGetStreamByIdSuccessful() {
	stream_id := uuid.NewString()
	stream := model.Stream{
        ID:        stream_id,
        Name:      "test_name",
        Cost:      10.99,
        StartDate: "2022-01-01T15:04:05Z",
        EndDate:   "2023-12-01T15:04:05Z",
    }

	response := stream.Validate()
	if response != nil {
		suite.Fail("Validation failed", response.ToString())
	}

	_, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
	if err != nil {
		suite.Fail("Failed to create stream", err.ToString())
	}

	createdStream, err := suite.dynamoDBLocalInfrastructure.GetStreamById(stream_id)
	if err != nil {
		suite.Fail("Failed to retrieve stream by ID", err.ToString())
	}
	suite.Equal(stream.ID, createdStream.ID)
	suite.Equal(stream.Name, createdStream.Name)
	suite.Equal(stream.Cost, createdStream.Cost)
	suite.Equal(stream.StartDate, createdStream.StartDate)
	suite.Equal(stream.EndDate, createdStream.EndDate)
}

func (suite *StreamDynamoDBSuite) TestGetAllStreamSuccessful() {
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

	originalStreamsMap := make(map[string]model.Stream)

    for _, stream := range streams {
        createdStream, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
        suite.Nil(err)
        suite.NotNil(createdStream)

        originalStreamsMap[createdStream.ID] = *createdStream
    }

    createdStreams, err := suite.dynamoDBLocalInfrastructure.GetAllStream()
    suite.Nil(err)
    suite.NotEmpty(createdStreams)

    suite.Len(createdStreams, len(streams), "Number of created streams should match")


	for _, createdStream := range createdStreams {
        originalStream, ok := originalStreamsMap[createdStream.ID]
        suite.True(ok, "Stream ID should exist in the original streams")

        suite.Equal(originalStream.Name, createdStream.Name, "Stream name should match")
        suite.Equal(originalStream.Cost, createdStream.Cost, "Stream cost should match")
        suite.Equal(originalStream.StartDate, createdStream.StartDate, "Stream start date should match")
		suite.Equal(originalStream.EndDate, createdStream.EndDate, "Stream end date should match")
	}
}

func (suite *StreamDynamoDBSuite) TestGetStreamByIdWithBadId() {
	stream_id := uuid.NewString()
	stream_bad_id := uuid.NewString()

	stream := model.Stream{
		ID:        stream_id,
		Name:      "test_name",
		Cost:      15.00,
		StartDate: "2022-01-01T15:04:05Z",
		EndDate:   "2023-01-01T15:04:05Z",
	}
	response := stream.Validate()
	if response != nil {
		suite.Fail("Validation failed", response.ToString())
	}

	_, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
	if err != nil {
		suite.Fail("Failed to create stream", err.ToString())
	}

	_, err = suite.dynamoDBLocalInfrastructure.GetStreamById(stream_bad_id)
	suite.NotNil(err, "Expected an error for bad ID")
}
