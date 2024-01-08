package configuration

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)


func GetDynamoDBAWSClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Failed to load DynamoDB AWS SDK v2 config")
		return nil, err
	}
	log.Printf("AWS Dynamo Client connected successfully")
	return dynamodb.NewFromConfig(cfg), nil
}

func GetLocalDynamoDBClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: "http://localhost:8000"}, nil
				},
			),
		),
	)

	if err != nil {
		log.Fatal("Failed to load Local DynamoDB AWS SDK v2 config")
		return nil, err
	}
	log.Printf("Local Client connected successfully")
	return dynamodb.NewFromConfig(cfg), nil
}
