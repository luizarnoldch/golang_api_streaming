package validation

import (
	appError "main/utils/error"
	"math"

	"regexp"
	"strings"
	"time"
)

func ValidateUUID(id string) *appError.Error {
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	if !uuidRegex.MatchString(id) {
		return appError.NewValidationError("Invalid ID: The ID must be a valid UUID.")
	}
	return nil
}

func ValidateStringField(field string, minLength int) *appError.Error {
	if len(strings.TrimSpace(field)) < minLength {
		return appError.NewValidationError("Field length is less than the required minimum")
	}
	return nil
}

func ValidateEmail(email string) *appError.Error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return appError.NewValidationError("Invalid email format")
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
