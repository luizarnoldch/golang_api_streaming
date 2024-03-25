package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CognitoEventUserPoolsPostAuthentication) (events.CognitoEventUserPoolsPostAuthentication, error) {
	log.Printf("PostAuthentication")

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return event, err
	}

	log.Printf("Event: %s\n", string(eventJSON))

	/**
	* TODO: Override Claims
	* TODO: Add to cognito group
	*/

	return event, nil
}

func main() {
	lambda.Start(handler)
}
