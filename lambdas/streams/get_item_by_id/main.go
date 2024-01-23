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
	id_stream, errParseRequest := apigateway.ParseAPIGatewayRequestParameters(request, "stream_id")
	if errParseRequest != nil {
		return apigateway.APIGatewayError(errParseRequest.Code, errParseRequest.ToError()), errParseRequest.ToError()
	}

	stream, errGetAllStream := stream_handler.GetItemById(ctx, id_stream)
	if errGetAllStream != nil {
		return apigateway.APIGatewayError(errGetAllStream.Code, errGetAllStream.ToError()), errGetAllStream.ToError()
	}

	return apigateway.APIGatewayDataResponse(http.StatusOK, stream), nil
}

func main() {
	lambda.Start(handler)
}
