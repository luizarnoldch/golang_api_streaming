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
		apigateway.APIGatewayError(http.StatusNoContent, err)
	}

	return apigateway.APIGatewayDataResponse(http.StatusOK, streams)
}

func main(){
	lambda.Start(handler)
}