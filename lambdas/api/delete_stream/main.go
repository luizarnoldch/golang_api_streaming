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
	id_stream, errParseRequest := apigateway.ParseAPIGatewayRequestParameters(request, "stream_id")
	if errParseRequest != nil {
		return apigateway.APIGatewayError(errParseRequest.Code, errParseRequest.ToError()), errParseRequest.ToError()
	}

	stream_handler := stream_micro.StreamUseCases{
		Ctx:       ctx,
		TableName: STREAM_TABLE,
	}

	errGetAllStream := stream_handler.DeleteStream(id_stream)

	if errGetAllStream != nil {
		return apigateway.APIGatewayError(errGetAllStream.Code, errGetAllStream.ToError()), errGetAllStream.ToError()
	}

	message := "Deleted stream with id: " + id_stream

	return apigateway.APIGatewayDataResponse(http.StatusOK, message), nil
}

func main() {
	lambda.Start(handler)
}
