package dynamodb

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoDBAWSClient(ctx context.Context) (*dynamodb.Client, error) {
	if (os.Getenv("ENV") == "dev") {
		log.Println(os.Getenv("ENV"))
		return GetLocalDynamoDBClient(ctx)
	}
	
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("error while getting client from AWS")
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}

func GetLocalEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return aws.Endpoint{URL: "http://localhost:8000"}, nil
}

func GetLocalDynamoDBClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, _ := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				GetLocalEndpoint,
			),
		),
		config.WithRetryer(func() aws.Retryer {
			return retry.AddWithMaxAttempts(retry.NewStandard(), 5) // Increase max attempts
		}),
	)
	log.Printf("Local Client connected successfully")
	return dynamodb.NewFromConfig(cfg), nil
}
