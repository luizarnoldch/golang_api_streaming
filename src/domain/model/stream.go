package model

type Stream struct {
	ID        string `json:"id,omitempty" dynamodbav:"id,omitempty"`
	Name      string `json:"name" dynamodbav:"name"`
	Cost      string `json:"cost" dynamodbav:"cost"`
	StartDate string `json:"start_date" dynamodbav:"start_date"`
	EndDate   string `json:"end_date" dynamodbav:"end_date"`
}

func (s Stream) Validate() error {
	return nil
}