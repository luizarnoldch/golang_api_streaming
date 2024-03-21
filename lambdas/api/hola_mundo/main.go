package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	STREAM_TABLE = os.Getenv("STREAM_TABLE")
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	message := "Hola mundo " + STREAM_TABLE

	return events.APIGatewayProxyResponse{
		Body: message,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
