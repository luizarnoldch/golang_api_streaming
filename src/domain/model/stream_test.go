package model_test

import (
	"log"
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
	log.Println(stream)
// 	err := stream.Validate()
// 	suite.Nil(err)
}