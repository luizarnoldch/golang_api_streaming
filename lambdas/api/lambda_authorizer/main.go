package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	cognitoUtil "main/utils/cognito"
	iamUtil "main/utils/iam"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	COGNITO_REGION     = os.Getenv("COGNITO_REGION")
	USER_POOL_ID       = os.Getenv("USER_POOL_ID")
	USER_APP_CLIENT_ID = os.Getenv("USER_APP_CLIENT_ID")
)

func logClaims(claims map[string]interface{}) {
	for key, value := range claims {
		switch v := value.(type) {
		case string:
			fmt.Printf("Claim %s: %s\n", key, v)
		case []interface{}:
			fmt.Printf("Claim %s: %v\n", key, v)
		case float64:
			fmt.Printf("Claim %s: %f\n", key, v)
		case bool:
			fmt.Printf("Claim %s: %t\n", key, v)
		default:
			fmt.Printf("Claim %s: unknown type\n", key)
		}
	}
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	log.Printf("Event Lambda Authorizer: %s\n", string(eventJSON))

	keyURL := cognitoUtil.GenerateCognitoKeyURL(COGNITO_REGION, USER_POOL_ID)
	token := event.AuthorizationToken

	claims, err := cognitoUtil.VerifyPublicKeyToken(ctx, token, keyURL, USER_APP_CLIENT_ID)
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("error: Invalid token")
	}
	resource := cognitoUtil.GenerateCognitoGlobalResource(event.MethodArn)

	// * Return Deny Access in cases of token valid but no authorized
	// return generatePolicy("user", "Deny", resource, claims), nil

	all_claims := map[string]interface{}{
		// Shared claims
		"acr":        claims["acr"].(string),
		"amr":        claims["amr"].([]interface{}),
		"at_hash":    claims["at_hash"].(string),
		"auth_time":  claims["auth_time"].(float64),
		"azp":        claims["azp"].(string),
		"exp":        claims["exp"].(float64),
		"iat":        claims["iat"].(float64),
		"iss":        claims["iss"].(string),
		"jti":        claims["jti"].(string),
		"nbf":        claims["nbf"].(float64),
		"nonce":      claims["nonce"].(string),
		"origin_jti": claims["origin_jti"].(string),
		"sub":        claims["sub"].(string),
		"token_use":  claims["token_use"].(string),

		// ID token claims
		"identities":       claims["identities"].([]interface{}),
		"aud":              claims["aud"].(string),
		"cognito:username": claims["cognito:username"].(string),

		// Access token claims
		"username":   claims["username"].(string),
		"client_id":  claims["client_id"].(string),
		"scope":      claims["scope"].(string),
		"device_key": claims["device_key"].(string),
		"event_id":   claims["event_id"].(string),
		"version":    claims["version"].(string),

		// Custom claims
		// "email":          claims["email"].(string),
		// "company_id":     claims["company_id"].(string),
		// "user_jira_id":   userI.UserJiraID,
		// "company_jira_id": companyI.CompanyJiraID,
	}

	logClaims(all_claims)

	custom_claims := map[string]interface{}{
		// Shared claims
		// "acr":            claims["acr"].(string),
		// "amr":            claims["amr"].([]interface{}),
		// "at_hash":        claims["at_hash"].(string),
		"auth_time": claims["auth_time"].(float64),
		// "azp":            claims["azp"].(string),
		"exp": claims["exp"].(float64),
		"iat": claims["iat"].(float64),
		"iss": claims["iss"].(string),
		"jti": claims["jti"].(string),
		// "nbf":            claims["nbf"].(float64),
		// "nonce":          claims["nonce"].(string),
		// "origin_jti":     claims["origin_jti"].(string),
		"sub":       claims["sub"].(string),
		"token_use": claims["token_use"].(string),

		// ID token claims
		// "identities":     claims["identities"].([]interface{}),
		"aud":              claims["aud"].(string),
		"cognito:username": claims["cognito:username"].(string),

		// Access token claims
		// "username":       claims["username"].(string),
		// "client_id":      claims["client_id"].(string),
		// "scope":          claims["scope"].(string),
		// "device_key":     claims["device_key"].(string),
		// "event_id":       claims["event_id"].(string),
		// "version":        claims["version"].(string),

		// Custom claims
		"email":          claims["email"].(string),
		// "email_verified": claims["email"].(string),
	}

	// * Include groups only if they are present in the token
	// if groups, ok := claims["cognito:groups"].([]interface{}); ok && len(groups) > 0 {
	// 	custom_claims["cognito:groups"] = groups[0].(string)
	// }

	// * Include email_verified only if it is present in the token
	// if emailVerified, ok := claims["email_verified"].(bool); ok {
	// 	custom_claims["email_verified"] = emailVerified
	// }

	return iamUtil.GenerateAPIPolicy("user", "Allow", resource, custom_claims), nil

}

func main() {
	lambda.Start(handleRequest)
}
