package model

import (
	appError "main/utils/error"
	"math"
	"regexp"
	"strings"
	"time"
)

type Stream struct {
    ID        string  `json:"ID,omitempty" dynamodbav:"ID,omitempty"`
    Name      string  `json:"name" dynamodbav:"name"`
    Cost      float64 `json:"cost" dynamodbav:"cost"`
    StartDate string  `json:"start_date" dynamodbav:"start_date"`
    EndDate   string  `json:"end_date" dynamodbav:"end_date"`
}

func (s Stream) Validate() *appError.Error {
	if err := ValidateUUID(s.ID); err != nil {
		return err
	}
	if err := ValidateStringField(s.Name, 2); err != nil {
		return err
	}
	if err := ValidateCost(s.Cost); err != nil {
		return err
	}
	if err := ValidateDate(s.StartDate); err != nil {
		return err
	}
	if err := ValidateDate(s.EndDate); err != nil {
		return err
	}
	return nil
}

func ValidateUUID(id string) *appError.Error {
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	if !uuidRegex.MatchString(id) {
		return appError.NewValidationError("Invalid ID: The ID must be a valid UUID.")
	}
	return nil
}

func ValidateStringField(field string, minLength int) *appError.Error {
	if len(strings.TrimSpace(field)) <= minLength {
		return appError.NewValidationError("Field length is less than required minimum")
	}
	return nil
}

func ValidateCost(cost float64) *appError.Error {
    if cost < 0 {
        return appError.NewValidationError("Cost cannot be negative")
    }
    if _, frac := math.Modf(cost * 100); frac != 0 {
        return appError.NewValidationError("Cost cannot have more than two decimal places")
    }
    return nil
}


func ValidateDate(dateStr string) *appError.Error {
	_, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		if _, err := time.Parse(time.RFC3339Nano, dateStr); err != nil {
			return appError.NewValidationError("Invalid date format, must be RFC3339 or RFC3339 with microseconds")
		}
	}
	return nil
}