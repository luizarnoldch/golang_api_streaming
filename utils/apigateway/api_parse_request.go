package apigateway

import (
	"encoding/json"

	appError "main/utils/error"

	"github.com/aws/aws-lambda-go/events"
)

func ParseAPIGatewayRequestBody(apiBody string, entity interface{}) error {
	err := json.Unmarshal([]byte(apiBody), entity)
	if err != nil {
		return err
	}
	return nil
}

func ParseAPIGatewayRequestParameters(request events.APIGatewayProxyRequest, parameter string) (string, *appError.Error) {
	param := request.PathParameters[parameter]
	if param == "" {
		return "", appError.NewNotFoundError("Path Parameter from APIGateway doesn't exits")
	}
	return param, nil
}
