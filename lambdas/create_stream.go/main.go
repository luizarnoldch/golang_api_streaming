package main

import (
	"context"
	"main/src/domain/model"
	"main/src/infrastructure/api/stream"
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

    streams, err := stream.CreateStream(ctx, &streamRequest)
    if err != nil {
        return apigateway.APIGatewayError(http.StatusInternalServerError, err), err
    }

    return apigateway.APIGatewayDataResponse(http.StatusOK, streams), nil
}


func main() {
	lambda.Start(handler)
}
