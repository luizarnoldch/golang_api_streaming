package cognito

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

func GenerateCognitoKeyURL(region, userPoolID string) string {
	return fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
}

func GenerateCognitoGlobalResource(methodArn string) string {
	apiGatewayArnTmp := strings.Split(methodArn, "/")
	return fmt.Sprintf("%s/*/*",apiGatewayArnTmp[0])
}

func VerifyPublicKeyToken(ctx context.Context, tokenString, keyURL, appClientID string) (jwt.MapClaims, error) {
	keySet, err := jwk.Fetch(ctx, keyURL)
	if err != nil {
		log.Printf("Error fetching JWK: %s\n", err)
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("value 'kid' not found in token header")
		}
		key, ok := keySet.LookupKeyID(kid)
		if !ok {
			return nil, fmt.Errorf("signature not found: %v", err)
		}
		keyRaw, err := jwk.PublicRawKeyOf(key)
		if err != nil {
			return nil, err
		}
		return keyRaw, nil
	})

	if err != nil {
		log.Printf("Error decoding token: %s\n", err)
		return nil, err
	}

	if !token.Valid {
		log.Printf("Invalid token\n")
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Invalid token claims\n")
		return nil, fmt.Errorf("invalid token claims")
	}
	exp, ok := claims["exp"].(float64)
	if ok && time.Now().Unix() > int64(exp) {
		log.Printf("Invalid token, token is expired\n")
		return nil, fmt.Errorf("invalid token, token is expired")
	}
	if !ok {
		log.Printf("Missing 'exp' claim\n")
		return nil, fmt.Errorf("missing 'exp' claim")
	}
	if aud, ok := claims["aud"].(string); !ok || aud != appClientID {
		log.Printf("Invalid token, 'aud' claim does not match appClientID\n")
		return nil, fmt.Errorf("invalid token, 'aud' claim does not match appClientID")
	}

	return claims, nil
}
