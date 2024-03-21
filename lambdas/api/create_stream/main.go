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
	
    stream_handler := stream_micro.StreamUseCases{
        Ctx: ctx,
        TableName: STREAM_TABLE,
    }

    err := apigateway.ParseAPIGatewayRequestBody(request.Body, &streamRequest)
    if err != nil {
        return apigateway.APIGatewayError(http.StatusBadRequest, err), err
    }

    stream, errCreate := stream_handler.CreateStream(&streamRequest)
    if errCreate != nil {
        return apigateway.APIGatewayError(errCreate.Code, errCreate.ToError()), errCreate.ToError()
    }

    return apigateway.APIGatewayDataResponse(http.StatusOK, stream), nil
}


func main() {
	lambda.Start(handler)
}
