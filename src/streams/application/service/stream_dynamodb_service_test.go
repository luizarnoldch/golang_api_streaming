package service_test

import (
	"testing"

	streamMock "main/mocks"
	"main/src/streams/application/service"
	"main/src/streams/domain/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type StreamServiceDynamoDBSuite struct {
	suite.Suite
	streamDynamoDBRepository *streamMock.StreamRepository
	streamServiceApplication service.StreamService
}

const (
	MethodUpdateStreamById = "UpdateStreamById"
	MethodGetStreamById    = "GetStreamById"
	MethodGetAllStream     = "GetAllStream"
	MethodDeleteStream     = "DeleteStream"
	MethodCreateStream     = "CreateStream"
)

func (suite *StreamServiceDynamoDBSuite) SetupTest() {
	suite.streamDynamoDBRepository = new(streamMock.StreamRepository)
	suite.streamServiceApplication = service.NewStreamDynamoDBService(suite.streamDynamoDBRepository)
}

func (suite *StreamServiceDynamoDBSuite) TestCreateStream() {
	new_stream := &model.Stream{
		Name:      "Test Stream",
		Cost:      10.99,
		StartDate: "2023-01-01T00:00:00Z",
		EndDate:   "2023-12-31T23:59:59Z",
	}

	suite.streamDynamoDBRepository.On(MethodCreateStream, new_stream).Return(new_stream, nil)
	created_stream, err := suite.streamServiceApplication.CreateStream(new_stream)
	suite.Nil(err, "Creating a stream should not return an error")
	suite.NotNil(created_stream, "Created stream should not be nil")
	suite.NotEqual("", created_stream.ID, "Created stream should have a non-empty ID")
	suite.Equal(new_stream.Name, created_stream.Name, "Created stream name should match the input")
	suite.streamDynamoDBRepository.AssertExpectations(suite.T())
}

func (suite *StreamServiceDynamoDBSuite) TestDeleteStream() {
	streamID := uuid.NewString()
	suite.streamDynamoDBRepository.On(MethodDeleteStream, streamID).Return(nil)
	err := suite.streamServiceApplication.DeleteStream(streamID)
	suite.Nil(err, "Deleting a stream should not return an error")
	suite.streamDynamoDBRepository.AssertExpectations(suite.T())
}

func (suite *StreamServiceDynamoDBSuite) TestGetStreamById() {
	streamID := uuid.NewString()
	expectedStream := &model.Stream{
		ID:   streamID,
		Name: "Test Stream",
	}

	suite.streamDynamoDBRepository.On(MethodGetStreamById, streamID).Return(expectedStream, nil)
	stream, err := suite.streamServiceApplication.GetStreamById(streamID)

	suite.Nil(err, "Getting a stream by ID should not return an error")
	suite.Equal(expectedStream.ID, stream.ID, "The returned stream ID should match the expected ID")

	suite.streamDynamoDBRepository.AssertExpectations(suite.T())
}

func (suite *StreamServiceDynamoDBSuite) TestUpdateStreamById() {
	streamID := uuid.NewString()
	updatedStream := &model.Stream{
		ID:        streamID,
		Name:      "Updated Stream",
		Cost:      15.99,
		StartDate: "2023-06-01T00:00:00Z",
		EndDate:   "2023-12-31T23:59:59Z",
	}
	suite.streamDynamoDBRepository.On(MethodUpdateStreamById, streamID, updatedStream).Return(updatedStream, nil)
	resultStream, err := suite.streamServiceApplication.UpdateStreamById(streamID, updatedStream)
	suite.Nil(err, "Updating a stream should not return an error")
	suite.Equal(updatedStream.Name, resultStream.Name, "Updated stream name should match")
	suite.Equal(updatedStream.Cost, resultStream.Cost, "Updated stream cost should match")
	suite.streamDynamoDBRepository.AssertExpectations(suite.T())
}

func (suite *StreamServiceDynamoDBSuite) TestGetAllStream() {
	expectedStreams := []model.Stream{
		{ID: uuid.NewString(), Name: "Stream 1"},
		{ID: uuid.NewString(), Name: "Stream 2"},
	}
	suite.streamDynamoDBRepository.On(MethodGetAllStream).Return(expectedStreams, nil)
	streams, err := suite.streamServiceApplication.GetAllStream()
	suite.Nil(err, "Getting all streams should not return an error")
	suite.Len(streams, len(expectedStreams), "The returned slice should have the same number of streams as expected")
	suite.streamDynamoDBRepository.AssertExpectations(suite.T())
}

func TestStreamServiceDynamoDBSuite(t *testing.T) {
	suite.Run(t, new(StreamServiceDynamoDBSuite))
}
