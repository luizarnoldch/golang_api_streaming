package model_test

import (
	"main/src/users/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	now := time.Now().Format(time.RFC3339)

	tests := []struct {
		name     string
		user     model.User
		expected bool
	}{
		{"Valid User", model.User{ID: "a0b1c2d3-e4f5-6789-0123-456789abcdef", Name: "John Doe", Email: "john@example.com", CreatedAt: now}, true},
		{"Invalid ID", model.User{ID: "invalid-uuid", Name: "Jane Doe", Email: "jane@example.com", CreatedAt: now}, false},
		{"Invalid Name", model.User{ID: "a0b1c2d3-e4f5-6789-0123-456789abcdef", Name: "J", Email: "jane@example.com", CreatedAt: now}, false},
		{"Invalid Email", model.User{ID: "a0b1c2d3-e4f5-6789-0123-456789abcdef", Name: "Jane Doe", Email: "not-an-email", CreatedAt: now}, false},
		{"Invalid Date", model.User{ID: "a0b1c2d3-e4f5-6789-0123-456789abcdef", Name: "Jane Doe", Email: "jane@example.com", CreatedAt: "invalid-date"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			assert.Equal(t, tt.expected, err == nil)
		})
	}
}
