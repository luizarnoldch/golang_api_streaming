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

    stream, errCreate := stream.CreateStream(ctx, &streamRequest)
    if errCreate != nil {
        return apigateway.APIGatewayError(errCreate.Code, errCreate.ToError()), errCreate.ToError()
    }

    return apigateway.APIGatewayDataResponse(http.StatusOK, stream), nil
}


func main() {
	lambda.Start(handler)
}
