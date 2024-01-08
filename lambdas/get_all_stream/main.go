package main

import (
	"context"
	"main/src/infrastructure/api/stream"
	"main/utils/apigateway"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	streams, err := stream.GetAllStream(ctx)

	if err != nil {
		return apigateway.APIGatewayError(http.StatusNoContent, err), err
	}

	return apigateway.APIGatewayDataResponse(http.StatusOK, streams), nil
}

func main() {
	lambda.Start(handler)
}
