package dynamodb_test

import (
	"main/src/streams/domain/model"
	"main/utils/dynamodb"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/suite"
)

type DynamoDBSuite struct {
	suite.Suite
}

func TestDynamoDBSuite(t *testing.T) {
	suite.Run(t, new(DynamoDBSuite))
}

func (suite *DynamoDBSuite) TestUnmarshalStream() {
	mockItem := map[string]types.AttributeValue{
		"ID":         &types.AttributeValueMemberS{Value: "123"},
		"name":       &types.AttributeValueMemberS{Value: "Test Stream"},
		"cost":       &types.AttributeValueMemberN{Value: "10.99"},
		"start_date": &types.AttributeValueMemberS{Value: "2024-01-01"},
		"end_date":   &types.AttributeValueMemberS{Value: "2024-12-31"},
	}

	stream, err := dynamodb.UnmarshalStream(mockItem)
	suite.NoError(err)
	suite.NotNil(stream)
	suite.Equal("123", stream.ID)
	suite.Equal("Test Stream", stream.Name)
	suite.Equal(10.99, stream.Cost)
	suite.Equal("2024-01-01", stream.StartDate)
	suite.Equal("2024-12-31", stream.EndDate)
}

func (suite *DynamoDBSuite) TestUnmarshalStreamError() {
    invalidItem := map[string]types.AttributeValue{
        "InvalidField": &types.AttributeValueMemberS{Value: "Invalid Data"},
        "cost":         &types.AttributeValueMemberS{Value: "Invalid Cost Format"},
    }

    result, err := dynamodb.UnmarshalStream(invalidItem)
    suite.Nil(result, "Result should be nil on unmarshal error")
    suite.Error(err, "Should return an error for invalid unmarshalling")
}

func (suite *DynamoDBSuite) TestMarshalMapStream() {
	mockStream := &model.Stream{
		ID:        "123",
		Name:      "Test Stream",
		Cost:      10.99,
		StartDate: "2024-01-01",
		EndDate:   "2024-12-31",
	}

	item, err := dynamodb.MarshalMapStream(mockStream)
	suite.NoError(err)
	suite.NotNil(item)
	suite.IsType(&types.AttributeValueMemberS{}, item["ID"])
	suite.IsType(&types.AttributeValueMemberS{}, item["name"])
	suite.IsType(&types.AttributeValueMemberN{}, item["cost"])
	suite.IsType(&types.AttributeValueMemberS{}, item["start_date"])
	suite.IsType(&types.AttributeValueMemberS{}, item["end_date"])
}