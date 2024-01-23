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
	if err := validateUUID(s.ID); err != nil {
		return err
	}
	if err := validateStringField(s.Name, 2); err != nil {
		return err
	}
	if err := validateCost(s.Cost); err != nil {
		return err
	}
	if err := validateDate(s.StartDate); err != nil {
		return err
	}
	if err := validateDate(s.EndDate); err != nil {
		return err
	}
	return nil
}

func validateUUID(id string) *appError.Error {
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89ab][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	if !uuidRegex.MatchString(id) {
		return appError.NewValidationError("Invalid UUID format for ID")
	}
	return nil
}

func validateStringField(field string, minLength int) *appError.Error {
	if len(strings.TrimSpace(field)) <= minLength {
		return appError.NewValidationError("Field length is less than required minimum")
	}
	return nil
}

func validateCost(cost float64) *appError.Error {
    if cost < 0 {
        return appError.NewValidationError("Cost cannot be negative")
    }
    if _, frac := math.Modf(cost * 100); frac != 0 {
        return appError.NewValidationError("Cost cannot have more than two decimal places")
    }
    return nil
}


func validateDate(dateStr string) *appError.Error {
	_, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		if _, err := time.Parse(time.RFC3339Nano, dateStr); err != nil {
			return appError.NewValidationError("Invalid date format, must be RFC3339 or RFC3339 with microseconds")
		}
	}
	return nil
}