package repositoryadapter_test

import (
	"context"
	"fmt"
	"main/src/domain/model"
	"main/src/infrastructure/configuration"
	repositoryadapter "main/src/infrastructure/repository_adapter"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type StreamSuite struct {
	suite.Suite
	tableName string
	dynamoClient *dynamodb.Client
	dynamoDBLocalInfrastructure *repositoryadapter.StreamDynamoDBRepository
}

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamSuite))
}

func (suite *StreamSuite) SetupSuite() {
	ctx := context.TODO()
	client, err := configuration.GetLocalDynamoDBClient(ctx)
	suite.NoError(err)
	suite.dynamoClient = client

	table_name := configuration.GetDynamoDBStreamTable()
	suite.Equal("Test_Stream_Table", table_name)
	suite.tableName = table_name

	stream_infrastructure := repositoryadapter.NewStreamDynamoDBRepository(client, ctx, table_name)
	suite.dynamoDBLocalInfrastructure = stream_infrastructure
}

func (suite *StreamSuite) SetupTest() {
	ctx := context.TODO()
    configuration.CreateLocalDynamoDBStreamTable(suite.dynamoClient, ctx, suite.tableName)
}

func (suite *StreamSuite) TearDownTest() {
	ctx := context.TODO()
	configuration.DeleteLocalDynamoDBStreamTable(suite.dynamoClient, ctx, suite.tableName)
}

func (suite *StreamSuite) TestCreateStreamAndGetStreamByIdSuccessful() {
    stream_id := uuid.NewString()
	
	stream := model.Stream{
        ID: stream_id,
        Name: "test_name",
        Cost: 15.00,
        StartDate: "01-01-24",
        EndDate: "12-12-24",
    }
    response := stream.Validate()
    suite.NoError(response)

    _, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
    suite.NoError(err)

    createdStream, err := suite.dynamoDBLocalInfrastructure.GetStreamById(stream_id)
    suite.NoError(err, "Failed to retrieve stream by ID")

    suite.Equal(stream.ID, createdStream.ID)
    suite.Equal(stream.Name, createdStream.Name)
    suite.Equal(stream.Cost, createdStream.Cost)
    suite.Equal(stream.StartDate, createdStream.StartDate)
    suite.Equal(stream.EndDate, createdStream.EndDate)
}

func (suite *StreamSuite) TestGetAllStreamSuccessful() {
    streams := []model.Stream{
        {
            ID:        "01" + uuid.NewString(),
            Name:      "test_name_1",
            Cost:      11.00,
            StartDate: "01-01-24",
            EndDate:   "01-12-24",
        },
        {
            ID:        "02" + uuid.NewString(),
            Name:      "test_name_2",
            Cost:      12.00,
            StartDate: "02-01-24",
            EndDate:   "02-12-24",
        },
        {
            ID:        "03" + uuid.NewString(),
            Name:      "test_name_3",
            Cost:      13.00,
            StartDate: "03-01-24",
            EndDate:   "03-12-24",
        },
    }

    for _, stream := range streams {
        _, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
        suite.NoError(err)
    }

    createdStreams, err := suite.dynamoDBLocalInfrastructure.GetAllStream()
    suite.NoError(err)
    suite.Len(createdStreams, len(streams), "Number of created streams should match")

    originalStreamsMap := make(map[string]model.Stream)
    for _, stream := range streams {
        originalStreamsMap[stream.ID] = stream
    }

    for _, createdStream := range createdStreams {
        originalStream, ok := originalStreamsMap[createdStream.ID]
        suite.True(ok, "Stream ID should exist in the original streams")

        suite.Equal(originalStream.Name, createdStream.Name, "Stream name should match")
        suite.Equal(originalStream.Cost, createdStream.Cost, "Stream cost should match")
        suite.Equal(originalStream.StartDate, createdStream.StartDate, "Stream start date should match")
        suite.Equal(originalStream.EndDate, createdStream.EndDate, "Stream end date should match")
    }
}

func (suite *StreamSuite) TestGetStreamByIdWithBadId() {
    stream_id := uuid.NewString()
    stream_bad_id := uuid.NewString()
    
    stream := model.Stream{
        ID: stream_id,
        Name: "test_name",
        Cost: 15.00,
        StartDate: "01-01-24",
        EndDate: "12-12-24",
    }
    response := stream.Validate()
    suite.NoError(response)

    _, err := suite.dynamoDBLocalInfrastructure.CreateStream(&stream)
    suite.NoError(err)

    retrievedStream, err := suite.dynamoDBLocalInfrastructure.GetStreamById(stream_bad_id)
    suite.Error(err, fmt.Errorf("GetStreamById: No stream found with ID: %s", stream_bad_id))
	suite.Nil(retrievedStream)
}

