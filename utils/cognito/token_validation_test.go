package cognito_test

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestCognito(t *testing.T) {
	MethodArn := "arn:aws:execute-api:us-east-1:681318508133:59ng6168w4/Prod/GET/"

	tmp := strings.Split(MethodArn, ":")

	apiGatewayArnTmp := strings.Split(tmp[5], "/")

	resource := fmt.Sprintf("%s:%s:%s:%s:%s:%s/*/*", tmp[0], tmp[1], tmp[2], tmp[3], tmp[4], apiGatewayArnTmp[0])

	log.Println(resource)
}

func TestCognitoValue2(t *testing.T) {
	MethodArn := "arn:aws:execute-api:us-east-1:681318508133:59ng6168w4/Prod/GET/"

	
	apiGatewayArnTmp := strings.Split(MethodArn, "/")

	resource := fmt.Sprintf("%s/*/*",apiGatewayArnTmp[0])

	log.Println(resource)
}