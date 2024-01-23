package model_test

import (
	"main/src/streams/domain/model"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StreamSuite struct {
	suite.Suite
}

func TestStreamSuite(t *testing.T) {
	suite.Run(t, new(StreamSuite))
}

func (suite *StreamSuite) TestStreamValidateSuccessful() {
	stream := model.Stream{
		ID:        "123e4567-e89b-12d3-a456-426614174000",
		Name:      "Valid Name",
		Cost:      15.99,
		StartDate: "2022-01-01T15:04:05Z",
		EndDate:   "2023-01-01T15:04:05Z",
	}
	err := stream.Validate()
	suite.Nil(err)
}

func (suite *StreamSuite) TestStreamValidateInvalidID() {
	stream := model.Stream{
		ID: "invalid-uuid",
	}
	err := stream.Validate()
	suite.NotNil(err)
	suite.Equal(err.Message, "Invalid UUID format for ID")
}

func (suite *StreamSuite) TestStreamValidateInvalidName() {
	stream := model.Stream{
		ID:   "123e4567-e89b-12d3-a456-426614174000",
		Name: "N",
	}
	err := stream.Validate()
	suite.NotNil(err)
	suite.Equal(err.Message, "Field length is less than required minimum")
}

func (suite *StreamSuite) TestStreamValidateInvalidCostNegative() {
	stream := model.Stream{
		ID:   "123e4567-e89b-12d3-a456-426614174000",
		Name: "Name",
		Cost: -10.00,
	}
	err := stream.Validate()
	suite.NotNil(err)
	suite.Equal(err.Message, "Cost cannot be negative")
}

func (suite *StreamSuite) TestStreamValidateInvalidCostMoreThan2Decimals() {
	stream := model.Stream{
		ID:   "123e4567-e89b-12d3-a456-426614174000",
		Name: "Name",
		Cost: 10.123,
	}
	err := stream.Validate()
	suite.NotNil(err)
	suite.Equal(err.Message, "Cost cannot have more than two decimal places")
}

func (suite *StreamSuite) TestStreamValidateInvalidStartDate() {
	// Test with invalid start date
	stream := model.Stream{
		ID:        "123e4567-e89b-12d3-a456-426614174000",
		Name: "Name",
		Cost: 10.00,
		StartDate: "invalid-date",
	}
	err := stream.Validate()
	suite.NotNil(err)
	suite.Equal(err.Message, "Invalid date format, must be RFC3339 or RFC3339 with microseconds")
}

func (suite *StreamSuite) TestStreamValidateInvalidEndDate() {
	// Test with invalid end date
	stream := model.Stream{
		ID:      "123e4567-e89b-12d3-a456-426614174000",
		Name: "Name",
		Cost: 10.00,
		StartDate: "2006-01-02T15:04:05.99Z",
		EndDate: "invalid-date",
	}
	err := stream.Validate()
	suite.NotNil(err)
	suite.Equal(err.Message, "Invalid date format, must be RFC3339 or RFC3339 with microseconds")
}
