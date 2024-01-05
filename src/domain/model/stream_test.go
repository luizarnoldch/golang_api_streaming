package model_test

import (
	"main/src/domain/model"
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
	stream := model.Stream{}
	response := stream.Validate()
	suite.NoError(response)
}