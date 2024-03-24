package model_test

import (
	"main/src/streams/domain/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type StreamModelSuite struct {
	suite.Suite
}

func (suite *StreamModelSuite) SetupSuite() {}
func (suite *StreamModelSuite) TearDownSuite() {}

func (suite *StreamModelSuite) TestStreamValidate() {
	now := time.Now().Format(time.RFC3339)
	future := time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	var tests = []struct {
		name     string
		stream   model.Stream
		expected bool // true if no error is expected, false otherwise
	}{
		{"Valid Stream", model.Stream{ID: uuid.NewString(), Name: "Valid Name", Cost: 100.00, StartDate: now, EndDate: future}, true},
		{"Invalid UUID", model.Stream{ID: "invalid-uuid", Name: "Valid Name", Cost: 100.00, StartDate: now, EndDate: future}, false},
		{"Invalid Name", model.Stream{ID: uuid.NewString(), Name: "No", Cost: 100.00, StartDate: now, EndDate: future}, false},
		{"Negative Cost", model.Stream{ID: uuid.NewString(), Name: "Valid Name", Cost: -100.00, StartDate: now, EndDate: future}, false},
		{"Invalid Cost Precision", model.Stream{ID: uuid.NewString(), Name: "Valid Name", Cost: 100.001, StartDate: now, EndDate: future}, false},
		{"Invalid StartDate", model.Stream{ID: uuid.NewString(), Name: "Valid Name", Cost: 100.00, StartDate: "invalid-date", EndDate: future}, false},
		{"Invalid EndDate", model.Stream{ID: uuid.NewString(), Name: "Valid Name", Cost: 100.00, StartDate: now, EndDate: "invalid-date"}, false},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := tt.stream.Validate()
			if tt.expected {
				suite.Nil(err)
			} else {
				suite.NotNil(err)
			}
		})
	}
}

func TestStreamModelSuite(t *testing.T) {
	suite.Run(t, new(StreamModelSuite))
}
