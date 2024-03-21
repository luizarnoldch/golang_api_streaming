package service_test

// import (
// 	"context"
// 	"main/src/streams/domain/model"
// 	"main/src/streams/application/service"
// 	"main/src/streams/infrastructure/adapter"
// 	"main/src/streams/infrastructure/configuration"
// 	dynamodbUtils "main/utils/dynamodb"
// 	"testing"

// 	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
// 	"github.com/stretchr/testify/suite"
// )

// type StreamDynamoDBServiceSuite struct {
// 	suite.Suite
// 	tableName                   string
// 	dynamoClient                *dynamodb.Client
// 	service service.StreamService
// }

// func TestStreamSuite(t *testing.T) {
// 	suite.Run(t, new(StreamDynamoDBServiceSuite))
// }

// func (suite *StreamDynamoDBServiceSuite) SetupSuite() {
// 	ctx := context.TODO()
// 	client, err := dynamodbUtils.GetLocalDynamoDBClient(ctx)
// 	suite.NoError(err)
// 	suite.dynamoClient = client

// 	table_name := configuration.GetDynamoDBStreamTable()
// 	suite.Equal("Test_Stream_Table", table_name)
// 	suite.tableName = table_name

// 	stream_infrastructure := adapter.NewStreamDynamoDBRepository(client, ctx, table_name)
// 	stream_application := service.NewStreamDynamoDBService(stream_infrastructure)

// 	suite.service = stream_application
// }

// func (suite *StreamDynamoDBServiceSuite) SetupTest() {
// 	ctx := context.TODO()
// 	configuration.CreateLocalDynamoDBStreamTable(suite.dynamoClient, ctx, suite.tableName)
// }

// func (suite *StreamDynamoDBServiceSuite) TearDownTest() {
// 	ctx := context.TODO()
// 	configuration.DeleteLocalDynamoDBStreamTable(suite.dynamoClient, ctx, suite.tableName)
// }

// func (suite *StreamDynamoDBServiceSuite) TestCreateStreamAndGetStreamByIdSuccessful() {
// 	stream := model.Stream{
//         Name:      "test_name",
//         Cost:      10.99,
//         StartDate: "2022-01-01T15:04:05Z",
//         EndDate:   "2023-12-01T15:04:05Z",
//     }

//     createdStream, err := suite.service.CreateStream(&stream)
//     suite.Nil(err)
//     suite.NotNil(createdStream)
//     suite.NotEmpty(createdStream.ID)

// 	retrievedStream, err := suite.service.GetStreamById(createdStream.ID)
// 	suite.Nil(err)

// 	suite.Equal(createdStream.ID, retrievedStream.ID)
// 	suite.Equal(createdStream.Name, retrievedStream.Name)
// 	suite.Equal(createdStream.Cost, retrievedStream.Cost)
// 	suite.Equal(createdStream.StartDate, retrievedStream.StartDate)
// 	suite.Equal(createdStream.EndDate, retrievedStream.EndDate)
// }

// func (suite *StreamDynamoDBServiceSuite) TestGetAllStreamSuccessful() {
//     streams := []model.Stream{
//         {
//             Name:      "test_name_1",
//             Cost:      11.00,
//             StartDate: "2022-01-01T15:04:05Z",
//             EndDate:   "2023-12-01T15:04:05Z",
//         },
//         {
//             Name:      "test_name_2",
//             Cost:      12.00,
//             StartDate: "2022-01-02T15:04:05Z",
//             EndDate:   "2023-12-02T15:04:05Z",
//         },
//         {
//             Name:      "test_name_3",
//             Cost:      13.00,
//             StartDate: "2022-01-03T15:04:05Z",
//             EndDate:   "2023-12-03T15:04:05Z",
//         },
//     }

//     originalStreamsMap := make(map[string]model.Stream)

//     for _, stream := range streams {
//         createdStream, err := suite.service.CreateStream(&stream)
//         suite.Nil(err)
//         suite.NotNil(createdStream)
//         suite.NotEmpty(createdStream.ID)

//         stream.ID = createdStream.ID
//         originalStreamsMap[createdStream.ID] = *createdStream
//     }

//     savedStreams, err := suite.service.GetAllStream()
//     suite.Nil(err)
//     suite.NotEmpty(savedStreams)

//     suite.Len(savedStreams, len(streams), "Number of created streams should match")

//     for _, stream := range savedStreams {
//         originalStream, ok := originalStreamsMap[stream.ID]
//         suite.True(ok, "Stream ID should exist in the original streams")
//         suite.NotNil(stream)
//         suite.NotEmpty(originalStream)

//         // suite.Equal(originalStream.Name, stream.Name, "Stream name should match")
//         // suite.Equal(originalStream.Cost, stream.Cost, "Stream cost should match")
//         // suite.Equal(originalStream.StartDate, stream.StartDate, "Stream start date should match")
// 		// suite.Equal(originalStream.EndDate, stream.EndDate, "Stream end date should match")
//     }
// }