package model_test

import (
	"main/src/users/domain/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserModelSuite struct {
	suite.Suite
}

func (suite *UserModelSuite) SetupSuite()    {}
func (suite *UserModelSuite) TearDownSuite() {}

func (suite *UserModelSuite) TestUserValidate(t *testing.T) {
	now := time.Now().Format(time.RFC3339)

	tests := []struct {
		name     string
		user     model.User
		expected bool
	}{
		{"Valid User", model.User{ID: uuid.NewString(), Name: "John Doe", Email: "john@example.com", CreatedAt: now}, true},
		{"Invalid ID", model.User{ID: "invalid-uuid", Name: "Jane Doe", Email: "jane@example.com", CreatedAt: now}, false},
		{"Invalid Name", model.User{ID: uuid.NewString(), Name: "J", Email: "jane@example.com", CreatedAt: now}, false},
		{"Invalid Email", model.User{ID: uuid.NewString(), Name: "Jane Doe", Email: "not-an-email", CreatedAt: now}, false},
		{"Invalid Date", model.User{ID: uuid.NewString(), Name: "Jane Doe", Email: "jane@example.com", CreatedAt: "invalid-date"}, false},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := tt.user.Validate()
			if tt.expected {
				suite.Nil(err)
			} else {
				suite.NotNil(err)
			}
		})
	}
}

func TestStreamModelSuite(t *testing.T) {
	suite.Run(t, new(UserModelSuite))
}
