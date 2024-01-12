package configuration

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoDBAWSClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, _ := config.LoadDefaultConfig(ctx)
	log.Printf("AWS Dynamo Client connected successfully")
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
	)
	log.Printf("Local Client connected successfully")
	return dynamodb.NewFromConfig(cfg), nil
}
