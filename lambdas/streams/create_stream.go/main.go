package main

import (
	"context"
	stream_handler "main/src/streams/application/use_cases"
	"main/src/streams/domain/model"
	"main/utils/apigateway"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var streamRequest model.Stream
	
    err := apigateway.ParseAPIGatewayRequestBody(request.Body, &streamRequest)
    if err != nil {
        return apigateway.APIGatewayError(http.StatusBadRequest, err), err
    }

    stream, errCreate := stream_handler.CreateStream(ctx, &streamRequest)
    if errCreate != nil {
        return apigateway.APIGatewayError(errCreate.Code, errCreate.ToError()), errCreate.ToError()
    }

    return apigateway.APIGatewayDataResponse(http.StatusOK, stream), nil
}


func main() {
	lambda.Start(handler)
}
