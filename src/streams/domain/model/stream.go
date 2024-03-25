package model

import (
	appError "main/utils/error"
	"main/utils/validation"
)

type Stream struct {
    ID        string  `json:"ID,omitempty" dynamodbav:"ID,omitempty"`
    Name      string  `json:"name" dynamodbav:"name"`
    Cost      float64 `json:"cost" dynamodbav:"cost"`
    StartDate string  `json:"start_date" dynamodbav:"start_date"`
    EndDate   string  `json:"end_date" dynamodbav:"end_date"`
}

func (s Stream) Validate() *appError.Error {
	if err := validation.ValidateUUID(s.ID); err != nil {
		return err
	}
	if err := validation.ValidateStringField(s.Name, 2); err != nil {
		return err
	}
	if err := validation.ValidateCost(s.Cost); err != nil {
		return err
	}
	if err := validation.ValidateDate(s.StartDate); err != nil {
		return err
	}
	if err := validation.ValidateDate(s.EndDate); err != nil {
		return err
	}
	return nil
}