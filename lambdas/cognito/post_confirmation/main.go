package main

import (
	"context"
	"encoding/json"
	"log"
	"main/src/users/domain/model"
	"os"

	user_micro "main/src/users/application/use_cases"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	USER_TABLE = os.Getenv("USER_TABLE")
)

func handler(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	log.Printf("PostConfirmation")

	// Doesn't Work for external Providers as Google and Microsoft
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return event, err
	}

	log.Printf("Event: %s\n", string(eventJSON))

	email := event.Request.UserAttributes["email"]
	name := event.Request.UserAttributes["name"]
	user_id := event.Request.UserAttributes["sub"]

	ctx := context.TODO()
	new_user := model.User{
		ID:    user_id,
		Name:  name,
		Email: email,
	}

	user_handler := user_micro.UserUseCases{
		Ctx:       ctx,
		TableName: USER_TABLE,
	}

	_, userErr := user_handler.CreateUser(&new_user)
	if userErr != nil {
		log.Println(userErr.ToString())
		return event, userErr.ToError()
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
