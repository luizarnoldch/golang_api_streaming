package apigateway

import (
	"encoding/json"
)

// ParseAPIGatewayRequestBody unmarshals a JSON string into the provided pointer to an interface.
func ParseAPIGatewayRequestBody(apiBody string, entity interface{}) error {
	err := json.Unmarshal([]byte(apiBody), entity)
	if err != nil {
		return err
	}
	return nil
}
