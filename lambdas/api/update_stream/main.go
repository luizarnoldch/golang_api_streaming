package main

import (
	"context"
	stream_micro "main/src/streams/application/use_cases"
	"main/src/streams/domain/model"
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
	var streamRequest model.Stream

	id_stream, errParseRequest := apigateway.ParseAPIGatewayRequestParameters(request, "stream_id")
	if errParseRequest != nil {
		return apigateway.APIGatewayError(errParseRequest.Code, errParseRequest.ToError()), errParseRequest.ToError()
	}

	err := apigateway.ParseAPIGatewayRequestBody(request.Body, &streamRequest)
	if err != nil {
		return apigateway.APIGatewayError(http.StatusBadRequest, err), err
	}

	stream_handler := stream_micro.StreamUseCases{
		Ctx:       ctx,
		TableName: STREAM_TABLE,
	}

	stream, errCreate := stream_handler.UpdateStreamById(id_stream, &streamRequest)
	if errCreate != nil {
		return apigateway.APIGatewayError(errCreate.Code, errCreate.ToError()), errCreate.ToError()
	}

	return apigateway.APIGatewayDataResponse(http.StatusOK, stream), nil
}

func main() {
	lambda.Start(handler)
}
