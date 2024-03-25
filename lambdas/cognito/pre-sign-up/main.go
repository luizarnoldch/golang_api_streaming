package main

import (
	"encoding/json"
	"log"
	"main/utils/validation"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {

	event.Response.AutoConfirmUser = false
	event.Response.AutoVerifyEmail = false
	event.Response.AutoVerifyPhone = false

	// [test] Doesn't Work for external Providers as Google and Microsoft
	log.Println("Pre signup Cognito confirmation")

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return event, err
	}

	log.Printf("Event: %s\n", string(eventJSON))

	email := event.Request.UserAttributes["email"]

	if err := validation.ValidateEmail(email); err != nil {
		log.Println(err.ToString())
		return event, err.ToError()
	}

	// [todo] In case need to autoconfirm something
	// event.Response.AutoConfirmUser = true
	// event.Response.AutoVerifyEmail = true
	// event.Response.AutoVerifyPhone = true
	return event, nil
}

func main() {
	lambda.Start(handler)
}
