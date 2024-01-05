package apigateway

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

var (
	HEADERS_ALL = map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "DELETE,GET,HEAD,POST,PUT",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Content-Type":                 "application/json",
	}

	HEADERS_JSON = map[string]string{
		"Content-Type": "application/json",
	}
)


func APIGatewayMessageResponse(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	responseData, _ := json.Marshal(map[string]string{"message": message})
	return events.APIGatewayProxyResponse{
		Body:       string(responseData),
		StatusCode: statusCode,
		Headers:    HEADERS_JSON,
	}, nil
}

func APIGatewayDataResponse(statusCode int, data interface{}) (events.APIGatewayProxyResponse, error) {
	responseData, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		Body:       string(responseData),
		StatusCode: statusCode,
		Headers:    HEADERS_JSON,
	}, nil
}


func APIGatewayError(statusCode int, err error) (events.APIGatewayProxyResponse, error) {
    errorMessage, _ := json.Marshal(map[string]string{"error": err.Error()})
    return events.APIGatewayProxyResponse{
        Body:       string(errorMessage),
        StatusCode: statusCode,
		Headers:    HEADERS_JSON,
    }, nil
}