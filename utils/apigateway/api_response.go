package apigateway

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// HeadersJSON - standard headers for JSON responses
var HeadersJSON = map[string]string{
	// "Access-Control-Allow-Origin":  "*",
	// "Access-Control-Allow-Methods": "DELETE,GET,HEAD,POST,PUT",
	// "Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
	"Content-Type":                 "application/json",
}

// APIGatewayResponse - centralize response creation logic
func APIGatewayResponse(statusCode int, body interface{}, headers map[string]string) events.APIGatewayProxyResponse {
	responseData, err := json.Marshal(body)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		// Handle marshaling error or return a generic error response
		return events.APIGatewayProxyResponse{
			Body:       string(`{"error": "Error marshaling response"}`),
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
		}
	}

	log.Printf("Response: %s", responseData)
	return events.APIGatewayProxyResponse{
		Body:       string(responseData),
		StatusCode: statusCode,
		Headers:    headers,
	}
}

// APIGatewayMessageResponse - create a response with a simple message
func APIGatewayMessageResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	return APIGatewayResponse(statusCode, map[string]string{"message": message}, HeadersJSON)
}

// APIGatewayDataResponse - create a response with more complex data
func APIGatewayDataResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	return APIGatewayResponse(statusCode, data, HeadersJSON)
}

// APIGatewayError - create an error response
func APIGatewayError(statusCode int, err error) events.APIGatewayProxyResponse {
	return APIGatewayResponse(statusCode, map[string]string{"error": err.Error()}, HeadersJSON)
}