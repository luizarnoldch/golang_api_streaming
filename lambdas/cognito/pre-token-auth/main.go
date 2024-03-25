package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {
	log.Printf("PreTokenGen")
	// Triggerts works before generate the token on exchange code

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return event, err
	}

	log.Printf("Event Before Override: %s\n", string(eventJSON))


	// claimsToAdd := map[string]string{
	// 	"company_id": companyID,
	// }

	event.Response.ClaimsOverrideDetails.ClaimsToSuppress = []string{"family_name"}
	// event.Response.ClaimsOverrideDetails.ClaimsToAddOrOverride = claimsToAdd
	// event.Response.ClaimsOverrideDetails.GroupOverrideDetails = event.Request.GroupConfiguration

	eventJSON, err = json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return event, err
	}

	log.Printf("Event After Override: %s\n", string(eventJSON))

	return event, nil
}

func main() {
	lambda.Start(handler)
}