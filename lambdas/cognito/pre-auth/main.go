package main

import (
	"encoding/json"
	"log"
	"main/utils/validation"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CognitoEventUserPoolsPreAuthentication) (events.CognitoEventUserPoolsPreAuthentication, error) {
	log.Printf("PreAuthentication")

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return event, err
	}

	log.Printf("Event: %s\n", string(eventJSON))

	/**
	* TODO: Auditoy changes.
	* TODO: Login control.
	*/

	email := event.Request.UserAttributes["email"]

	if err := validation.ValidateEmail(email); err != nil {
		log.Println(err.ToString())
		return event, err.ToError()
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
