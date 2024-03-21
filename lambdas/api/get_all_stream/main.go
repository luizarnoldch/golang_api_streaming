package main

import (
	"context"
	stream_micro "main/src/streams/application/use_cases"
	"main/utils/apigateway"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	STREAM_TABLE = os.Getenv("STREAM_TABLE")
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	stream_handler := stream_micro.StreamUseCases{
        Ctx: ctx,
        TableName: STREAM_TABLE,
    }

	streams, errGetAllStream := stream_handler.GetAllStream()

	if errGetAllStream != nil {
		return apigateway.APIGatewayError(errGetAllStream.Code, errGetAllStream.ToError()), errGetAllStream.ToError()
	}

	return apigateway.APIGatewayDataResponse(http.StatusOK, streams), nil
}

func main() {
	lambda.Start(handler)
}
