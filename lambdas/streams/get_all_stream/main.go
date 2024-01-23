package main

import (
	"context"
	stream_handler "main/src/streams/application/use_cases"
	"main/utils/apigateway"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	streams, errGetAllStream := stream_handler.GetAllStream(ctx)

	if errGetAllStream != nil {
		return apigateway.APIGatewayError(errGetAllStream.Code, errGetAllStream.ToError()), errGetAllStream.ToError()
	}

	return apigateway.APIGatewayDataResponse(http.StatusOK, streams), nil
}

func main() {
	lambda.Start(handler)
}
