package model

import (
	appError "main/utils/error"
	"strings"
)

type Stream struct {
	ID        string  `json:"ID,omitempty" dynamodbav:"ID,omitempty"`
	Name      string  `json:"name" dynamodbav:"name"`
	Cost      float64 `json:"cost" dynamodbav:"cost"`
	StartDate string  `json:"start_date" dynamodbav:"start_date"`
	EndDate   string  `json:"end_date" dynamodbav:"end_date"`
}

func (s Stream) Validate() *appError.Error {
	if strings.TrimSpace(s.ID) == "" {
        return appError.NewValidationError("ID is required")
    }
    // if strings.TrimSpace(s.Name) == "" {
    //     return errors.New("Name is required")
    // }
    // if strings.TrimSpace(s.StartDate) == "" {
    //     return errors.New("StartDate is required")
    // }
    // if strings.TrimSpace(s.EndDate) == "" {
    //     return errors.New("EndDate is required")
    // }
	return nil
}